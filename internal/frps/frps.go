package frps

import (
	v1 "github.com/fatedier/frp/pkg/config/v1"
	_ "github.com/xxl6097/go-frp-panel/assets/frps"
	"github.com/xxl6097/go-frp-panel/pkg/comm/iface"
	"github.com/xxl6097/go-service/gservice/gore"
)

func NewFrps(cfg *v1.ServerConfig, install gore.IGService) (iface.IFrps, error) {
	return New(cfg, install)
}
