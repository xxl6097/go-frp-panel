package frpc

import (
	"fmt"

	httppkg "github.com/fatedier/frp/pkg/util/http"
	"github.com/fatedier/frp/pkg/util/system"
	"github.com/xxl6097/glog/pkg/z"
	"github.com/xxl6097/go-frp-panel/internal/com/model"
	"github.com/xxl6097/go-frp-panel/internal/frpc/sub"
	"github.com/xxl6097/go-frp-panel/pkg/comm"
	"github.com/xxl6097/go-frp-panel/pkg/comm/iface"
	"github.com/xxl6097/go-frp-panel/pkg/frp"
	"github.com/xxl6097/go-frp-panel/pkg/utils"
	"github.com/xxl6097/go-service/pkg/gs/igs"
	"go.uber.org/zap"
)

type Admin struct {
	install igs.Service
	upgrade iface.IComm
	node    *model.Node
	nodes   map[string]*model.Node
}

func NewConfig(i igs.Service) iface.IFrpc {
	this := &Admin{
		install: i,
		nodes:   make(map[string]*model.Node),
		upgrade: comm.NewCommApi(i),
	}
	return this
}

func (this *Admin) Run() error {
	cfgDir, err := frp.GetFrpcTomlDir()
	if err != nil {
		return err
	}
	cfgFilePath, err := frp.GetFrpcMainTomlFilePath()
	if err != nil {
		return err
	}
	z.L().Debug("加载配置文件", zap.String("cfgFilePath", cfgFilePath))
	system.EnableCompatibilityMode()
	err = this.registerHandlers(nil)
	err = sub.Run(cfgFilePath, cfgDir)
	return err
}
func (this *Admin) registerHandlers(node *model.Node) error {
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
