package frpc

import (
	"encoding/json"
	v1 "github.com/fatedier/frp/pkg/config/v1"
	"github.com/xxl6097/glog/glog"
	"github.com/xxl6097/go-frp-panel/pkg"
	"github.com/xxl6097/go-frp-panel/pkg/utils"
	"github.com/xxl6097/go-service/gservice/ukey"
	"os"
	"path/filepath"
)

var cfgData *CfgModel
var cfgBytes []byte

type CfgModel struct {
	Frpc v1.ClientCommonConfig `json:"frpc"`
	Data any                   `json:"data"`
}

func load() error {
	defer glog.Flush()
	temp := os.TempDir()
	glog.SetLogFile(filepath.Join(temp, "frpc", "logs"), "frpc.log")
	glog.SetCons(true)
	byteArray, err := ukey.Load()
	if err != nil {
		//glog.Error(err)
		return err
	}
	cfgBytes = byteArray
	//c := CfgModel{}
	c := ukey.ClientCommonConfig{}
	err = json.Unmarshal(cfgBytes, &c)
	if err != nil {
		glog.Println("cfgBytes解析错误", err)
		return err
	}
	cfgData = &CfgModel{Frpc: v1.ClientCommonConfig{
		ServerAddr: c.Addr,
		ServerPort: c.Port,
		User:       c.User,
		Metadatas: map[string]string{
			"token": c.Token,
		},
	}}
	//glog.Printf("%d 配置加载成功：%+v\n", os.Getpid(), cfgData)
	pkg.Version()
	return nil
}

func GetCfgBytes() []byte {
	load()
	return cfgBytes
}

func GetCfgModel() *CfgModel {
	return cfgData
}

func PrintCfg() {
	if cfgBytes != nil {
		glog.Println(string(cfgBytes))
	}
}

func IsInit() error {
	defer glog.Flush()
	err := load()
	if err != nil {
		//glog.Println(err)
		return err
	}
	return nil
}

func Assert() {
	if IsInit() != nil {
		if utils.IsMacOs() {
			return
		}
		os.Exit(0)
	}
}
