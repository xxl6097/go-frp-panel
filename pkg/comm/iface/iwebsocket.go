package iface

import (
	"github.com/gorilla/websocket"
	"net/http"
)

type IWebSocketMessage interface {
	onMessageRecv(ws *websocket.Conn, r *http.Request)
}
type IWebSocket interface {
	HandleConnections(http.ResponseWriter, *http.Request)
	Send(id string, payload []byte) error
}
