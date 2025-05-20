package frpc

import (
	_ "github.com/xxl6097/go-frp-panel/assets/frpc"
	"github.com/xxl6097/go-frp-panel/internal/frpc"
	"github.com/xxl6097/go-frp-panel/pkg/comm/iface"
	"github.com/xxl6097/go-service/gservice/gore"
)

func NewFrpc(i gore.IGService) (iface.IFrpc, error) {
	return frpc.New(i)
}
