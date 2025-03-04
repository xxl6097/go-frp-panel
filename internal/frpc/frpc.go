package frpc

import (
	"cmp"
	"context"
	"encoding/json"
	"fmt"
	"github.com/fatedier/frp/client"
	"github.com/fatedier/frp/client/proxy"
	"github.com/fatedier/frp/pkg/config"
	v1 "github.com/fatedier/frp/pkg/config/v1"
	"github.com/fatedier/frp/pkg/config/v1/validation"
	httppkg "github.com/fatedier/frp/pkg/util/http"
	"github.com/fatedier/frp/pkg/util/log"
	"github.com/fatedier/frp/pkg/util/system"
	"github.com/xxl6097/glog/glog"
	_ "github.com/xxl6097/go-frp-panel/assets/frpc"
	"github.com/xxl6097/go-frp-panel/internal/comm"
	"github.com/xxl6097/go-frp-panel/internal/comm/iface"
	"github.com/xxl6097/go-frp-panel/pkg/utils"
	"github.com/xxl6097/go-service/gservice/gore"
	utils2 "github.com/xxl6097/go-service/gservice/utils"
	"io/fs"
	"os"
	"os/signal"
	"path"
	"path/filepath"
	"slices"
	"strings"
	"syscall"
	"time"
)

type frpClient struct {
	svr *client.Service
	cfg *v1.ClientCommonConfig
}
type frpc struct {
	svr            *client.Service
	install        gore.IGService
	configFilePath string
	upgrade        iface.IComm
	svrs           map[string]*frpClient
}

func NewFrpc(i gore.IGService) (*frpc, error) {
	baseDir, err := os.Executable()
	if err != nil {
		return nil, err
	}
	cfgFilePath := filepath.Join(filepath.Dir(baseDir), "config.toml")
	if !utils2.FileExists(cfgFilePath) {
		return nil, fmt.Errorf("config file %s not exists", cfgFilePath)
	}
	cfg, proxyCfgs, visitorCfgs, isLegacyFormat, err := config.LoadClientConfig(cfgFilePath, true)
	if err != nil {
		return nil, fmt.Errorf("load config file %s not exists", cfgFilePath)
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
		return nil, fmt.Errorf("ValidateAllClientConfig config file %v err", err)
	}

	system.EnableCompatibilityMode()
	log.InitLogger(cfg.Log.To, cfg.Log.Level, int(cfg.Log.MaxDays), cfg.Log.DisablePrintColor)
	svr, err := client.NewService(client.ServiceOptions{
		Common:         cfg,
		ProxyCfgs:      proxyCfgs,
		VisitorCfgs:    visitorCfgs,
		ConfigFilePath: cfgFilePath,
	})
	if err != nil {
		return nil, err
	}
	this := &frpc{
		svr:            svr,
		install:        i,
		configFilePath: cfgFilePath,
		svrs:           make(map[string]*frpClient),
		upgrade:        comm.NewCommApi(i, GetCfgModel()),
	}

	shouldGracefulClose := cfg.Transport.Protocol == "kcp" || cfg.Transport.Protocol == "quic"
	if shouldGracefulClose {
		go this.handleTermSignal(svr)
	}

	webServer := utils.GetPointerInstance[httppkg.Server]("webServer", svr)
	if webServer == nil {
		return nil, fmt.Errorf("can't find webServer")
	}
	webServer.RouteRegister(this.adminHandlers)
	err = this.runMultipleClients(filepath.Join(filepath.Dir(baseDir), "config"))
	if err != nil {
		glog.Errorf("runMultipleClients err: %v", err)
	}
	return this, nil
}

func (this *frpc) handleTermSignal(svr *client.Service) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	svr.GracefulClose(500 * time.Millisecond)
}

func (this *frpc) Run() error {
	err := this.svr.Run(context.Background())
	if err != nil {
		glog.Errorf("frpc run error: %v", err)
	}
	return err
}

func (this *frpc) runMultipleClients(cfgDir string) error {
	err := filepath.WalkDir(cfgDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return nil
		}
		ext := strings.ToLower(filepath.Ext(d.Name()))
		if ext != ".toml" {
			return nil
		}
		time.Sleep(time.Millisecond)
		err = this.runClient(path)
		if err != nil {
			glog.Errorf("创建客户端【%s】失败:%v", d.Name(), err)
		} else {
			glog.Infof("创建客户端【%s】成功", d.Name())
		}
		return err
	})
	return err
}

func (this *frpc) deleteClient(cfgFilePath string) error {
	name := path.Base(cfgFilePath)
	glog.Debug("delete frpc", name)
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

func (this *frpc) statusClient(cfgFilePath string) ([]byte, error) {
	name := path.Base(cfgFilePath)
	glog.Debug("status frpc", name)
	cls := this.svrs[name]
	if cls == nil {
		return nil, fmt.Errorf("没有找到客户端句柄")
	}
	svr := cls.svr
	if svr == nil {
		return nil, fmt.Errorf("没有找到客户端服务句柄")
	}
	ctl := utils.GetPointerInstance[client.Control]("ctl", svr)
	if ctl == nil {
		return nil, fmt.Errorf("没有找到服务控制器")
	}
	pm := utils.GetPointerInstance[proxy.Manager]("pm", ctl)
	if pm == nil {
		return nil, fmt.Errorf("没有找到服务代理器")
	}
	var (
		buf []byte
		res client.StatusResp = make(map[string][]client.ProxyStatusResp)
	)
	ps := pm.GetAllProxyStatus()
	glog.Debug("GetAllProxyStatus", len(ps))
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
	log.Infof("Http response [/api/status]")
	buf, _ = json.Marshal(&res)
	return buf, nil
}

func (this *frpc) updateClient(cfgFilePath string) error {
	name := path.Base(cfgFilePath)
	glog.Debug("update frpc", name)
	cls := this.svrs[name]
	if cls == nil {
		return fmt.Errorf("can't find client")
	}
	svr := cls.svr
	if svr == nil {
		return fmt.Errorf("can't find service")
	}
	cliCfg, proxyCfgs, visitorCfgs, _, err := config.LoadClientConfig(cfgFilePath, true)
	if err != nil {
		return fmt.Errorf("reload frpc config error: %v", err)
	}
	if _, err := validation.ValidateAllClientConfig(cliCfg, proxyCfgs, visitorCfgs); err != nil {
		return fmt.Errorf("validate frpc proxy config error: %v", err)
	}

	if err := svr.UpdateAllConfigurer(proxyCfgs, visitorCfgs); err != nil {
		return fmt.Errorf("update frpc proxy config error: %v", err)
	}
	return nil
}

func (this *frpc) runClient(cfgFilePath string) error {
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
		return err
	}
	err = this.startService(cfg, proxyCfgs, visitorCfgs, cfgFilePath)
	return err
}

func (this *frpc) startService(
	cfg *v1.ClientCommonConfig,
	proxyCfgs []v1.ProxyConfigurer,
	visitorCfgs []v1.VisitorConfigurer,
	cfgFile string,
) error {
	cfg.WebServer = v1.WebServerConfig{}
	if cfg.Log.To == "" {
		temp := os.TempDir()
		temp = filepath.Join(temp, "frpc", cfg.User, "logs", "frpc.log")
		utils.DirCheck(temp)
		cfg.Log = v1.LogConfig{
			To:      temp,
			MaxDays: 15,
		}
	}

	log.InitLogger(cfg.Log.To, cfg.Log.Level, int(cfg.Log.MaxDays), cfg.Log.DisablePrintColor)

	if cfgFile != "" {
		log.Infof("start frpc service for config file [%s]", cfgFile)
		defer log.Infof("frpc service for config file [%s] stopped", cfgFile)
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

	name := path.Base(cfgFile)
	this.svrs[name] = &frpClient{
		svr: svr,
		cfg: cfg,
	}
	glog.Debug("create frpc", name)
	shouldGracefulClose := cfg.Transport.Protocol == "kcp" || cfg.Transport.Protocol == "quic"
	// Capture the exit signal if we use kcp or quic.
	if shouldGracefulClose {
		go this.handleTermSignal(svr)
	}
	go func() {
		err1 := svr.Run(context.Background())
		if err1 != nil {
			glog.Errorf("frpc service run err: %v", err1)
		}
	}()
	return nil
}
