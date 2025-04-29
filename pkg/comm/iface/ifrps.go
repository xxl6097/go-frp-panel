package iface

import (
	httppkg "github.com/fatedier/frp/pkg/util/http"
)

type IFrps interface {
	Run()
	GetServer() *httppkg.Server
	CfgFile(string)
	OnFrpcConfigExport(string) (error, string)
	OnFrpcConfigImport(string) error
	Close()
}
