package frpc

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/xxl6097/glog/glog"
	"github.com/xxl6097/go-frp-panel/pkg/comm/iface"
	"github.com/xxl6097/go-frp-panel/pkg/comm/ws"
	"net/http"
)

func (this *frpc) onWebSocketMessageHandle(data []byte) {
	if data != nil {
		var msg iface.Message[any]
		err := json.Unmarshal(data, &msg)
		if err != nil {
			glog.Error(err)
			return
		}
		switch msg.Action {
		case "clientInfo":
			this.getClientInfo(msg.SseID)
			break
		case "mainTomlUpgrade":
			this.recvMainTomlUpgrade(msg.Data)
			break
		}
	}
}

func (this *frpc) onWebSocketOpenHandle(conn *websocket.Conn, response *http.Response) {
	glog.Debugf("连接成功: %v,%v,Status:%v", conn.LocalAddr(), conn.RemoteAddr(), response.Status)
}

func (this *frpc) getClientInfo(sseId string) {
	body, err := this.getClientMainConfig()
	if err != nil {
		glog.Error(err)
		return
	}
	msg := iface.Message[string]{
		Action: "clientInfo",
		Data:   string(body),
		SseID:  sseId,
	}
	err = ws.GetClientInstance().SendJSON(msg)
	if err != nil {
		glog.Error(err)
	} else {
		glog.Debug("getClientInfo sucess")
	}
}

func (this *frpc) recvMainTomlUpgrade(data any) {
	body, ok := data.(string)
	if !ok {
		glog.Error("data is not []byte")
		return
	}
	err := this.upgradeMainTomlContent(body)
	if err != nil {
		glog.Error(err)
	}
	glog.Debug("ConfigUpgrade sucess", body)
}
