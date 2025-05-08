package iface

import (
	httppkg "github.com/fatedier/frp/pkg/util/http"
	"github.com/xxl6097/go-frp-panel/pkg/model"
)

type IFrps interface {
	Run()
	GetServer() *httppkg.Server
	CfgFile(string)
	OnFrpcConfigExport(string) (error, string)
	OnFrpcConfigImport(string) error
	GetCloudApi() *model.CloudApi
	SetCloudApi(*model.CloudApi)
	Close()
}
