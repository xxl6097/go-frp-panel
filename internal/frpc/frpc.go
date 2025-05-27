package frpc

import (
	"context"
	"encoding/base64"
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
	"github.com/xxl6097/go-frp-panel/internal/com/model"
	comm2 "github.com/xxl6097/go-frp-panel/pkg/comm"
	"github.com/xxl6097/go-frp-panel/pkg/comm/iface"
	"github.com/xxl6097/go-frp-panel/pkg/comm/ws"
	"github.com/xxl6097/go-frp-panel/pkg/frp"
	"github.com/xxl6097/go-frp-panel/pkg/utils"
	"github.com/xxl6097/go-service/pkg/gs/igs"
	"os"
	"os/signal"
	"path"
	"syscall"
	"time"
)

type frpClient struct {
	svr            *client.Service
	config         *model.FrpcBuffer
	configFilePath string
	cfg            *v1.ClientCommonConfig
	proxyCfg       []v1.ProxyConfigurer
	visitorCfg     []v1.VisitorConfigurer
}
type frpc struct {
	install igs.Service
	upgrade iface.IComm
	cls     *frpClient
	svrs    map[string]*frpClient
}

func New(i igs.Service) (iface.IFrpc, error) {
	cfgDir, err := frp.GetFrpcTomlDir()
	if err != nil {
		return nil, err
	}
	cfgFilePath, err := frp.GetFrpcMainTomlFilePath()
	if err != nil {
		return nil, err
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
	this := &frpc{
		install: i,
		svrs:    make(map[string]*frpClient),
		upgrade: comm2.NewCommApi(i),
		cls: &frpClient{
			svr:            svr,
			configFilePath: cfgFilePath,
			cfg:            cfg,
			proxyCfg:       proxyCfgs,
			visitorCfg:     visitorCfgs,
		},
	}

	decodeConfigAndRunWebSocket(this, this.cls)

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

	go this.runMultipleClients(cfgDir)
	name := path.Base(cfgFilePath)
	this.svrs[name] = this.cls
	return this, nil
}

func decodeConfigAndRunWebSocket(this *frpc, cls *frpClient) {
	defer glog.Flush()
	if cls != nil && cls.cfg != nil && cls.cfg.Metadatas != nil {
		secret := cls.cfg.Metadatas["secret"]
		glog.Debugf("secret %+v", secret)
		if secret != "" {
			cls.config = frp.DecodeSecret(secret)
			glog.Debugf("解析secret %+v", cls.config)
			if cls.config == nil {
				glog.Debug("config nil 无法启动wensocket ")
				return
			}
			id := cls.config.User.ID
			addr := fmt.Sprintf("%s:%d", cls.config.ServerAddr, cls.config.ServerAdminPort)
			authorization := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", cls.config.AdminUser, cls.config.AdminPass)))
			ws.GetClientInstance().NewClient(id, addr, authorization)
			ws.GetClientInstance().SetMessageHandler(this.onWebSocketMessageHandle)
			ws.GetClientInstance().SetOpenHandler(this.onWebSocketOpenHandle)
		}
	} else {
		glog.Error("cfg.Metadatas is nil")
	}
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
			glog.Errorf("frpc客户端连接失败[%s]: %v", this.cls.configFilePath, e)
		}
		return e
	}, retry.Delay(time.Second*5), retry.Attempts(10))

	if err != nil {
		glog.Error("启动失败", err)
	}
	return err
}
