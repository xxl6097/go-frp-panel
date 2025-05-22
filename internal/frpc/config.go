package frpc

import (
	"encoding/json"
	v1 "github.com/fatedier/frp/pkg/config/v1"
	"github.com/xxl6097/glog/glog"
	"github.com/xxl6097/go-frp-panel/pkg"
	"github.com/xxl6097/go-service/gservice/ukey"
)

var cfgData *CfgModel
var cfgBytes []byte

type CfgModel struct {
	Frpc v1.ClientConfig `json:"frpc"`
}

func load() error {
	defer glog.Flush()
	byteArray, err := ukey.Load()
	if err != nil {
		return err
	}
	cfgBytes = byteArray
	var cfg v1.ClientConfig
	err = json.Unmarshal(cfgBytes, &cfg)
	if err != nil {
		glog.Println("ClientConfig解析错误", err)
		return err
	}
	cfgData = &CfgModel{
		Frpc: cfg,
	}
	pkg.Version()
	return nil
}

func GetCfgModel() *CfgModel {
	if cfgData == nil {
		err := load()
		if err != nil {
			return nil
		}
	}
	return cfgData
}

func SetCfgModel(c *CfgModel) {
	cfgData = c
}
