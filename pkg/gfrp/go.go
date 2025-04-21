package gfrp

import (
	v1 "github.com/fatedier/frp/pkg/config/v1"
	"github.com/xxl6097/go-frp-panel/internal/comm/iface"
	"github.com/xxl6097/go-frp-panel/internal/frps"
	"github.com/xxl6097/go-service/gservice/gore"
)

func NewFrps(cfg *v1.ServerConfig, install gore.IGService) (iface.IFrps, error) {
	return frps.New(cfg, install)
}
