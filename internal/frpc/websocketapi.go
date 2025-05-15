package frpc

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/xxl6097/glog/glog"
	"github.com/xxl6097/go-frp-panel/pkg/comm/iface"
	"github.com/xxl6097/go-frp-panel/pkg/comm/ws"
	"github.com/xxl6097/go-frp-panel/pkg/frp"
	"github.com/xxl6097/go-frp-panel/pkg/utils"
	"net/http"
	"os"
	"path/filepath"
)

func (this *frpc) onWebSocketMessageHandle(data []byte) {
	if data != nil {
		var msg iface.Message[any]
		err := json.Unmarshal(data, &msg)
		if err != nil {
			glog.Error(err)
			return
		}
		glog.Debugf("msg:%+v", msg)
		switch msg.Action {
		case ws.CLIENT_INFO:
			this.getClientInfo(msg.SseID)
			break
		case ws.CONFIG_LIST:
			this.getConfigs(msg.SseID)
			break
		case ws.TOML_UPGRADE:
			this.recvTomlUpgrade(msg.Data)
			break
		case ws.REBOOT:
			if this.install == nil {
				return
			}
			err = this.install.Restart()
			if err != nil {
				glog.Error(err)
			}
			break
		case ws.UNINSTALL:
			if this.install == nil {
				return
			}
			err = this.install.Uninstall()
			if err != nil {
				glog.Error(err)
			}
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
		Action: ws.CLIENT_INFO,
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

func (this *frpc) recvTomlUpgrade(data any) {
	body, ok := data.(map[string]interface{})
	if !ok {
		glog.Error("data is not Toml")
		return
	}
	err := this.upgradeTomlContent(body["label"].(string), body["content"].(string))
	if err != nil {
		glog.Error(err)
		return
	}
	glog.Debug("ConfigUpgrade sucess")
}

func (this *frpc) getConfigs(sseId string) {
	cfgDir, err := frp.GetFrpcTomlDir()
	if err != nil {
		glog.Error(err)
		return
	}
	files, err := os.ReadDir(cfgDir)
	if err != nil {
		glog.Error(err)
		return
	}
	type Option struct {
		Label   string `json:"label"`
		Value   string `json:"value"`
		Content string `json:"content"`
	}
	var list []Option
	for _, file := range files {
		fileName := file.Name()
		buffer, e := utils.Read(filepath.Join(cfgDir, fileName))
		if e == nil {
			item := Option{
				Label:   fileName,
				Value:   fileName,
				Content: string(buffer),
			}
			list = append(list, item)
		}
	}
	msg := iface.Message[[]Option]{
		Action: ws.CONFIG_LIST,
		Data:   list,
		SseID:  sseId,
	}
	err = ws.GetClientInstance().SendJSON(msg)
	if err != nil {
		glog.Error(err)
	} else {
		glog.Debugf("getClients sucess")
	}
}
