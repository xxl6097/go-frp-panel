package frpc

import (
	"context"
	"fmt"
	"github.com/avast/retry-go/v4"
	"github.com/fatedier/frp/client"
	"github.com/fatedier/frp/pkg/config"
	v1 "github.com/fatedier/frp/pkg/config/v1"
	"github.com/fatedier/frp/pkg/config/v1/validation"
	httppkg "github.com/fatedier/frp/pkg/util/http"
	"github.com/fatedier/frp/pkg/util/log"
	"github.com/fatedier/frp/pkg/util/system"
	"github.com/xxl6097/glog/glog"
	_ "github.com/xxl6097/go-frp-panel/assets/frpc"
	comm2 "github.com/xxl6097/go-frp-panel/pkg/comm"
	"github.com/xxl6097/go-frp-panel/pkg/comm/iface"
	"github.com/xxl6097/go-frp-panel/pkg/utils"
	"github.com/xxl6097/go-service/gservice/gore"
	utils2 "github.com/xxl6097/go-service/gservice/utils"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

type frpClient struct {
	svr            *client.Service
	configFilePath string
	cfg            *v1.ClientCommonConfig
	proxyCfg       []v1.ProxyConfigurer
	visitorCfg     []v1.VisitorConfigurer
}
type frpc struct {
	install   gore.IGService
	upgrade   iface.IComm
	cls       *frpClient
	cfgBuffer *comm2.BufferConfig
	svrs      map[string]*frpClient
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
	glog.Debug("加载配置文件", cfgFilePath)
	cfg, proxyCfgs, visitorCfgs, isLegacyFormat, err := config.LoadClientConfig(cfgFilePath, true)
	if err != nil {
		glog.Debug("加载配置文件失败", cfgFilePath, err)
		return nil, fmt.Errorf("load config file %s not exists", cfgFilePath)
	}
	if isLegacyFormat {
		fmt.Printf("WARNING: ini format is deprecated and the support will be removed in the future, " +
			"please use yaml/json/toml format instead!\n")
	}

	warning, err := validation.ValidateAllClientConfig(cfg, proxyCfgs, visitorCfgs)
	if warning != nil {
		glog.Errorf("加载配置文件告警: %v\n", warning)
	}
	if err != nil {
		glog.Errorf("配置文件校验失败: %v\n", err)
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
	cfgModel := GetCfgModel()
	this := &frpc{
		install: i,
		svrs:    make(map[string]*frpClient),
		upgrade: comm2.NewCommApi(i, cfgModel),
		cls: &frpClient{
			svr:            svr,
			configFilePath: cfgFilePath,
			cfg:            cfg,
			proxyCfg:       proxyCfgs,
			visitorCfg:     visitorCfgs,
		},
	}
	if cfgModel.Cfg != nil {
		this.cfgBuffer = cfgModel.Cfg
	}

	shouldGracefulClose := cfg.Transport.Protocol == "kcp" || cfg.Transport.Protocol == "quic"
	if shouldGracefulClose {
		go this.handleTermSignal(svr)
	}

	webServer, err := utils.GetPointerInstance[httppkg.Server]("webServer", svr)
	if err != nil {
		return nil, err
	}
	if webServer == nil {
		return nil, fmt.Errorf("can't find webServer")
	}
	webServer.RouteRegister(this.adminHandlers)
	go this.runMultipleClients(filepath.Join(filepath.Dir(baseDir), "config"))
	return this, nil
}

func (this *frpc) handleTermSignal(svr *client.Service) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	svr.GracefulClose(500 * time.Millisecond)
}

func (this *frpc) Run() error {
	err := retry.Do(func() error {
		e := this.cls.svr.Run(context.Background())
		if e != nil {
			glog.Errorf("frpc启动错误[%s]: %v", this.cls.configFilePath, e)
		}
		return e
	}, retry.Delay(time.Second*5), retry.Attempts(10))

	if err != nil {
		glog.Error("启动失败", err)
	}
	return err
}
