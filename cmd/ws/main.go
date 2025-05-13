package main

import (
	"fmt"
	"github.com/xxl6097/go-frp-panel/pkg/comm/ws"
	"github.com/xxl6097/go-frp-panel/pkg/utils"
)

func main() {
	//ws.GetClientInstance().Init("ws://uuxia.cn:6500/frp", "admin", "het002402")
	ws.GetClientInstance().Init("ws://192.168.0.3:6500/frp", "admin", "het002402")
	// 获取主IP地址
	primary, err := utils.GetDeviceInfo()
	if err != nil {
		fmt.Printf("\n获取主IP错误: %v\n", err)
	} else {
		fmt.Printf("\n主IP地址: %+v\n", primary)
	}
	select {}
}
