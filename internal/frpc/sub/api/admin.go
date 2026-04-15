package api

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/fatedier/frp/client"
	v1 "github.com/fatedier/frp/pkg/config/v1"
	httppkg "github.com/fatedier/frp/pkg/util/http"
	"github.com/gorilla/websocket"
	"github.com/xxl6097/glog/pkg/z"
	"github.com/xxl6097/go-frp-panel/internal/com/model"
	"github.com/xxl6097/go-frp-panel/pkg/comm/ws"
	"github.com/xxl6097/go-frp-panel/pkg/frp"
	"github.com/xxl6097/go-frp-panel/pkg/utils"
)

func (this *OneInstance) PostRun(svr *client.Service,
	cfg *v1.ClientCommonConfig,
	proxyCfgs []v1.ProxyConfigurer,
	visitorCfgs []v1.VisitorConfigurer) error {
	if this.Cb != nil {
		return this.Cb(svr, cfg, proxyCfgs, visitorCfgs)
	}
	return svr.Run(context.Background())
}

func decodeConfig(node *model.Node) *model.Node {
	if node != nil && node.Cfg != nil && node.Cfg.Metadatas != nil {
		secret := node.Cfg.Metadatas["secret"]
		z.L().Debug(fmt.Sprint("secret: %s", secret))
		if secret != "" {
			node.AdminConfig = frp.DecodeSecret(secret)
			z.L().Debug(fmt.Sprintf("解析secret %+v", node.AdminConfig))
			if node.AdminConfig == nil {
				z.L().Debug("adminConfig nil 无法启动wensocket ")
				return nil
			}
			return node

		}
	} else {
		z.L().Error("cfg.Metadatas is nil")
	}
	return nil
}

func runWebSocket(node *model.Node, recv func([]byte), openHandler func(*websocket.Conn, *http.Response)) {
	id := node.AdminConfig.User.ID
	addr := fmt.Sprintf("%s:%d", node.AdminConfig.ServerAddr, node.AdminConfig.ServerAdminPort)
	authorization := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", node.AdminConfig.AdminUser, node.AdminConfig.AdminPass)))
	ws.GetClientInstance().NewClient(id, addr, authorization)
	ws.GetClientInstance().SetMessageHandler(recv)
	ws.GetClientInstance().SetOpenHandler(openHandler)
}

func registerHandlers(node *model.Node) error {
	webServer, err := utils.GetPointerInstance[httppkg.Server]("webServer", node.Svr)
	if err != nil {
		return err
	}
	if webServer == nil {
		return fmt.Errorf("can't find webServer")
	}
	//webServer.RouteRegister(this.adminHandlers)
	return nil
}
