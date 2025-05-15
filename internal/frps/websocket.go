package frps

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/xxl6097/glog/glog"
	"github.com/xxl6097/go-frp-panel/pkg/comm"
	iface2 "github.com/xxl6097/go-frp-panel/pkg/comm/iface"
	"github.com/xxl6097/go-frp-panel/pkg/utils"
	"net/http"
)

func (this *frps) OnServerWebSocketMessageReceive(messageType int, payload []byte) {
	if payload != nil {
		var msg iface2.Message[any]
		err := json.Unmarshal(payload, &msg)
		if err != nil {
			glog.Error(err)
			return
		}
		switch msg.Action {
		case "clientInfo":
			this.recvClientInfo(msg.SseID, msg.Data)
			break
		}
	}
}

func (this *frps) OnServerWebSocketDisconnect(session *iface2.WSSession) {
}

func (this *frps) OnServerWebSocketNewConnection(session *iface2.WSSession) {
}

func (this *frps) recvClientInfo(sseId string, data any) {
	if data == nil {
		glog.Error("data is nil")
		return
	}
	body, ok := data.(string)
	if !ok {
		glog.Error("data is not []byte")
		return
	}
	if this.sseApi != nil {
		eve := iface2.SSEEvent{
			Event:   CLIENT_INFO,
			Payload: body,
		}
		okk := this.sseApi.BroadcastTo(sseId, eve)
		if !okk {
			glog.Errorf("Send error: %v", okk)
		} else {
			glog.Infof("Send success %s %v", CLIENT_INFO, sseId)
		}
	}
}

func (this *frps) apiClientConfigUpgrade(w http.ResponseWriter, r *http.Request) {
	res, f := comm.Response(r)
	defer f(w)
	body, err := utils.GetDataByJson[struct {
		Toml   string `json:"toml"`
		FrpId  string `json:"frpId"`
		SecKey string `json:"secKey"`
	}](r)
	if err != nil {
		glog.Error("解析Json对象失败", err)
		return
	}
	if body == nil {
		msg := "json对象nil"
		glog.Error(msg)
		http.Error(w, "json对象nil", http.StatusInternalServerError)
		return
	}
	glog.Debugf("body:%+v", body)
	if this.webSocketApi == nil {
		res.Error(fmt.Sprintf("webSocketApi is nil"))
	}

	if this.webSocketApi != nil {
		msg := iface2.Message[string]{
			Action: "mainTomlUpgrade",
			Data:   body.Toml,
		}
		b, e := json.Marshal(msg)
		if e != nil {
			glog.Errorf("getClientInfo error: %v", e)
			return
		}
		e = this.webSocketApi.SendByKey(body.FrpId, body.SecKey, websocket.TextMessage, b)
		if e != nil {
			glog.Errorf("getClientInfo error: %v", e)
		} else {
			glog.Infof("Send success %v", body.FrpId)
		}
	}
}

func (this *frps) apiClientCMD(w http.ResponseWriter, r *http.Request) {
	res, f := comm.Response(r)
	defer f(w)
	body, err := utils.GetDataByJson[struct {
		Cmd    string `json:"cmd"`
		FrpId  string `json:"frpId"`
		SecKey string `json:"secKey"`
	}](r)
	if err != nil {
		glog.Error("解析Json对象失败", err)
		res.Err(err)
		return
	}
	if body == nil {
		msg := "json对象nil"
		glog.Error(msg)
		res.Err(fmt.Errorf(msg))
		return
	}
	glog.Debugf("body:%+v", body)
	if this.webSocketApi == nil {
		res.Error(fmt.Sprintf("webSocketApi is nil"))
		return
	}

	msg := iface2.Message[string]{
		Action: body.Cmd,
	}
	b, e := json.Marshal(msg)
	if e != nil {
		glog.Errorf("apiClientCMD error: %v", e)
		res.Err(e)
		return
	}
	e = this.webSocketApi.SendByKey(body.FrpId, body.SecKey, websocket.TextMessage, b)
	if e != nil {
		glog.Errorf("apiClientCMD error: %v", e)
		res.Err(e)
	} else {
		glog.Infof("Send success %v", body.FrpId)
		res.Ok("执行成功～")
	}
}
