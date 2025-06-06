package frps

import (
	"encoding/json"
	plugin "github.com/fatedier/frp/pkg/plugin/server"
	httppkg "github.com/fatedier/frp/pkg/util/http"
	"github.com/xxl6097/glog/glog"
	"github.com/xxl6097/go-frp-panel/internal/com/model"
	"github.com/xxl6097/go-frp-panel/pkg/frp"
	"github.com/xxl6097/go-frp-panel/pkg/utils"
	"log"
	"net/http"
)

type HTTPError struct {
	Code int
	Err  error
}
type Response struct {
	Msg string `json:"msg"`
}

func (this *frps) handlers(helper *httppkg.RouterRegisterHelper) {
	subRouter := helper.Router.NewRoute().Name("admin").Subrouter()
	subRouter.HandleFunc("/handler", this.apiHandler).Methods("POST")
}

func decodeMetas(mapData map[string]string) *model.User {
	//v, ok := mapData["secret"]
	//if !ok {
	//	return nil
	//}
	//buffer := frp.DecodeSecret(v)
	//if buffer == nil {
	//	return nil
	//}
	//return &buffer.User
	return frp.DecodeMetas(mapData)
}

func (c *frps) apiHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	request, err := utils.BindJSON[plugin.Request](r)
	jsonStr, err := json.Marshal(request.Content)
	var response plugin.Response
	if request.Op == "Login" {
		content := plugin.LoginContent{}
		err = json.Unmarshal(jsonStr, &content)
		response = c.HandleLogin(&content)
	} else if request.Op == "NewProxy" {
		content := plugin.NewProxyContent{}
		err = json.Unmarshal(jsonStr, &content)
		response = c.HandleNewProxy(&content)
	} else if request.Op == "Ping" {
		content := plugin.PingContent{}
		err = json.Unmarshal(jsonStr, &content)
		response = c.HandlePing(&content)
	} else if request.Op == "NewWorkConn" {
		content := plugin.NewWorkConnContent{}
		err = json.Unmarshal(jsonStr, &content)
		response = c.HandleNewWorkConn(&content)
	} else if request.Op == "NewUserConn" {
		content := plugin.NewUserConnContent{}
		err = json.Unmarshal(jsonStr, &content)
		response = c.HandleNewUserConn(&content)
	}

	if err != nil {
		glog.Printf("handle %s error: %v\n", r.URL.Path, err)
		response.RejectReason = err.Error()
		response.Reject = true
	}
	bb, err := json.Marshal(response)
	if err != nil {
		glog.Printf("【%s】Failed %v\n", request.Op, err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		//glog.Printf("【%s】Sucess %s\n", request.Op, string(bb))
		w.Write(bb)
	}
}

func (c *frps) HandleLogin(content *plugin.LoginContent) plugin.Response {
	//token := content.Metas["token"]
	//id := content.Metas["id"]
	user := decodeMetas(content.Metas)
	if user == nil {
		var res plugin.Response
		res.Reject = true
		res.RejectReason = "ID和Token不能为空"
		return res
	}
	return c.JudgeToken(user.ID, user.Token)
}

func (c *frps) HandleNewProxy(content *plugin.NewProxyContent) plugin.Response {
	//token := content.User.Metas["token"]
	//id := content.User.Metas["id"]
	//judgeToken := c.JudgeToken(id, token)
	//if judgeToken.Reject {
	//	return judgeToken
	//}
	//return c.JudgePort(content)

	user := decodeMetas(content.User.Metas)
	if user == nil {
		var res plugin.Response
		res.Reject = true
		res.RejectReason = "ID和Token不能为空"
		return res
	}
	judgeToken := c.JudgeToken(user.ID, user.Token)
	if judgeToken.Reject {
		return judgeToken
	}
	return c.JudgePort(content)
}

func (c *frps) HandlePing(content *plugin.PingContent) plugin.Response {
	//token := content.User.Metas["token"]
	//id := content.User.Metas["id"]
	//return c.JudgeToken(id, token)
	user := decodeMetas(content.User.Metas)
	if user == nil {
		var res plugin.Response
		res.Reject = true
		res.RejectReason = "ID和Token不能为空"
		return res
	}
	return c.JudgeToken(user.ID, user.Token)
}

func (c *frps) HandleNewWorkConn(content *plugin.NewWorkConnContent) plugin.Response {
	//token := content.User.Metas["token"]
	//id := content.User.Metas["id"]
	//return c.JudgeToken(id, token)
	user := decodeMetas(content.User.Metas)
	if user == nil {
		var res plugin.Response
		res.Reject = true
		res.RejectReason = "ID和Token不能为空"
		return res
	}
	return c.JudgeToken(user.ID, user.Token)
}

func (c *frps) HandleNewUserConn(content *plugin.NewUserConnContent) plugin.Response {
	//token := content.User.Metas["token"]
	//id := content.User.Metas["id"]
	//return c.JudgeToken(id, token)
	user := decodeMetas(content.User.Metas)
	if user == nil {
		var res plugin.Response
		res.Reject = true
		res.RejectReason = "ID和Token不能为空"
		return res
	}
	return c.JudgeToken(user.ID, user.Token)
}
func (c *frps) JudgeToken(id string, token string) plugin.Response {
	var res plugin.Response
	if id == "" || token == "" {
		res.Reject = true
		res.RejectReason = "ID和Token不能为空"
	} else {
		ok, err := model.JudgeToken(id, token)
		if ok {
			res.Unchange = true
		} else {
			res.Reject = true
			if err != nil {
				res.RejectReason = err.Error()
			}
		}
	}

	return res
}

func (c *frps) JudgePort(content *plugin.NewProxyContent) plugin.Response {
	var res plugin.Response
	supportProxyTypes := []string{
		"tcp", "tcpmux", "udp", "http", "https",
	}
	proxyType := content.ProxyType
	if !utils.StringContains(proxyType, supportProxyTypes) {
		log.Printf("proxy type [%v] not support, plugin do nothing", proxyType)
		res.Unchange = true
		return res
	}

	user := decodeMetas(content.User.Metas)
	if user == nil {
		res.Reject = true
		res.RejectReason = "ID和Token不能为空"
		return res
	}
	id := user.ID
	//user := content.User.User
	//id := content.User.Metas["id"]
	userPort := content.RemotePort
	userDomains := content.CustomDomains
	userSubdomain := content.SubDomain

	ok, err := model.JudgePort(id, proxyType, userPort, userDomains, userSubdomain)
	if ok {
		res.Reject = true
		res.RejectReason = err.Error()
	} else {
		res.Unchange = true
	}
	return res
}
