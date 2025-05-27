package frps

import (
	"context"
	"fmt"
	v1 "github.com/fatedier/frp/pkg/config/v1"
	"github.com/fatedier/frp/pkg/config/v1/validation"
	httppkg "github.com/fatedier/frp/pkg/util/http"
	logfrps "github.com/fatedier/frp/pkg/util/log"
	"github.com/fatedier/frp/pkg/util/system"
	"github.com/fatedier/frp/server"
	"github.com/xxl6097/glog/glog"
	"github.com/xxl6097/go-frp-panel/pkg/comm"
	iface2 "github.com/xxl6097/go-frp-panel/pkg/comm/iface"
	"github.com/xxl6097/go-frp-panel/pkg/comm/sse"
	"github.com/xxl6097/go-frp-panel/pkg/comm/ws"
	"github.com/xxl6097/go-frp-panel/pkg/model"
	"github.com/xxl6097/go-frp-panel/pkg/utils"
	"github.com/xxl6097/go-service/pkg/gs/igs"
	"os"
	"path/filepath"
)

type frps struct {
	svr                    *server.Service
	webServer              *httppkg.Server
	cfg                    *v1.ServerConfig
	install                igs.Service
	upgrade                iface2.IComm
	cloudApi               *model.CloudApi
	webSocketApi           iface2.IWebSocket
	sseApi                 iface2.ISSE
	binDir                 string
	cfgFilePath            string
	frpcGithubDownloadUrls []string
	frpsGithubDownloadUrls []string
	githubProxys           []string
}

func (this *frps) SetCloudApi(api *model.CloudApi) {
	if api == nil {
		return
	}
	this.cloudApi = api
}

func (this *frps) GetCloudApi() *model.CloudApi {
	return this.cloudApi
}

func New(cfg *v1.ServerConfig, install igs.Service) (iface2.IFrps, error) {
	sseApi := sse.NewServer()
	sseApi.Start()

	binPath, err := os.Executable()
	if err != nil {
		glog.Error(fmt.Sprintf("获取当前可执行文件路径出错: %v\n", err))
		return nil, err
	}
	if cfg == nil {
		return nil, fmt.Errorf("服务器配置信息空")
	}

	if cfg.HTTPPlugins == nil || len(cfg.HTTPPlugins) <= 0 {
		cfg.HTTPPlugins = make([]v1.HTTPPluginOptions, 0)
	}

	cfg.HTTPPlugins = append(cfg.HTTPPlugins, v1.HTTPPluginOptions{
		Name: "frps-plugin",
		Addr: fmt.Sprintf("%s:%d", cfg.WebServer.Addr, cfg.WebServer.Port),
		Path: "/handler",
		Ops:  []string{"Login", "NewWorkConn", "NewUserConn", "NewProxy", "Ping"},
	})

	if GetCfgModel() == nil {
		SetCfgModel(&CfgModel{
			Frps: *cfg,
		})
	}

	//cfg := &v1.ServerConfig{}
	//err = json.Unmarshal(content, cfg)
	//if err != nil {
	//	glog.Error(err)
	//	return nil, err
	//}
	cfg.Complete()
	warning, err := validation.ValidateServerConfig(cfg)
	if warning != nil {
		fmt.Printf("WARNING: %v\n", warning)
	}

	system.EnableCompatibilityMode()
	logfrps.InitLogger(cfg.Log.To, cfg.Log.Level, int(cfg.Log.MaxDays), cfg.Log.DisablePrintColor)
	svr, err := server.NewService(cfg)
	if err != nil {
		glog.Fatalf("new frps err: %v", err)
	}
	webServer, err := utils.GetPointerInstance[httppkg.Server]("webServer", svr)
	if err != nil {
		glog.Fatalf("new frps err: %v", err)
	}
	f := &frps{
		cfg:          cfg,
		webServer:    webServer,
		svr:          svr,
		cloudApi:     nil,
		install:      install,
		upgrade:      comm.NewCommApi(install),
		binDir:       filepath.Dir(binPath),
		webSocketApi: ws.NewWebSocket(),
		sseApi:       sseApi,
	}
	f.webSocketApi.SetWebSocket(f)
	f.sseApi.SetSSECallBack(f)
	f.InitClientsConfig()
	//webServer.RouteRegister(f.proxyHandlers)
	webServer.RouteRegister(f.handlers)
	webServer.RouteRegister(f.adminHandlers)
	webServer.RouteRegister(f.userHandlers)
	webServer.RouteRegister(f.webSocketHandler)
	webServer.RouteRegister(f.sseHandler)
	f.CheckClients()
	return f, nil
}

func (this *frps) GetServer() *httppkg.Server {
	return this.webServer
}

func (this *frps) CfgFile(cfgFilePath string) {
	this.cfgFilePath = cfgFilePath
}
func (this *frps) Close() {
	if this.svr == nil {
		return
	}
	this.svr.Close()
}

func (this *frps) Run() {
	inet, _ := utils.GetDeviceInfo()
	ipaddr := fmt.Sprintf("http://localhost")
	if inet.Ipv4 != "" {
		ipaddr = fmt.Sprintf("http://%s", inet.Ipv4)
	}
	fmt.Printf("Frps Admin %s:%d\n", ipaddr, this.cfg.WebServer.Port)
	this.svr.Run(context.Background())
}

func test() {
	//err := config.LoadConfigure(content, this.svrCfg, strict)
	//if err != nil {
	//	fmt.Println("Serve", err)
	//	return err
	//}
	//this.svrCfg.Complete()
	//warning, err := validation.ValidateServerConfig(this.svrCfg)
	//if warning != nil {
	//	fmt.Printf("WARNING: %v\n", warning)
	//}
	//if err != nil {
	//	fmt.Println("Serve", err)
	//	return err
	//}
}
