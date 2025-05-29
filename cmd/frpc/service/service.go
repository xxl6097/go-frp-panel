package service

import (
	"fmt"
	v1 "github.com/fatedier/frp/pkg/config/v1"
	"github.com/kardianos/service"
	"github.com/xxl6097/glog/glog"
	"github.com/xxl6097/go-frp-panel/internal/frpc"
	"github.com/xxl6097/go-frp-panel/pkg"
	"github.com/xxl6097/go-frp-panel/pkg/frp"
	frpc2 "github.com/xxl6097/go-frp-panel/pkg/frp/frpc"
	"github.com/xxl6097/go-frp-panel/pkg/utils"
	"github.com/xxl6097/go-service/pkg/gs"
	"github.com/xxl6097/go-service/pkg/gs/igs"
	utils2 "github.com/xxl6097/go-service/pkg/utils"
	"path/filepath"
)

type Service struct {
	wsc *v1.WebServerConfig
}

func (this *Service) OnFinish() {
	if this.wsc != nil {
		face, e := utils.GetDeviceInfo()
		var ip string
		if e == nil {
			ip = face.Ipv4
		}
		glog.Infof("\n登录地址：http://%s:%d\n用户信息：%s/%s", ip, this.wsc.Port, this.wsc.User, this.wsc.Password)
	}
}

func Bootstrap() {
	defer glog.Flush()
	servs := Service{}
	err := gs.Run(&servs)
	if err != nil {
		glog.Error("程序启动出错了", err)
	}
	//glog.Println("服务程序启动成功", os.Getegid())
}

func (s *Service) OnConfig() *service.Config {
	return &service.Config{
		Name:        pkg.AppName,
		DisplayName: pkg.DisplayName,
		Description: pkg.Description,
	}
}
func (s *Service) OnVersion() string {
	//fmt.Println(string(ukey.GetBuffer()))
	pkg.Version()
	return pkg.AppVersion
}

func (this *Service) OnRun(i igs.Service) error {
	//frpc.Assert()
	glog.Printf("启动frpc_%s\n", pkg.AppVersion)
	cfg := frpc.GetCfgModel()
	if cfg == nil {
		return fmt.Errorf("程序配置文件未初始化")
	}
	//svv, err := frpc.NewFrpc(i)
	svv, err := frpc2.NewFrpc(i)
	if err != nil {
		glog.Error("启动frpc失败", err)
		glog.Printf("启动frp_%s失败\n", pkg.AppVersion)
		return err
	}
	err = svv.Run()
	return err
}

func (this *Service) GetAny(binDir string) []byte {
	cfg := this.menu()
	this.wsc = &cfg.Frpc.WebServer
	err := frp.WriteFrpcMainConfigWithDir(binDir, cfg.Frpc)
	if err != nil {
		glog.Warnf("write content to frpc config file error: %v", err)
		return nil
	}
	return cfg.Bytes()
}

func (this *Service) menu() *frpc.CfgModel {
	var cfg *v1.ClientConfig
	cfm := frpc.GetCfgModel()
	if cfm == nil {
		cfg = &v1.ClientConfig{}
	} else {
		cfg = &cfm.Frpc
	}
	cfg.Log = v1.LogConfig{
		To:      filepath.Join(glog.AppHome("frpc", "log"), "frpc.log"),
		MaxDays: 7,
		Level:   "error",
	}
	if cfg.ClientCommonConfig.ServerAddr == "" {
		cfg.ClientCommonConfig.ServerAddr = utils2.InputString("Frps服务器地址:")
	}
	if cfg.ClientCommonConfig.ServerPort <= 0 || cfg.ClientCommonConfig.ServerPort > 65535 {
		cfg.ClientCommonConfig.ServerPort = utils2.InputInt("Frps服务器绑定端口:")
	}
	if cfg.WebServer.Addr == "" {
		cfg.WebServer.Addr = "0.0.0.0"
	}
	if cfg.WebServer.Port == 0 {
		cfg.WebServer.Port = utils2.InputIntDefault("管理后台端口(6400)", 6400)
	}
	if cfg.WebServer.User == "" {
		cfg.WebServer.User = utils2.InputStringEmpty("管理后台用户名(admin):", "admin")
	}
	if cfg.WebServer.Password == "" {
		cfg.WebServer.Password = utils2.InputString("管理后台密码：")
	}

	return &frpc.CfgModel{
		Frpc: *cfg,
	}
}
