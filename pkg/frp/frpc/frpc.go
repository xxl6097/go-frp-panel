package frpc

import (
	_ "github.com/xxl6097/go-frp-panel/assets/frpc"
	"github.com/xxl6097/go-frp-panel/internal/frpc"
	"github.com/xxl6097/go-frp-panel/pkg/comm/iface"
	"github.com/xxl6097/go-service/pkg/gs/igs"
)

func NewFrpc(i igs.Service) (iface.IFrpc, error) {
	return frpc.New(i)
}
