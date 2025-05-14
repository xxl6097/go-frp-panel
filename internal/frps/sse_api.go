package frps

import (
	iface2 "github.com/xxl6097/go-frp-panel/pkg/comm/iface"
	"strings"
)

var (
	CLIENT_LIST   = "list"
	CLIENT_DETAIL = "detail"
)

func (this *frps) OnSseDisconnect(client *iface2.SSEClient) {
}

func (this *frps) OnSseNewConnection(client *iface2.SSEClient) {
	if strings.HasPrefix(client.SseId, CLIENT_LIST) {
		//列表
		if this.webSocketApi != nil {
			list := this.webSocketApi.GetListSize()
			if list != nil && len(list) > 0 {
				//this.sseApi.BroadcastByType(CLIENT_LIST, iface2.SSEEvent{
				//	Event: CLIENT_LIST,
				//	SseId:    client.SseId,
				//	Data:  list,
				//})
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
		if this.webSocketApi != nil {
			detail := this.webSocketApi.GetDetail(client.FrpID, client.SecKey)
			if detail != nil {
				this.sseApi.BroadcastByType(CLIENT_DETAIL, iface2.SSEEvent{
					Event: CLIENT_DETAIL,
					Payload: map[string]interface{}{
						"id":   client.SseId,
						"data": detail,
					},
				})
			}
		}

	}
}
