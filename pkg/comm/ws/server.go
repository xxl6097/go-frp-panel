package ws

import (
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
	clients  map[string]map[string]*iface.WSSession
	callback iface.OnWebSocketCallBack
}

func (this *FrpWebSocket) GetList(key string) []*iface.WSSession {
	if this.clients != nil {
		v, ok := this.clients[key]
		if ok {
			var list []*iface.WSSession
			for _, vv := range v {
				list = append(list, vv)
			}
			return list
		}
	}
	return nil
}

func (this *FrpWebSocket) GetListSize() map[string]int {
	if this.clients != nil {
		var list = make(map[string]int)
		for id, c := range this.clients {
			list[id] = len(c)
		}
		return list
	}
	return nil
}

func (this *FrpWebSocket) GetDetail(id, key string) *iface.WSSession {
	if this.clients != nil {
		v, ok := this.clients[id]
		if ok {
			return v[key]
		}
	}
	return nil
}

func (this *FrpWebSocket) SetWebSocket(back iface.OnWebSocketCallBack) {
	this.callback = back
}

func (this *FrpWebSocket) Send(id string, messageType int, payload []byte) error {
	v, ok := this.clients[id]
	if ok && v != nil && len(v) > 0 {
		for _, conn := range v {
			err := conn.Conn.WriteMessage(messageType, payload)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (this *FrpWebSocket) SendByKey(id, secKey string, messageType int, payload []byte) error {
	v, ok := this.clients[id]
	if ok && v != nil && len(v) > 0 {
		conn, okk := v[secKey]
		if okk && conn != nil {
			err := conn.Conn.WriteMessage(messageType, payload)
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
			pointAddress := fmt.Sprintf("%p", ws)
			glog.Errorf("websocket断开:%v,address:%v,messageType:%v,err:%v", ws.RemoteAddr().String(), pointAddress, messageType, err)
			break
		} else {
			glog.Printf("Received:%+v %+v %+v\n", ws.RemoteAddr().String(), messageType, message)
			if this.callback != nil {
				this.callback.OnServerWebSocketMessageReceive(messageType, message)
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
	osType := r.Header.Get("OsType")
	id := r.Header.Get("ClientID")
	localMacAddress := r.Header.Get("LocalMacAddress")
	localIpv4 := r.Header.Get("LocalIpv4")
	secKey := r.Header.Get("Sec-Websocket-Key")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		glog.Errorf("ClientID空：%+v", r)
		return
	}
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
	glog.Warn("websocket客户端连接成功", secKey, localMacAddress, localIpv4, id)
	childMap := this.clients[id]
	defer func() {
		session := childMap[secKey]
		if this.callback != nil {
			this.callback.OnServerWebSocketDisconnect(session)
		}
		delete(childMap, secKey)
	}()
	session := iface.WSSession{Conn: ws, SecKey: secKey, DevMac: localMacAddress, DevIp: localIpv4, ID: id, OsType: osType}
	if childMap == nil {
		childMap = make(map[string]*iface.WSSession)
		childMap[secKey] = &session
		this.clients[id] = childMap
	} else {
		childMap[secKey] = &session
	}
	if this.callback != nil {
		this.callback.OnServerWebSocketNewConnection(&session)
	}
	this.onMessageRecv(ws, r)
}

func NewWebSocket() iface.IWebSocket {
	return &FrpWebSocket{clients: make(map[string]map[string]*iface.WSSession)}
}
