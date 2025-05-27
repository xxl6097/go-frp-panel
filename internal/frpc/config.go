package frpc

import (
	v1 "github.com/fatedier/frp/pkg/config/v1"
	"github.com/xxl6097/glog/glog"
	"github.com/xxl6097/go-frp-panel/pkg"
	"github.com/xxl6097/go-frp-panel/pkg/utils"
	"github.com/xxl6097/go-service/pkg/ukey"
)

var cfgData *CfgModel
var cfgBytes []byte

type CfgModel struct {
	Frpc v1.ClientConfig `json:"frpc"`
}

func (this *CfgModel) Bytes() []byte {
	return utils.ObjectToTomlText(this)
}
func load() error {
	defer glog.Flush()
	byteArray, err := ukey.Load()
	if err != nil {
		return err
	}
	cfgBytes = byteArray
	var cfg v1.ClientConfig
	err = utils.TomlTextToObject(byteArray, &cfg)
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
