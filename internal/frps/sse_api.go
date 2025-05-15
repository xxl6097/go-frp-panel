package frps

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/xxl6097/glog/glog"
	iface2 "github.com/xxl6097/go-frp-panel/pkg/comm/iface"
	"strings"
)

var (
	CLIENT_LIST   = "list"
	CLIENT_DETAIL = "detail"
	CLIENT_INFO   = "client-info"
)

func (this *frps) OnSseDisconnect(client *iface2.SSEClient) {
}

func (this *frps) OnSseNewConnection(client *iface2.SSEClient) {
	if strings.HasPrefix(client.SseId, CLIENT_LIST) {
		//列表
		if this.webSocketApi != nil {
			list := this.webSocketApi.GetListSize()
			if list != nil && len(list) > 0 {
				this.sseApi.BroadcastByType(CLIENT_LIST, iface2.SSEEvent{
					Event: CLIENT_LIST,
					Payload: map[string]interface{}{
						"id":   client.SseId,
						"data": list,
					},
				})
			}
		}
	} else if strings.HasPrefix(client.SseId, CLIENT_DETAIL) {
		//详情
		if this.webSocketApi != nil && this.sseApi != nil {
			this.getClientInfo(client)
			detail := this.webSocketApi.GetDetail(client.FrpID, client.SecKey)
			if detail != nil {
				eve := iface2.SSEEvent{
					Event: CLIENT_DETAIL,
					Payload: map[string]interface{}{
						"id":   client.SseId,
						"data": detail,
					},
				}
				err := this.sseApi.Send(client, eve)
				if err != nil {
					glog.Errorf("Send error: %s", err)
				} else {
					//glog.Infof("Send success %v", client.SseId)
				}
			} else {
				glog.Errorf("No Detail: %v", client.SseId)
			}
		}

	}
}

func (this *frps) getClientInfo(client *iface2.SSEClient) {
	if this.webSocketApi != nil {
		msg := iface2.Message[any]{
			Action: "clientInfo",
			SseID:  client.SseId,
		}
		b, e := json.Marshal(msg)
		if e != nil {
			glog.Errorf("getClientInfo error: %v", e)
			return
		}
		e = this.webSocketApi.SendByKey(client.FrpID, client.SecKey, websocket.TextMessage, b)
		if e != nil {
			glog.Errorf("getClientInfo error: %v", e)
		}
	}
}
