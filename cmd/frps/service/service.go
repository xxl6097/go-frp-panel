package service

import (
	"fmt"
	v1 "github.com/fatedier/frp/pkg/config/v1"
	"github.com/fatedier/frp/pkg/util/version"
	"github.com/kardianos/service"
	"github.com/xxl6097/glog/glog"
	"github.com/xxl6097/go-frp-panel/internal/frps"
	"github.com/xxl6097/go-frp-panel/pkg"
	"github.com/xxl6097/go-frp-panel/pkg/comm/iface"
	frps2 "github.com/xxl6097/go-frp-panel/pkg/frp/frps"
	"github.com/xxl6097/go-frp-panel/pkg/utils"
	"github.com/xxl6097/go-service/pkg/gs"
	"github.com/xxl6097/go-service/pkg/gs/igs"
	utils2 "github.com/xxl6097/go-service/pkg/utils"
	"path/filepath"
)

type Service struct {
	ifrps     iface.IFrps
	webServer *v1.WebServerConfig
}

func (this *Service) OnFinish() {
	if this.webServer != nil {
		face, e := utils.GetDeviceInfo()
		var ip string
		if e == nil {
			ip = face.Ipv4
		}
		fmt.Printf("登录地址：http://%s:%d\n用户信息：%s/%s\n", ip, this.webServer.Port, this.webServer.User, this.webServer.Password)
	}
}
func Bootstrap() {
	defer glog.Flush()
	svr := &Service{}
	err := gs.Run(svr)
	if err != nil {
		glog.Error("程序启动出错了", err)
	}
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
	//这里需要打印config中buffer原始信息
	pkg.Version()
	return fmt.Sprintf("frps version:%s", version.Full())
}

func (this *Service) OnRun(i igs.Service) error {
	//frps.Assert()
	glog.Printf("启动 %s %s\n", pkg.AppName, pkg.AppVersion)
	cfg := frps.GetCfgModel()
	if cfg == nil {
		return fmt.Errorf("程序配置文件未初始化")
	}
	conf := frps.GetCfgModel().Frps
	//content, err := json.Marshal(conf)
	//if err != nil {
	//	glog.Error(err)
	//	return err
	//}
	//cfgConfig := &v1.ServerConfig{}
	//err = json.Unmarshal(content, cfgConfig)
	//if err != nil {
	//	glog.Error(err)
	//	return err
	//}
	//svv, err := frps2.NewFrps(cfgConfig, i)
	svv, err := frps2.NewFrps(&conf, i)
	if err != nil {
		glog.Printf("启动 %s %s 失败:%v\n%v", pkg.AppName, pkg.AppVersion, err, conf)
		return err
	}
	this.ifrps = svv
	svv.Run()
	return err
}

func (this *Service) GetAny(binDir string) []byte {
	a := this.menu()
	if a == nil {
		return nil
	}
	this.webServer = &a.Frps.WebServer
	return a.Bytes()
}

func (this *Service) menu() *frps.CfgModel {
	cfg := frps.GetCfgModel()
	if cfg == nil {
		cfg = &frps.CfgModel{}
	}
	if cfg.Frps.BindPort <= 0 {
		cfg.Frps.BindPort = utils2.InputIntDefault("Frps绑定端口(默认:6000):", 6000)
	}
	if cfg.Frps.WebServer.Port <= 0 {
		cfg.Frps.WebServer.Port = utils2.InputIntDefault("管理后台端口(默认:6500):", 6500)
	}
	if cfg.Frps.WebServer.Addr == "" {
		cfg.Frps.WebServer.Addr = utils2.InputStringEmpty("管理后台地址(默认:0.0.0.0)：", "0.0.0.0")
	}
	if cfg.Frps.WebServer.User == "" {
		cfg.Frps.WebServer.User = utils2.InputStringEmpty("管理后台用户名(默认:admin)：", "admin")
	}
	if cfg.Frps.WebServer.Password == "" {
		cfg.Frps.WebServer.Password = utils2.InputString("管理后台密码：")
	}
	temp := glog.AppHome("frps", "log")
	cfg.Frps.Log = v1.LogConfig{
		To:      filepath.Join(temp, "frps.log"),
		MaxDays: 3,
		Level:   "error",
	}
	return cfg
}
