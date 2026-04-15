package frpc

import (
	"cmp"
	"context"
	"fmt"
	"io/fs"
	"path"
	"path/filepath"
	"slices"
	"strings"
	"time"

	"github.com/avast/retry-go/v4"
	"github.com/fatedier/frp/client"
	"github.com/fatedier/frp/client/proxy"
	"github.com/fatedier/frp/pkg/config"
	v1 "github.com/fatedier/frp/pkg/config/v1"
	"github.com/fatedier/frp/pkg/config/v1/validation"
	"github.com/fatedier/frp/pkg/util/log"
	"github.com/xxl6097/glog/pkg/z"
	"github.com/xxl6097/glog/pkg/zutil"
	"github.com/xxl6097/go-frp-panel/pkg/frp"
	"github.com/xxl6097/go-frp-panel/pkg/utils"
	utils2 "github.com/xxl6097/go-service/pkg/utils"
)

func (this *frpc) retry(cfgPath string) {
	err := retry.Do(func() error {
		return this.newClient(cfgPath)
	}, retry.Delay(time.Second*5), retry.Attempts(0))

	if err != nil {
		z.Error("启动失败", err)
	}
}

func (this *frpc) runMultipleClients(cfgDir string) {
	err := filepath.WalkDir(cfgDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return nil
		}
		if strings.Compare(strings.ToLower(d.Name()), strings.ToLower(frp.GetFrpcMainTomlFileName())) == 0 {
			return nil
		}
		ext := strings.ToLower(filepath.Ext(d.Name()))
		if ext != ".toml" {
			return nil
		}
		time.Sleep(time.Millisecond)
		err = this.newClient(path)
		if err != nil {
			z.Errorf("创建客户端【%s】失败:%v", d.Name(), err)
			go this.retry(path)
		} else {
			z.Infof("创建客户端【%s】成功", d.Name())
		}
		return err
	})
	if err != nil {
		z.Error(err)
	}
}

func (this *frpc) startService(
	cfg *v1.ClientCommonConfig,
	proxyCfgs []v1.ProxyConfigurer,
	visitorCfgs []v1.VisitorConfigurer,
	cfgFile string,
) error {
	cfg.WebServer = v1.WebServerConfig{}
	if cfg.Log.To == "" {
		temp := filepath.Join(zutil.AppHome(), cfg.User, "app.log")
		cfg.Log = v1.LogConfig{
			To:      temp,
			MaxDays: 7,
		}
	}

	log.InitLogger(cfg.Log.To, cfg.Log.Level, int(cfg.Log.MaxDays), cfg.Log.DisablePrintColor)

	if cfgFile != "" {
		log.Infof("start frpc service for adminConfig file [%s]", cfgFile)
		defer log.Infof("frpc service for adminConfig file [%s] stopped", cfgFile)
	}

	svr, err := client.NewService(client.ServiceOptions{
		Common:         cfg,
		ProxyCfgs:      proxyCfgs,
		VisitorCfgs:    visitorCfgs,
		ConfigFilePath: cfgFile,
	})
	if err != nil {
		return err
	}
	fc := frpClient{
		svr:         svr,
		cfg:         cfg,
		proxyCfg:    proxyCfgs,
		visitorCfg:  visitorCfgs,
		cfgFilePath: cfgFile,
		err:         nil,
	}
	name := path.Base(cfgFile)
	this.svrs[name] = &fc
	decodeConfigAndRunWebSocket(this, &fc)
	z.Debug("创建frpc客户端", name)
	shouldGracefulClose := cfg.Transport.Protocol == "kcp" || cfg.Transport.Protocol == "quic"
	// Capture the exit signal if we use kcp or quic.
	if shouldGracefulClose {
		go this.handleTermSignal(svr)
	}

	//e := retry.Do(func() error {
	//	e := svr.Run(context.Background())
	//	if e != nil {
	//		z.Errorf("创建frpc客户端失败: %s %v\n", cfgFile, e)
	//	}
	//	return e
	//}, retry.Delay(time.Second*5), retry.Attempts(10))

	fc.err = nil
	e := svr.Run(context.Background())
	if e != nil {
		z.Errorf("[%s]创建客户端失败: %v\n", name, e)
		fc.err = e
	}
	//因为Run是阻塞的，能执行到这一行，说明失败了
	//delete(this.svrs, name) // 注释掉，不然获取不到最新的错误信息
	return e
}

func (this *frpc) deleteClient(cfgFilePath string) error {
	name := path.Base(cfgFilePath)
	z.Debug("delete", name)
	cls := this.svrs[name]
	if cls == nil {
		return fmt.Errorf("can't find client")
	}
	svr := cls.svr
	if svr == nil {
		return fmt.Errorf("can't find service")
	}
	svr.Close()
	svr.GracefulClose(100 * time.Millisecond)
	//svr.StatusExporter().GetProxyStatus()
	utils.Delete(cfgFilePath, fmt.Sprintf("客户端:%s", cfgFilePath))
	return nil
}

func (this *frpc) statusClient(cfgFilePath string) (map[string][]client.ProxyStatusResp, error) {
	name := path.Base(cfgFilePath)
	z.Debug("status frpc", name)
	cls := this.svrs[name]
	if cls == nil {
		return nil, fmt.Errorf("客户端未创建")
	}
	if cls.err != nil {
		return nil, cls.err
	}
	svr := cls.svr
	if svr == nil {
		return nil, fmt.Errorf("客户端服务未创建")
	}
	ctl, err := utils.GetPointerInstance[client.Control]("ctl", svr)
	if err != nil {
		z.Debug("GetPointerInstance[client.Control] err", err)
		return nil, err
	}
	if ctl == nil {
		return nil, fmt.Errorf("没有找到服务控制器")
	}
	pm, err := utils.GetPointerInstance[proxy.Manager]("pm", ctl)
	if err != nil {
		z.Debug("GetPointerInstance[proxy.Manager] err", err)
		return nil, err
	}
	if pm == nil {
		return nil, fmt.Errorf("没有找到服务代理器")
	}
	var (
		//buf []byte
		res client.StatusResp = make(map[string][]client.ProxyStatusResp)
	)
	ps := pm.GetAllProxyStatus()
	z.Debug("GetAllProxyStatus", len(ps))
	for _, status := range ps {
		res[status.Type] = append(res[status.Type], client.NewProxyStatusResp(status, cls.cfg.ServerAddr))
	}

	for _, arrs := range res {
		if len(arrs) <= 1 {
			continue
		}
		slices.SortFunc(arrs, func(a, b client.ProxyStatusResp) int {
			return cmp.Compare(a.Name, b.Name)
		})
	}
	//buf, err = json.Marshal(&res)
	//if err != nil {
	//	z.Errorf("json error: %v", err)
	//	return nil, err
	//}
	//return buf, nil
	return res, nil
}

func (this *frpc) updateClient(cfgFilePath string) error {
	name := path.Base(cfgFilePath)
	z.Debug("update clilent", cfgFilePath)
	cls := this.svrs[name]
	if cls == nil {
		return fmt.Errorf("can't find client [%+v]", this.svrs)
	}
	svr := cls.svr
	if svr == nil {
		return fmt.Errorf("can't find service")
	}
	cliCfg, proxyCfgs, visitorCfgs, _, err := config.LoadClientConfig(cfgFilePath, true)
	if err != nil {
		return fmt.Errorf("reload frpc adminConfig error: %v", err)
	}
	if _, err := validation.ValidateAllClientConfig(cliCfg, proxyCfgs, visitorCfgs); err != nil {
		return fmt.Errorf("validate frpc proxy adminConfig error: %v", err)
	}

	if err := svr.UpdateAllConfigurer(proxyCfgs, visitorCfgs); err != nil {
		return fmt.Errorf("update frpc proxy adminConfig error: %v", err)
	}
	cls.cfg = cliCfg
	cls.proxyCfg = proxyCfgs
	cls.visitorCfg = visitorCfgs
	return nil
}

func (this *frpc) upgradeMainConfig() error {
	if this.mainFrpcClient == nil {
		return fmt.Errorf("can't find client")
	}
	svr := this.mainFrpcClient.svr
	if svr == nil {
		return fmt.Errorf("can't find service")
	}
	cliCfg, proxyCfgs, visitorCfgs, _, err := config.LoadClientConfig(this.mainFrpcClient.cfgFilePath, true)
	if err != nil {
		return fmt.Errorf("reload frpc adminConfig error: %v", err)
	}
	if _, err := validation.ValidateAllClientConfig(cliCfg, proxyCfgs, visitorCfgs); err != nil {
		return fmt.Errorf("validate frpc proxy adminConfig error: %v", err)
	}

	if err := svr.UpdateAllConfigurer(proxyCfgs, visitorCfgs); err != nil {
		return fmt.Errorf("update frpc proxy adminConfig error: %v", err)
	}
	z.Infof("success reload conf")
	return nil
}

func (this *frpc) getTcpProxyArray(name string) []int {
	z.Debug("info frpc", name)
	var cls *frpClient
	if name == "" {
		cls = this.mainFrpcClient
	} else {
		cls = this.svrs[name]
	}
	if cls == nil {
		return nil
	}
	if cls.adminConfig == nil {
		return nil
	}
	if cls.adminConfig.User.Ports == nil {
		return nil
	}
	//主客户端
	ports := frp.ParsePorts(cls.adminConfig.User.Ports)
	for _, c := range cls.proxyCfg {
		port := frp.GetPort(c)
		if port > 0 {
			ports = utils.RemoveSlice[int](ports, port)
		}
	}
	return ports
}

func (this *frpc) newClient(cfgFilePath string) error {
	cfg, proxyCfgs, visitorCfgs, isLegacyFormat, err := config.LoadClientConfig(cfgFilePath, true)
	if err != nil {
		return err
	}
	if isLegacyFormat {
		fmt.Printf("WARNING: ini format is deprecated and the support will be removed in the future, " +
			"please use yaml/json/toml format instead!\n")
	}

	warning, err := validation.ValidateAllClientConfig(cfg, proxyCfgs, visitorCfgs)
	if warning != nil {
		fmt.Printf("WARNING: %v\n", warning)
	}
	if err != nil {
		z.Errorf("配置文件校验失败: %s %v\n", cfgFilePath, err)
		return err
	}
	//给一个可能阻塞 / 耗时的同步函数添加超时控制，防止函数无限阻塞导致程序卡住。
	e, _ := utils2.BlockingFunction[error](context.Background(), time.Second*3, func() error {
		return this.startService(cfg, proxyCfgs, visitorCfgs, cfgFilePath)
	})
	if e == nil {
	}
	z.Warnf("运行客户端: %s %v\n", cfgFilePath, e)
	return e
}
