package iface

import "net/http"

// SSEClient 表示一个客户端连接
type SSEClient struct {
	SseId     string
	Send      chan SSEEvent
	CloseConn chan struct{}
	SecKey    string
	FrpID     string
}

// SSEEvent 表示一个SSE事件
type SSEEvent struct {
	Event   string      `json:"event,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
}
type SSEEvent1 struct {
	ID    string      `json:"id,omitempty"`
	Event string      `json:"event,omitempty"`
	Data  interface{} `json:"data"`
	Retry int         `json:"retry,omitempty"`
}

type OnSSECallBack interface {
	OnSseDisconnect(*SSEClient)
	OnSseNewConnection(*SSEClient)
}
type ISSE interface {
	http.Handler
	Broadcast(SSEEvent)
	BroadcastTo(string, SSEEvent) bool
	BroadcastByType(string, SSEEvent) bool
	Send(client *SSEClient, event SSEEvent) error
	GetClientCount() int
	GetClientIDs() []string
	CloseClient(clientID string) bool
	Start()
	SetSSECallBack(OnSSECallBack)
}
