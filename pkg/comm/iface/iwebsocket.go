package iface

import (
	"github.com/gorilla/websocket"
	"net/http"
)

type WSSession struct {
	Conn   *websocket.Conn `json:"-"`
	OsType string          `json:"osType"`
	SecKey string          `json:"secKey"`
	DevMac string          `json:"devMac"`
	DevIp  string          `json:"devIp"`
	ID     string          `json:"id"`
}

// Message 发送JSON消息
type Message struct {
	Action string `json:"action"`
	DevMac string `json:"devMac"`
	DevIp  string `json:"devIp"`
	Data   string `json:"data"`
}

// OnWebSocketCallBack 观察者接口，定义了更新方法
type OnWebSocketCallBack interface {
	OnServerWebSocketMessageReceive(messageType int, payload []byte)
	OnServerWebSocketDisconnect(*WSSession)
	OnServerWebSocketNewConnection(*WSSession)
}

type IWebSocket interface {
	SetWebSocket(OnWebSocketCallBack)
	HandleConnections(http.ResponseWriter, *http.Request)
	Send(id string, messageType int, payload []byte) error
	SendByKey(id, secKey string, messageType int, payload []byte) error
	GetListSize() map[string]int
	GetList(string) []*WSSession
	GetDetail(string, string) *WSSession
}
