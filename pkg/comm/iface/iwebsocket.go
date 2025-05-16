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
	FrpID  string          `json:"frpId"`
}

// Message 发送JSON消息
type Message[T any] struct {
	Action string `json:"action"`
	DevMac string `json:"devMac"`
	SseID  string `json:"sseId"`
	DevIp  string `json:"devIp"`
	FrpId  string `json:"frpId"`
	SecKey string `json:"secKey"`
	Data   T      `json:"data"`
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
	SendByKey(frpId, webSocketId string, messageType int, payload []byte) error
	GetListSize() map[string]int
	GetList(string) []*WSSession
	GetDetail(string, string) *WSSession
}
