package frps

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/xxl6097/glog/glog"
	"github.com/xxl6097/go-frp-panel/pkg/comm"
	iface2 "github.com/xxl6097/go-frp-panel/pkg/comm/iface"
	"github.com/xxl6097/go-frp-panel/pkg/comm/ws"
	"io"
	"net/http"
)

func (this *frps) OnServerWebSocketMessageReceive(messageType int, payload []byte) {
	if payload != nil {
		glog.Debugf("ws msg %s", string(payload))
		var msg iface2.Message[any]
		err := json.Unmarshal(payload, &msg)
		if err != nil {
			glog.Error(err)
			return
		}
		//glog.Debugf("ws recv:%+v", msg)
		this.recvClientInfo(msg.SseID, msg.Action, msg.Data)
	}
}

func (this *frps) OnServerWebSocketDisconnect(session *iface2.WSSession) {
	if this.sseApi != nil {
		eve := iface2.SSEEvent{
			Event:   ws.DISCONNECT,
			Payload: session,
		}
		this.sseApi.Broadcast(eve)
	}
}

func (this *frps) OnServerWebSocketNewConnection(session *iface2.WSSession) {
}

func (this *frps) recvClientInfo(sseId, event string, data any) {
	if data == nil {
		glog.Error("data is nil")
		return
	}
	if this.sseApi != nil {
		eve := iface2.SSEEvent{
			Event:   event,
			Payload: data,
		}
		okk := this.sseApi.BroadcastTo(sseId, eve)
		if !okk {
			glog.Errorf("Send error: %s sseId:%v", event, sseId)
		} else {
			glog.Infof("Send success %s %v", event, sseId)
		}
	}
}

func (this *frps) apiClientCMD(w http.ResponseWriter, r *http.Request) {
	res, f := comm.Response(r)
	defer f(w)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		glog.Error("body读取失败", err)
		res.Err(err)
		return
	}
	if body == nil {
		msg := "body is nil"
		glog.Error(msg)
		res.Err(fmt.Errorf(msg))
		return
	}
	var msg iface2.Message[any]
	err = json.Unmarshal(body, &msg)
	if err != nil {
		glog.Error("解析Json对象失败", err)
		res.Err(err)
		return
	}
	glog.Debugf("body:%s", string(body))
	if this.webSocketApi == nil {
		res.Error(fmt.Sprintf("webSocketApi is nil"))
		return
	}
	e := this.webSocketApi.SendByKey(msg.FrpId, msg.SecKey, websocket.TextMessage, body)
	if e != nil {
		glog.Errorf("apiClientCMD error: %v", e)
		res.Err(e)
	} else {
		glog.Infof("Send success %v", msg.FrpId)
		res.Ok("执行成功～")
	}
}

//func (this *frps) apiClientCreate(w http.ResponseWriter, r *http.Request) {
//	res, f := comm.Response(r)
//	defer f(w)
//
//	body, err := utils.GetDataByJson[struct {
//		Label   string `json:"label"`
//		Content string `json:"content"`
//		FrpId   string `json:"frpId"`
//		SecKey  string `json:"secKey"`
//		SseId   string `json:"sseId"`
//	}](r)
//	if err != nil {
//		res.Err(err)
//		glog.Error(res.Msg)
//		return
//	}
//	if body == nil {
//		res.Error("body is empty")
//		glog.Error(res.Msg)
//		return
//	}
//	if body.Label == "" {
//		res.Error("名称不能为空～")
//		glog.Error(res.Msg)
//		return
//	}
//	if body.Content == "" {
//		res.Error("toml配置空")
//		glog.Error(res.Msg)
//		return
//	}
//
//	if filepath.Ext(body.Label) != ".toml" {
//		body.Label = fmt.Sprintf("%s.toml", body.Label)
//	}
//
//	if this.webSocketApi == nil {
//		res.Error(fmt.Sprintf("webSocketApi is nil"))
//	}
//
//	type Option struct {
//		Label   string `json:"label"`
//		Content string `json:"content"`
//	}
//	if this.webSocketApi != nil {
//		msg := iface2.Message[Option]{
//			Action: ws.CLIENT_NEW,
//			Data: Option{
//				Label:   body.Label,
//				Content: body.Content,
//			},
//		}
//		b, e := json.Marshal(msg)
//		if e != nil {
//			glog.Errorf("getClientInfo error: %v", e)
//			return
//		}
//		e = this.webSocketApi.SendByKey(body.FrpId, body.SecKey, websocket.TextMessage, b)
//		if e != nil {
//			glog.Errorf("getClientInfo error: %v", e)
//		} else {
//			glog.Infof("Send success %v", body.FrpId)
//			// 请求一次客户端列表
//			this.getConfigs(body.SseId, body.FrpId, body.SecKey)
//		}
//	}
//}
