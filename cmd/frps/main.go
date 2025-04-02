package main

import (
	"github.com/xxl6097/glog/glog"
	"github.com/xxl6097/go-frp-panel/cmd/frps/service"
	"github.com/xxl6097/go-frp-panel/pkg"
	"github.com/xxl6097/go-frp-panel/pkg/utils"
	"github.com/xxl6097/go-service/gservice"
)

func init() {
	if utils.IsMacOs() {
		pkg.AppName = "acfrps"
		pkg.DisplayName = "acfrps"
		pkg.Description = "acfrps"
	}
}

//go:generate goversioninfo -icon=resource/icon.ico -manifest=resource/goversioninfo.exe.manifest
func main() {
	err := gservice.Run(service.Service{})
	if err != nil {
		glog.Error("程序启动出错了", err)
	}
}
