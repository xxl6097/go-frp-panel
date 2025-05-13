package service

import (
	"fmt"
	v1 "github.com/fatedier/frp/pkg/config/v1"
	"github.com/fatedier/frp/pkg/util/version"
	"github.com/kardianos/service"
	"github.com/xxl6097/glog/glog"
	"github.com/xxl6097/go-frp-panel/internal/frpc"
	"github.com/xxl6097/go-frp-panel/pkg"
	"github.com/xxl6097/go-frp-panel/pkg/utils"
	"github.com/xxl6097/go-service/gservice"
	"github.com/xxl6097/go-service/gservice/gore"
	"github.com/xxl6097/go-service/gservice/ukey"
	utils2 "github.com/xxl6097/go-service/gservice/utils"
	"os"
	"path/filepath"
)

type Service struct {
	wsc *v1.WebServerConfig
}

func Bootstrap() {
	svr := &Service{}
	err := gservice.Run(svr)
	if err != nil {
		glog.Error("程序启动出错了", err)
	}
	if svr.wsc != nil {
		glog.Infof("登录信息：\nhttp://%s:%d\n用户名密码：%s/%s", utils.GetLocalIp(), svr.wsc.Port, svr.wsc.User, svr.wsc.Password)
	}
	glog.Warnf("OnFinish %+v", svr.wsc)
	glog.Println("服务程序启动成功", os.Getegid())
}

func (s *Service) OnInit() *service.Config {
	return &service.Config{
		Name:        pkg.AppName,
		DisplayName: pkg.DisplayName,
		Description: pkg.Description,
	}
}
func (s *Service) OnVersion() string {
	fmt.Println(string(ukey.GetBuffer()))
	ver := fmt.Sprintf("frpc version:%s", version.Full())
	pkg.Version()
	return ver
}

func (this *Service) OnRun(i gore.IGService) error {
	frpc.Assert()
	glog.Printf("启动frpc_%s\n", pkg.AppVersion)
	cfg := frpc.GetCfgModel()
	if cfg == nil {
		return fmt.Errorf("程序配置文件未初始化")
	}
	svv, err := frpc.NewFrpc(i)
	if err != nil {
		glog.Error("启动frpc失败", err)
		glog.Printf("启动frp_%s失败\n", pkg.AppVersion)
		return err
	}
	err = svv.Run()
	return err
}

func (this *Service) GetAny(binDir string) any {
	cfg := this.menu()
	cfgPath := filepath.Join(binDir, "config.toml")
	this.wsc = &cfg.Frpc.WebServer
	if err := os.WriteFile(cfgPath, utils.ObjectToTomlText(cfg.Frpc), 0o600); err != nil {
		glog.Warnf("write content to frpc config file error: %v", err)
	} else {
		glog.Infof("write content to frpc config file success %s", cfgPath)
	}
	return cfg
}

func (this *Service) menu() *frpc.CfgModel {
	var bindAddr, userName, token, id string
	var bindPort int
	err := frpc.IsInit()
	c := frpc.GetCfgModel()
	//glog.Infof("Frpc: %+v", c.Frpc)
	//glog.Infof("Cfg: %+v", c.Cfg)
	//glog.Error(err)
	if err != nil || c == nil {
		bindAddr = utils2.InputString("Frps服务器地址:")
		bindPort = utils2.InputInt("Frps服务器绑定端口:")
		userName = utils2.InputStringEmpty("请输入用户名(admin):", "admin")
		id = utils2.InputString("请输入ID：")
		token = utils2.InputString("请输入密钥：")
	} else {
		bindAddr = c.Frpc.ServerAddr
		bindPort = c.Frpc.ServerPort
		userName = c.Frpc.User
		token = c.Frpc.Metadatas["token"]
		id = c.Frpc.Metadatas["id"]
	}
	webServer := &v1.WebServerConfig{}
	if c.Cfg != nil && c.Cfg.WebServer != nil && c.Cfg.WebServer.Port != 0 && c.Cfg.WebServer.Addr != "" && c.Cfg.WebServer.User != "" && c.Cfg.WebServer.Password != "" {
		webServer = c.Cfg.WebServer
	}
	if webServer.Addr == "" {
		webServer.Addr = "0.0.0.0"
	}
	if webServer.Port == 0 {
		webServer.Port = utils2.InputIntDefault("管理后台端口(6400)", 6400)
	}
	if webServer.User == "" {
		webServer.User = utils2.InputStringEmpty("管理后台用户名(admin):", "admin")
	}
	if webServer.Password == "" {
		webServer.Password = utils2.InputString("管理后台密码：")
	}

	temp := glog.GetCrossPlatformDataDir("frpc", "log")
	ccc := v1.ClientCommonConfig{
		ServerAddr: bindAddr,
		ServerPort: bindPort,
		User:       userName,
		Metadatas: map[string]string{
			"token": token,
			"id":    id,
		},
		Log: v1.LogConfig{
			To:      filepath.Join(temp, "frpc.log"),
			MaxDays: 7,
			Level:   "error",
		},
		WebServer: *webServer,
	}

	var proxies []v1.TypedProxyConfig
	if c.Cfg.Proxy.GetBaseConfig().LocalPort != 0 && c.Cfg.Proxy.GetBaseConfig().LocalIP != "" {
		proxies = append(proxies, *c.Cfg.Proxy)
	}

	cc := v1.ClientConfig{
		ClientCommonConfig: ccc,
		Proxies:            proxies,
	}
	cfg := &frpc.CfgModel{
		Frpc: cc,
	}

	//glog.Infof("menu: %+v", cfg)
	//proxy := v1.TypedProxyConfig{
	//	Type: "tcp",
	//}
	//v1.TCPProxyConfig{
	//	v1.ProxyBaseConfig{Type: "tcp"},
	//}
	return cfg
}
