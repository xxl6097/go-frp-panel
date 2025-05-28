package frpc

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/xxl6097/glog/glog"
	"github.com/xxl6097/go-frp-panel/pkg"
	"github.com/xxl6097/go-frp-panel/pkg/comm/iface"
	"github.com/xxl6097/go-frp-panel/pkg/comm/ws"
	"github.com/xxl6097/go-frp-panel/pkg/frp"
	"github.com/xxl6097/go-frp-panel/pkg/utils"
	"github.com/xxl6097/go-service/pkg/github"
	"net/http"
	"os"
	"os/exec"
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
		glog.Debugf("recv msg %+v", msg)
		switch msg.Action {
		case ws.CLIENT_REBOOT:
			if this.install == nil {
				glog.Error("install is nil")
				return
			}
			err = this.install.Restart()
			if err != nil {
				glog.Error("重启失败", err)
			}
			break
		case ws.CLIENT_UNINSTALL:
			if this.install == nil {
				glog.Error("install is nil")
				return
			}
			//err = this.install.RunCmd("uninstall")
			err = this.install.UnInstall()
			if err != nil {
				glog.Error("卸载失败", err)
			}
			break
		case ws.CLIENT_VERSION_CHECK:
			result, e := github.Api().DefaultRequest().CheckUpgrade(pkg.BinName, nil).Result()
			//args, e := utils.CheckVersionFromGithub()
			if e != nil {
				glog.Error(e)
				msg.Data = e.Error()
			} else {
				msg.Data = result
			}
			//if args == nil {
			//	msg.Data = "已经是最新版本～"
			//}
			_ = this.sendMessageToWebSocketServer(&msg)
			break
		case ws.CLIENT_NETWORLD:
			arr, e := utils.GetNetworkInterfaces()
			if e != nil {
				msg.Data = e.Error()
			} else {
				msg.Data = arr
			}
			_ = this.sendMessageToWebSocketServer(&msg)
			break
		case ws.CMD:
			sourceData, ok := msg.Data.(map[string]interface{})
			if ok {
				glog.Infof("sourceData %+v", sourceData)
				d := sourceData["data"]
				if d == nil {
					glog.Errorf("data is nil %+v", msg.Data)
					break
				}
				v, okk := d.(string)
				if !okk {
					glog.Infof("string err %+v", d)
					break
				}
				arrData := strings.Split(v, " ")
				var cmd *exec.Cmd
				if len(arrData) >= 2 {
					cmd = exec.Command(arrData[0], arrData[1:]...)
				} else {
					cmd = exec.Command(arrData[0])
				}
				output, err := cmd.CombinedOutput()
				if err != nil {
					msg.Data = err.Error()
				} else {
					msg.Data = string(output)
				}
			} else {
				msg.Data = fmt.Errorf("cmd err %+v", msg.Data)
			}
			_ = this.sendMessageToWebSocketServer(&msg)
			break
		case ws.CLIENT_VERSION_UPGRADE:
			body, ok := msg.Data.(map[string]interface{})
			if ok {
				glog.Infof("upgrade %+v", body)
				d := body["data"]
				if d == nil {
					glog.Errorf("data is nil %+v", msg.Data)
					break
				}
				url, okk := d.(string)
				if !okk {
					glog.Infof("类型不正确，upgrade %+v", d)
					break
				}
				msg.Data = fmt.Sprintf("开始下载 %v", url)
				_ = this.sendMessageToWebSocketServer(&msg)
				baseUrl := this.getUpgradeUrl(url)
				if baseUrl == "" {
					msg.Data = "下载链接空～"
					_ = this.sendMessageToWebSocketServer(&msg)
					break
				}
				msg.Data = fmt.Sprintf("下载成功，准备升级 %v", baseUrl)
				_ = this.sendMessageToWebSocketServer(&msg)
				err = this.Upgrade(context.Background(), baseUrl)
				//err = this.update(url)
				msg.Data = fmt.Sprintf("升级成功～")
				if err != nil {
					glog.Error(err)
					msg.Data = fmt.Sprintf("升级失败 %v", err.Error())
				}
				_ = this.sendMessageToWebSocketServer(&msg)
			} else {
				glog.Errorf("upgrade err %+v", msg.Data)
			}
			break
		default:
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
		glog.Debugf("send sucess %+v %+v %+v", obj.Action, obj.SseID, obj.DevIp)
	}
	return err
}
