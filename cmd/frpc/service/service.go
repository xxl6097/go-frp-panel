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
	"github.com/xxl6097/go-service/gservice/gore"
	"github.com/xxl6097/go-service/gservice/ukey"
	utils2 "github.com/xxl6097/go-service/gservice/utils"
	"os"
	"path/filepath"
)

type Service struct {
}

func (s Service) OnInit() *service.Config {
	return &service.Config{
		Name:        pkg.AppName,
		DisplayName: pkg.DisplayName,
		Description: pkg.Description,
	}
}
func (s Service) OnVersion() string {
	fmt.Println(string(ukey.GetBuffer()))
	ver := fmt.Sprintf("frpc version:%s", version.Full())
	pkg.Version()
	return ver
}

func (this Service) OnRun(i gore.IGService) error {
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

func (this Service) GetAny(binDir string) any {
	cfg := this.menu()
	cfgPath := filepath.Join(binDir, "config.toml")
	if err := os.WriteFile(cfgPath, utils.ObjectToTomlText(cfg.Frpc), 0o600); err != nil {
		glog.Warnf("write content to frpc config file error: %v", err)
	} else {
		glog.Infof("write content to frpc config file success %s", cfgPath)
	}
	return cfg
}

//func (s Service) OnInstall(binPath string) (bool, []string) {
//	cfg := s.menu()
//	//cfg.Frpc.Complete()
//	newBufferBytes, err := ukey.GenConfig(cfg, false)
//	if err != nil {
//		panic(fmt.Errorf("构建签名信息错误: %v", err))
//	}
//	//glog.Printf("--->%s\n", string(newBufferBytes))
//	currentBinPath, err := os.Executable()
//	if err != nil {
//		glog.Fatal("os.Executable() error", err)
//	}
//	//安装程序，需要对程序进行签名，那么需要传入两个参数：
//	//1、最原始的key；
//	//2、需写入的data
//	buffer := ukey.GetBuffer()
//	glog.Info("buffer大小", len(buffer))
//	err = utils.GenerateBin(currentBinPath, binPath, buffer, newBufferBytes)
//	if err != nil {
//		glog.Fatal("签名错误：", err)
//	}
//	cfgPath := filepath.Join(filepath.Dir(binPath), "config.toml")
//	if err := os.WriteFile(cfgPath, utils.ObjectToTomlText(cfg.Frpc), 0o600); err != nil {
//		glog.Warnf("write content to frpc config file error: %v", err)
//	}
//	return false, nil
//}

func (this *Service) menu() *frpc.CfgModel {
	var bindAddr, userName, password, id string
	var bindPort int
	err := frpc.IsInit()
	c := frpc.GetCfgModel()
	glog.Infof("c: %+v", c)
	//glog.Error(err)
	if err != nil || c == nil {
		bindAddr = utils2.InputString("Frps服务器地址:")
		bindPort = utils2.InputInt("Frps服务器绑定端口:")
		userName = utils2.InputStringEmpty("请输入用户名(admin):", "admin")
		id = utils2.InputString("请输入ID：")
		password = utils2.InputString("请输入密钥：")
	} else {
		bindAddr = c.Frpc.ServerAddr
		bindPort = c.Frpc.ServerPort
		userName = c.Frpc.User
		password = c.Frpc.Metadatas["token"]
		password = c.Frpc.Metadatas["id"]
	}
	adminPort := utils2.InputIntDefault("管理后台端口(6400)", 6400)
	adminUser := utils2.InputStringEmpty("管理后台用户名(admin):", "admin")
	adminPass := utils2.InputString("管理后台密码：")
	temp := glog.GetCrossPlatformDataDir("frpc", "log")
	fCfg := v1.ClientCommonConfig{
		ServerAddr: bindAddr,
		ServerPort: bindPort,
		User:       userName,
		Metadatas: map[string]string{
			"token": password,
			"id":    id,
		},
		Log: v1.LogConfig{
			To:      filepath.Join(temp, "frpc.log"),
			MaxDays: 7,
			Level:   "error",
		},
		WebServer: v1.WebServerConfig{
			Addr:     "0.0.0.0",
			Port:     adminPort,
			User:     adminUser,
			Password: adminPass,
		},
	}
	cfg := &frpc.CfgModel{
		Frpc: fCfg,
	}

	glog.Infof("menu: %+v", cfg)
	//proxy := v1.TypedProxyConfig{
	//	Type: "tcp",
	//}
	return cfg
}
