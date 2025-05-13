package ws

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/xxl6097/glog/glog"
	"github.com/xxl6097/go-frp-panel/pkg/comm/iface"
	"net/http"
)

// upgrader 用于升级HTTP连接到WebSocket连接
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// 允许所有的源，生产环境下需要配置跨域策略
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type FrpWebSocket struct {
	clients map[string]map[string]*websocket.Conn
}

func (this *FrpWebSocket) Send(id string, payload []byte) error {
	v, ok := this.clients[id]
	if ok && v != nil && len(v) > 0 {
		for _, conn := range v {
			err := conn.WriteMessage(websocket.TextMessage, payload)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (this *FrpWebSocket) onMessageRecv(ws *websocket.Conn, r *http.Request) {
	for {
		// 读取消息
		messageType, message, err := ws.ReadMessage()
		if err != nil {
			//delete(this.clients, ws.RemoteAddr().String())
			pointAddress := fmt.Sprintf("%p", ws)
			glog.Errorf("websocket断开:%v,address:%v,messageType:%v,err:%v", ws.RemoteAddr().String(), pointAddress, messageType, err)
			break
		} else {
			glog.Printf("Received:%+v %+v %+v\n", ws.RemoteAddr().String(), messageType, message)
			//this.clients[ws.RemoteAddr().String()] = ws
			if messageType == websocket.TextMessage {
				var msg Message[any]
				e := json.Unmarshal(message, &msg)
				if e == nil {
					glog.Warnf("Received: %+v", msg)
				}
			}
		}
	}
}

// HandleConnections 处理WebSocket连接
func (this *FrpWebSocket) HandleConnections(w http.ResponseWriter, r *http.Request) {
	glog.Debugf("WebSocket请求：%+v", r)
	//for key, values := range r.Header {
	//	fmt.Printf("Header[%q] = %q\n", key, values)
	//	// 若需处理单个值，可以遍历 values 切片
	//	//for _, v := range values {
	//	//	fmt.Printf("Value: %s\n", v)
	//	//}
	//}
	secKey := r.Header.Get("Sec-Websocket-Key")
	if secKey == "" {
		w.WriteHeader(http.StatusBadRequest)
		glog.Errorf("Sec-Websocket-Key空：%+v", r)
		return
	}
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		glog.Error("websocket连接错误", ws, err)
		return
	}
	defer ws.Close()
	pointAddress := fmt.Sprintf("%p", ws)
	glog.Warn("websocket客户端连接成功", secKey, pointAddress)
	childMap := this.clients[secKey]
	defer delete(childMap, pointAddress)
	if childMap == nil {
		childMap = make(map[string]*websocket.Conn)
		childMap[pointAddress] = ws
		this.clients[secKey] = childMap
	} else {
		childMap[pointAddress] = ws
	}
	this.onMessageRecv(ws, r)
}

func NewWebSocket() iface.IWebSocket {
	return &FrpWebSocket{clients: make(map[string]map[string]*websocket.Conn)}
}
