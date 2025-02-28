package service

import (
	"encoding/json"
	"fmt"
	v1 "github.com/fatedier/frp/pkg/config/v1"
	"github.com/fatedier/frp/pkg/util/version"
	"github.com/kardianos/service"
	"github.com/xxl6097/glog/glog"
	"github.com/xxl6097/go-frp-panel/internal/comm/ukey"
	"github.com/xxl6097/go-frp-panel/internal/frps"
	"github.com/xxl6097/go-frp-panel/pkg"
	"github.com/xxl6097/go-frp-panel/pkg/utils"
	"github.com/xxl6097/go-service/gservice/gore"
	"os"
	"path/filepath"
)

type Service struct {
}

func (s Service) OnVersion() string {
	fmt.Println(string(ukey.GetBuffer()))
	//这里需要打印config中buffer原始信息
	b := frps.GetCfgBytes()
	if b != nil {
		glog.Println(string(b))
	}
	ver := fmt.Sprintf("frps version:%s", version.Full())
	pkg.Version()
	return ver
}

func (s Service) OnConfig() *service.Config {
	return &service.Config{
		Name:        pkg.AppName,
		DisplayName: pkg.DisplayName,
		Description: pkg.Description,
	}
}

func (s Service) OnInstall(binPath string) (bool, []string) {
	if frps.IsInit() == nil {
		return false, nil
	}
	cfg := s.menu()
	//cfg.Frps.Complete()
	newBufferBytes, err := ukey.GenConfig(cfg, false)
	if err != nil {
		panic(fmt.Errorf("构建签名信息错误: %v", err))
	}
	//glog.Printf("--->%s\n", string(newBufferBytes))
	currentBinPath, err := os.Executable()
	if err != nil {
		glog.Fatal("os.Executable() error", err)
	}
	if gore.FileExists(binPath) {
		utils.Delete(binPath, "旧运行文件")
	}
	//安装程序，需要对程序进行签名，那么需要传入两个参数：
	//1、最原始的key；
	//2、需写入的data
	buffer := ukey.GetBuffer()
	glog.Info("buffer大小", len(buffer))
	err = utils.GenerateBin(currentBinPath, binPath, buffer, newBufferBytes)
	if err != nil {
		glog.Fatal("签名错误：", err)
	}
	return false, nil
}

func (this Service) OnRun(i gore.Install) error {
	frps.Assert()
	glog.Printf("启动frps_%s\n", pkg.AppVersion)
	cfg := frps.GetCfgModel()
	if cfg == nil {
		return fmt.Errorf("程序配置文件未初始化")
	}
	conf := frps.GetCfgModel().Frps
	content, err := json.Marshal(conf)
	if err != nil {
		glog.Error(err)
		return err
	}
	svv, err := frps.NewFrps(content, i)
	if err != nil {
		glog.Error("启动frps失败", err)
		glog.Printf("启动frps_%s失败\n", pkg.AppVersion)
		glog.Println(conf)
		return err
	}
	svv.Run()
	return err
}

func (this *Service) menu() *frps.CfgModel {
	bindPort := gore.InputInt("请输入Frps服务器绑定端口：")
	adminPort := gore.InputInt("请输入管理后台端口：")
	addr := gore.InputStringEmpty("请输入管理后台地址(默认0.0.0.0)：", "0.0.0.0")
	username := gore.InputStringEmpty("请输入管理后台用户名(admin)：", "admin")
	password := gore.InputString("请输入管理后台密码：")
	temp := os.TempDir()
	temp = filepath.Join(temp, "frps", "logs")
	err := os.MkdirAll(temp, 0755)
	if err != nil {
		fmt.Println(err)
	}
	cfg := &frps.CfgModel{
		Frps: v1.ServerConfig{
			BindPort: bindPort,
			HTTPPlugins: []v1.HTTPPluginOptions{
				{
					Name: "frps-panel",
					Addr: fmt.Sprintf("%s:%d", addr, adminPort),
					Path: "/handler",
					Ops:  []string{"Login", "NewWorkConn", "NewUserConn", "NewProxy", "Ping"},
				},
			},
			WebServer: v1.WebServerConfig{
				User:     username,
				Password: password,
				Port:     adminPort,
				Addr:     addr,
			},
			Log: v1.LogConfig{
				To:      filepath.Join(temp, "frps.log"),
				MaxDays: 15,
			},
		},
	}
	return cfg
}
