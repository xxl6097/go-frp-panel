package gfrp

import (
	v1 "github.com/fatedier/frp/pkg/config/v1"
	"github.com/xxl6097/go-frp-panel/internal/frps"
	"github.com/xxl6097/go-frp-panel/pkg/comm/iface"
	"github.com/xxl6097/go-service/pkg/gs/igs"
)

func NewFrps(cfg *v1.ServerConfig, install igs.Service) (iface.IFrps, error) {
	return frps.New(cfg, install)
}
