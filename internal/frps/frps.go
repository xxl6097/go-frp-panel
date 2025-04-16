package frps

import (
	"context"
	"encoding/json"
	"fmt"
	v1 "github.com/fatedier/frp/pkg/config/v1"
	"github.com/fatedier/frp/pkg/config/v1/validation"
	httppkg "github.com/fatedier/frp/pkg/util/http"
	logfrps "github.com/fatedier/frp/pkg/util/log"
	"github.com/fatedier/frp/pkg/util/system"
	"github.com/fatedier/frp/server"
	"github.com/xxl6097/glog/glog"
	_ "github.com/xxl6097/go-frp-panel/assets/frps"
	"github.com/xxl6097/go-frp-panel/internal/comm"
	"github.com/xxl6097/go-frp-panel/internal/comm/iface"
	"github.com/xxl6097/go-frp-panel/pkg/utils"
	"github.com/xxl6097/go-service/gservice/gore"
	"log"
	"os"
	"path/filepath"
)

type frps struct {
	svr       *server.Service
	webServer *httppkg.Server
	cfg       *v1.ServerConfig
	install   gore.IGService
	upgrade   iface.IComm
	binDir    string
}

func NewFrps(content []byte, install gore.IGService) (iface.IFrps, error) {
	binPath, err := os.Executable()
	if err != nil {
		glog.Error(fmt.Sprintf("获取当前可执行文件路径出错: %v\n", err))
		return nil, err
	}

	cfg := &v1.ServerConfig{}
	err = json.Unmarshal(content, cfg)
	if err != nil {
		glog.Error(err)
		return nil, err
	}
	cfg.Complete()
	warning, err := validation.ValidateServerConfig(cfg)
	if warning != nil {
		fmt.Printf("WARNING: %v\n", warning)
	}

	system.EnableCompatibilityMode()
	logfrps.InitLogger(cfg.Log.To, cfg.Log.Level, int(cfg.Log.MaxDays), cfg.Log.DisablePrintColor)
	svr, err := server.NewService(cfg)
	if err != nil {
		log.Fatalf("new frps err: %v", err)
	}
	webServer := utils.GetPointerInstance[httppkg.Server]("webServer", svr)
	f := &frps{
		cfg:       cfg,
		webServer: webServer,
		svr:       svr,
		install:   install,
		upgrade:   comm.NewCommApi(install, GetCfgModel()),
		binDir:    filepath.Dir(binPath),
	}
	webServer.RouteRegister(f.handlers)
	webServer.RouteRegister(f.adminHandlers)
	webServer.RouteRegister(f.userHandlers)
	f.CheckClients()
	return f, nil
}

func (this *frps) Close() {
	if this.svr == nil {
		return
	}
	this.svr.Close()
}

func (this *frps) Run() {
	fmt.Printf("Run: http://127.0.0.1:%d\n", this.cfg.WebServer.Port)
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
