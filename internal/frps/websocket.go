package frps

import iface2 "github.com/xxl6097/go-frp-panel/pkg/comm/iface"

func (this *frps) OnServerWebSocketMessageReceive(messageType int, payload []byte) {
}

func (this *frps) OnServerWebSocketDisconnect(session *iface2.WSSession) {
}

func (this *frps) OnServerWebSocketNewConnection(session *iface2.WSSession) {
}
