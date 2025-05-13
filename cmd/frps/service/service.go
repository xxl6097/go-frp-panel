package service

import (
	"encoding/json"
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
	"github.com/xxl6097/go-service/gservice"
	"github.com/xxl6097/go-service/gservice/gore"
	"github.com/xxl6097/go-service/gservice/ukey"
	utils2 "github.com/xxl6097/go-service/gservice/utils"
	"os"
	"path/filepath"
)

type Service struct {
	ifrps     iface.IFrps
	webServer *v1.WebServerConfig
}

func Bootstrap() {
	svr := &Service{}
	err := gservice.Run(svr)
	if err != nil {
		glog.Error("程序启动出错了", err)
	}
	if svr.webServer != nil {
		glog.Infof("\n登录地址：http://%s:%d\n用户信息：%s/%s", utils.GetLocalIp(), svr.webServer.Port, svr.webServer.User, svr.webServer.Password)
	}
	glog.Warnf("OnFinish %+v", svr.webServer)
	glog.Println("服务程序启动成功", os.Getegid())
}

func (s *Service) OnInit() *service.Config {
	return &service.Config{
		Name:        pkg.AppName,
		DisplayName: pkg.DisplayName,
		Description: pkg.Description,
	}
}

func (s *Service) OnStop(ss service.Service) {
	s.ifrps.Close()
}

func (s *Service) ShutDown(ss service.Service) {
	s.ifrps.Close()
}

func (s *Service) OnVersion() string {
	fmt.Println(string(ukey.GetBuffer()))
	//这里需要打印config中buffer原始信息
	ver := fmt.Sprintf("frps version:%s", version.Full())
	fmt.Println("GetCrossPlatformDataDir", glog.GetCrossPlatformDataDir())
	pkg.Version()
	return ver
}

func (this *Service) OnRun(i gore.IGService) error {
	frps.Assert()
	glog.Printf("启动 %s %s\n", pkg.AppName, pkg.AppVersion)
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
	cfgConfig := &v1.ServerConfig{}
	err = json.Unmarshal(content, cfgConfig)
	if err != nil {
		glog.Error(err)
		return err
	}

	svv, err := frps2.NewFrps(cfgConfig, i)
	if err != nil {
		glog.Printf("启动 %s %s 失败:%v\n%v", pkg.AppName, pkg.AppVersion, err, conf)
		return err
	}
	this.ifrps = svv
	svv.Run()
	return err
}

func (this *Service) GetAny(binDir string) any {
	a := this.menu()
	if a == nil {
		return nil
	}
	this.webServer = &a.Frps.WebServer
	return a
}

//func (s *Service) OnUpgrade(oldBinPath string, newFileUrlOrLocalPath string) (bool, []string) {
//	//1、读取老文件特征数据；
//	//2、下载新文件
//	//3、替换新文件特征数据
//	//4、数据写到安装目录地址（oldBinPath）
//	cfgBufferBytes := ukey.GetCfgBufferFromFile(oldBinPath)
//	if cfgBufferBytes == nil {
//		return false, nil
//	}
//	glog.Debug("获取配置数据成功", len(cfgBufferBytes))
//	if _, err := os.Stat(oldBinPath); !os.IsNotExist(err) {
//		err := os.Remove(oldBinPath)
//		if err != nil {
//			glog.Error("删除失败", oldBinPath)
//			return false, nil
//		}
//	}
//	var newFilePath string
//	if utils2.FileExists(newFileUrlOrLocalPath) {
//		newFilePath = newFileUrlOrLocalPath
//	} else if utils2.IsURL(newFileUrlOrLocalPath) {
//		glog.Debug("下载文件", newFileUrlOrLocalPath)
//		temp, err := utils.DownLoad(newFileUrlOrLocalPath)
//		if err != nil {
//			glog.Error("下载失败", err)
//			return false, nil
//		}
//		glog.Debug("下载成功.", temp)
//		newFilePath = temp
//	}
//	if newFilePath != "" {
//		oldBuffer := ukey.GetBuffer()
//		err := utils.GenerateBin(newFilePath, oldBinPath, oldBuffer, cfgBufferBytes)
//		if err != nil {
//			glog.Error("签名错误：", err)
//			return false, nil
//		}
//		return true, nil
//	}
//	return false, nil
//}
//
//func (s *Service) OnInstall(binPath string) (bool, []string) {
//	if frps.IsInit() == nil {
//		return false, nil
//	}
//	cfg := s.menu()
//	//cfg.Frps.Complete()
//	newBufferBytes, err := ukey.GenConfig(cfg, false)
//	if err != nil {
//		panic(fmt.Errorf("构建签名信息错误: %v", err))
//	}
//	//glog.Printf("--->%s\n", string(newBufferBytes))
//	currentBinPath, err := os.Executable()
//	if err != nil {
//		glog.Fatal("os.Executable() error", err)
//	}
//	if utils2.FileExists(binPath) {
//		utils.Delete(binPath, "旧运行文件")
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
//	return false, nil
//}

func (this *Service) menu() *frps.CfgModel {
	frps.Assert()
	cfg := frps.GetCfgModel()
	if cfg != nil {
		this.webServer = &cfg.Frps.WebServer
		if cfg.Frps.BindPort > 0 {
			return nil
		}
	}

	bindPort := utils2.InputIntDefault("Frps绑定端口(默认:6000):", 6000)
	adminPort := utils2.InputIntDefault("管理后台端口(默认:6500):", 6500)
	addr := utils2.InputStringEmpty("管理后台地址(默认:0.0.0.0)：", "0.0.0.0")
	username := utils2.InputStringEmpty("管理后台用户名(默认:admin)：", "admin")
	password := utils2.InputString("管理后台密码：")
	temp := glog.GetCrossPlatformDataDir("frps", "log")
	cm := &frps.CfgModel{
		Frps: v1.ServerConfig{
			BindPort: bindPort,
			WebServer: v1.WebServerConfig{
				User:     username,
				Password: password,
				Port:     adminPort,
				Addr:     addr,
			},
			Log: v1.LogConfig{
				To:      filepath.Join(temp, "frps.log"),
				MaxDays: 3,
				Level:   "error",
			},
		},
	}
	return cm
}
