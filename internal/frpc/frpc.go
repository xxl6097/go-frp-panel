package frpc

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"os/signal"
	"path"
	"syscall"
	"time"

	"github.com/avast/retry-go/v4"
	"github.com/fatedier/frp/client"
	"github.com/fatedier/frp/pkg/config"
	v1 "github.com/fatedier/frp/pkg/config/v1"
	"github.com/fatedier/frp/pkg/config/source"
	"github.com/fatedier/frp/pkg/config/v1/validation"
	httppkg "github.com/fatedier/frp/pkg/util/http"
	"github.com/fatedier/frp/pkg/util/log"
	"github.com/fatedier/frp/pkg/util/system"
	"github.com/xxl6097/glog/pkg/z"
	"github.com/xxl6097/go-frp-panel/internal/com/model"
	comm2 "github.com/xxl6097/go-frp-panel/pkg/comm"
	"github.com/xxl6097/go-frp-panel/pkg/comm/iface"
	"github.com/xxl6097/go-frp-panel/pkg/comm/ws"
	"github.com/xxl6097/go-frp-panel/pkg/frp"
	"github.com/xxl6097/go-frp-panel/pkg/utils"
	"github.com/xxl6097/go-service/pkg/gs/igs"
)

type frpClient struct {
	err         error
	svr         *client.Service
	adminConfig *model.FrpcBuffer
	cfgFilePath string
	cfg         *v1.ClientCommonConfig
	proxyCfg    []v1.ProxyConfigurer
	visitorCfg  []v1.VisitorConfigurer
}
type frpc struct {
	install        igs.Service
	upgrade        iface.IComm
	mainFrpcClient *frpClient
	svrs           map[string]*frpClient
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
	z.Debug("加载配置文件", cfgFilePath)
	cfg, proxyCfgs, visitorCfgs, isLegacyFormat, err := config.LoadClientConfig(cfgFilePath, true)
	if err != nil {
		z.Debug("加载配置文件失败", cfgFilePath, err)
		return nil, fmt.Errorf("load adminConfig file %s not exists", cfgFilePath)
	}
	if isLegacyFormat {
		fmt.Printf("WARNING: ini format is deprecated and the support will be removed in the future, " +
			"please use yaml/json/toml format instead!\n")
	}

	warning, err := validation.ValidateAllClientConfig(cfg, proxyCfgs, visitorCfgs, nil)
	if warning != nil {
		z.Errorf("加载配置文件告警: %v\n", warning)
	}
	if err != nil {
		z.Errorf("配置文件校验失败: %v\n", err)
		return nil, fmt.Errorf("ValidateAllClientConfig adminConfig file %v err", err)
	}

	system.EnableCompatibilityMode()
	log.InitLogger(cfg.Log.To, cfg.Log.Level, int(cfg.Log.MaxDays), cfg.Log.DisablePrintColor)
	aggregator, err := newConfigAggregator(proxyCfgs, visitorCfgs)
	if err != nil {
		return nil, err
	}
	svr, err := client.NewService(client.ServiceOptions{
		Common:                 cfg,
		ConfigSourceAggregator: aggregator,
		ConfigFilePath:         cfgFilePath,
	})
	if err != nil {
		return nil, err
	}
	this := &frpc{
		install: i,
		svrs:    make(map[string]*frpClient),
		upgrade: comm2.NewCommApi(i),
		mainFrpcClient: &frpClient{
			err:         nil,
			svr:         svr,
			cfgFilePath: cfgFilePath,
			cfg:         cfg,
			proxyCfg:    proxyCfgs,
			visitorCfg:  visitorCfgs,
		},
	}

	decodeConfigAndRunWebSocket(this, this.mainFrpcClient)

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
	this.svrs[name] = this.mainFrpcClient
	return this, nil
}

func (this *frpc) Run() error {
	err := retry.Do(func() error {
		this.mainFrpcClient.err = nil
		e := this.mainFrpcClient.svr.Run(context.Background())
		if e != nil {
			z.Errorf("mainfrpc 客户端连接失败[%s]: %v", this.mainFrpcClient.cfgFilePath, e)
			this.mainFrpcClient.err = e
		}
		return e
	}, retry.Delay(time.Second*5), retry.Attempts(0))

	if err != nil {
		z.Error("启动失败", err)
	}
	return err
}

func decodeConfigAndRunWebSocket(this *frpc, cls *frpClient) {
	//defer glog.Flush()
	if cls != nil && cls.cfg != nil && cls.cfg.Metadatas != nil {
		secret := cls.cfg.Metadatas["secret"]
		z.Debugf("secret %+v", secret)
		if secret != "" {
			cls.adminConfig = frp.DecodeSecret(secret)
			z.Debugf("解析secret %+v", cls.adminConfig)
			if cls.adminConfig == nil {
				z.Debug("adminConfig nil 无法启动wensocket ")
				return
			}
			id := cls.adminConfig.User.ID
			addr := fmt.Sprintf("%s:%d", cls.adminConfig.ServerAddr, cls.adminConfig.ServerAdminPort)
			authorization := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", cls.adminConfig.AdminUser, cls.adminConfig.AdminPass)))
			ws.GetClientInstance().NewClient(id, addr, authorization)
			ws.GetClientInstance().SetMessageHandler(this.onWebSocketMessageHandle)
			ws.GetClientInstance().SetOpenHandler(this.onWebSocketOpenHandle)
		}
	} else {
		z.Error("cfg.Metadatas is nil")
	}
}

func (this *frpc) handleTermSignal(svr *client.Service) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	svr.GracefulClose(500 * time.Millisecond)
}

// newConfigAggregator 构建 frp v0.70+ 客户端所需的配置聚合器（替代旧版 ServiceOptions 的 ProxyCfgs/VisitorCfgs 字段）。
func newConfigAggregator(proxyCfgs []v1.ProxyConfigurer, visitorCfgs []v1.VisitorConfigurer) (*source.Aggregator, error) {
	configSource := source.NewConfigSource()
	if err := configSource.ReplaceAll(proxyCfgs, visitorCfgs); err != nil {
		return nil, fmt.Errorf("failed to set config source: %w", err)
	}
	return source.NewAggregator(configSource), nil
}
