package service

import (
	"fmt"
	v1 "github.com/fatedier/frp/pkg/config/v1"
	"github.com/fatedier/frp/pkg/util/version"
	"github.com/kardianos/service"
	"github.com/xxl6097/glog/glog"
	"github.com/xxl6097/go-frp-panel/internal/frpc"
	"github.com/xxl6097/go-frp-panel/pkg"
	"github.com/xxl6097/go-frp-panel/pkg/frp"
	frpc2 "github.com/xxl6097/go-frp-panel/pkg/frp/frpc"
	"github.com/xxl6097/go-frp-panel/pkg/utils"
	"github.com/xxl6097/go-service/pkg/gs"
	"github.com/xxl6097/go-service/pkg/gs/igs"
	"github.com/xxl6097/go-service/pkg/ukey"
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
	fmt.Println(string(ukey.GetBuffer()))
	ver := fmt.Sprintf("frpc version:%s", version.Full())
	pkg.Version()
	return ver
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
	//dir, err := frp.GetFrpcTomlDirByDir(binDir)
	//if err != nil {
	//	return nil
	//}
	//cfgPath := filepath.Join(dir, frp.GetFrpcMainTomlFileName())
	//if err := frp.WriteFrpToml(cfgPath, cfg.Frpc); err != nil {
	//	glog.Warnf("write content to frpc config file error: %v", err)
	//	return nil
	//} else {
	//	glog.Infof("write content to frpc config file success %s", cfgPath)
	//}
	return cfg.Bytes()
}

//func (this *Service) menu1() *frpc.CfgModel {
//	var bindAddr, userName, token, id, apiPort, authorization string
//	var bindPort int
//	err := frpc.IsInit()
//	c := frpc.GetCfgModel()
//	//glog.Infof("Frpc: %+v", c.Frpc)
//	//glog.Infof("Cfg: %+v", c.Cfg)
//	//glog.Error(err)
//	if err != nil || c == nil {
//		bindAddr = utils2.InputString("Frps服务器地址:")
//		bindPort = utils2.InputInt("Frps服务器绑定端口:")
//		userName = utils2.InputStringEmpty("请输入用户名(admin):", "admin")
//		id = utils2.InputString("请输入ID：")
//		token = utils2.InputString("请输入密钥：")
//	} else {
//		bindAddr = c.Frpc.ServerAddr
//		bindPort = c.Frpc.ServerPort
//		userName = c.Frpc.User
//		token = c.Frpc.Metadatas["token"]
//		id = c.Frpc.Metadatas["id"]
//		apiPort = c.Frpc.Metadatas["apiPort"]
//	}
//	webServer := &v1.WebServerConfig{}
//	if c != nil && c.Cfg != nil {
//		if c.Cfg.WebServer != nil && c.Cfg.WebServer.Port != 0 && c.Cfg.WebServer.Addr != "" && c.Cfg.WebServer.User != "" && c.Cfg.WebServer.Password != "" {
//			webServer = c.Cfg.WebServer
//		}
//		authorization = c.Cfg.Authorization
//	}
//	if webServer.Addr == "" {
//		webServer.Addr = "0.0.0.0"
//	}
//	if webServer.Port == 0 {
//		webServer.Port = utils2.InputIntDefault("管理后台端口(6400)", 6400)
//	}
//	if webServer.User == "" {
//		webServer.User = utils2.InputStringEmpty("管理后台用户名(admin):", "admin")
//	}
//	if webServer.Password == "" {
//		webServer.Password = utils2.InputString("管理后台密码：")
//	}
//
//	temp := glog.GetCrossPlatformDataDir("frpc", "log")
//	ccc := v1.ClientCommonConfig{
//		ServerAddr: bindAddr,
//		ServerPort: bindPort,
//		User:       userName,
//		Metadatas:  frp.GetMetadatas(token, id, apiPort, authorization), //frpc服务安装 菜单显示 写入文件
//		Log: v1.LogConfig{
//			To:      filepath.Join(temp, "frpc.log"),
//			MaxDays: 7,
//			Level:   "error",
//		},
//		WebServer: *webServer,
//	}
//
//	var proxies []v1.TypedProxyConfig
//	if c != nil && c.Cfg != nil && c.Cfg.Proxy != nil && comm.HasProxyes(c.Cfg.Proxy) {
//		proxies = append(proxies, *c.Cfg.Proxy)
//	}
//
//	cc := v1.ClientConfig{
//		ClientCommonConfig: ccc,
//		Proxies:            proxies,
//	}
//	cfg := &frpc.CfgModel{
//		Frpc: cc,
//	}
//
//	//glog.Infof("menu: %+v", cfg)
//	//proxy := v1.TypedProxyConfig{
//	//	Type: "tcp",
//	//}
//	//v1.TCPProxyConfig{
//	//	v1.ProxyBaseConfig{Type: "tcp"},
//	//}
//	return cfg
//}

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
