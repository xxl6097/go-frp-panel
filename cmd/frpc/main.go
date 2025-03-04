package main

import (
	"github.com/xxl6097/glog/glog"
	"github.com/xxl6097/go-frp-panel/cmd/frpc/service"
	"github.com/xxl6097/go-service/gservice"
)

//go:generate goversioninfo -icon=resource/icon.ico -manifest=resource/goversioninfo.exe.manifest
func main() {
	err := gservice.Run(service.Service{})
	if err != nil {
		glog.Error("程序启动出错了", err)
	}
	//glog.Println("服务程序启动成功，主进程退出", os.Getegid())
}
