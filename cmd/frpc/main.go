package main

import (
	"github.com/xxl6097/go-frp-panel/cmd/frpc/service"
)

//go:generate goversioninfo -icon=resource/icon.ico -manifest=resource/goversioninfo.exe.manifest
func main() {
	service.Bootstrap()
}
