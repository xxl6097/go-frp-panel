package frpc

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/xxl6097/glog/glog"
	"github.com/xxl6097/go-frp-panel/pkg/comm/iface"
	"github.com/xxl6097/go-frp-panel/pkg/comm/ws"
	"github.com/xxl6097/go-frp-panel/pkg/frp"
	"github.com/xxl6097/go-frp-panel/pkg/utils"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func (this *frpc) onWebSocketMessageHandle(data []byte) {
	if data != nil {
		//glog.Debugf("recv:%+v", string(data))
		var msg iface.Message[any]
		err := json.Unmarshal(data, &msg)
		if err != nil {
			glog.Error(err)
			return
		}
		glog.Debugf("recv msg:%+v", msg)
		switch msg.Action {
		case ws.CLIENT_REBOOT:
			if this.install == nil {
				return
			}
			err = this.install.Restart()
			if err != nil {
				glog.Error(err)
			}
			break
		case ws.CLIENT_UNINSTALL:
			if this.install == nil {
				return
			}
			err = this.install.Uninstall()
			if err != nil {
				glog.Error(err)
			}
			break
		default:
			//glog.Debugf("msg:%+v", msg)
			this.recvClientEvent(&msg)
			break
		}
	}
}

func (this *frpc) onWebSocketOpenHandle(conn *websocket.Conn, response *http.Response) {
	glog.Debugf("连接成功: %v,%v,Status:%v", conn.LocalAddr(), conn.RemoteAddr(), response.Status)
}

func (this *frpc) recvClientEvent(msg *iface.Message[any]) {
	if msg == nil {
		glog.Error("msg is nil")
		return
	}
	defer this.clientRefresh(msg)
	data := msg.Data
	if data == nil {
		glog.Error("data is nil")
		return
	}
	body, ok := data.(map[string]interface{})
	if !ok {
		glog.Errorf("body is err, the value is %+v", data)
		return
	}
	switch msg.Action {
	case ws.CLIENT_NEW:
		if body["content"] != nil && body["name"] != nil {
			err := this.clientNew(body["name"].(string), body["content"].(string))
			if err != nil {
				glog.Error(err)
				return
			}
			glog.Debug("new sucess")
		}
		break
	case ws.CLIENT_DELETE:
		if body["name"] != nil {
			err := this.clientDelete(body["name"].(string))
			if err != nil {
				glog.Error(err)
				return
			}
			glog.Debug("delete sucess")
		}
		break
	case ws.CLIENT_CHANGE:
		if body["content"] != nil && body["name"] != nil {
			err := this.upgradeTomlContent(body["name"].(string), body["content"].(string))
			if err != nil {
				glog.Error(err)
				return
			}
			glog.Debug("change sucess")
		}
		break
	}

}

func (this *frpc) getClientConfigs() (any, error) {
	cfgDir, err := frp.GetFrpcTomlDir()
	if err != nil {
		glog.Error(err)
		return nil, err
	}
	files, err := os.ReadDir(cfgDir)
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	glog.Debugf("files %+v", files)
	type Option struct {
		Label   string `json:"label"`
		Value   string `json:"value"`
		Content string `json:"content"`
	}

	var list []Option
	for _, f := range files {
		ext := strings.ToLower(filepath.Ext(f.Name()))
		if !f.IsDir() && ext == ".toml" {
			buffer, e := utils.Read(filepath.Join(cfgDir, f.Name()))
			if e == nil {
				list = append(list, Option{
					Label:   f.Name(),
					Value:   f.Name(),
					Content: string(buffer),
				})
			}

		}
	}
	return list, nil
}

func (this *frpc) clientRefresh(obj *iface.Message[any]) {
	if obj == nil {
		glog.Error("obj is nil")
	}
	data, err := this.getClientConfigs()
	if err != nil {
		glog.Error(err)
	} else {
		obj.Data = data
	}
	obj.Action = ws.CLIENT_REFRESH
	_ = this.sendMessageToWebSocketServer(obj)
}

func (this *frpc) sendMessageToWebSocketServer(obj *iface.Message[any]) error {
	if obj == nil {
		return fmt.Errorf("obj is nil")
	}
	//glog.Debugf("send %+v", *obj)
	err := ws.GetClientInstance().SendJSON(obj)
	if err != nil {
		glog.Error(err)
	} else {
		glog.Debugf("send sucess")
	}
	return err
}
