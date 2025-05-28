package main

import (
	"fmt"
	"github.com/xxl6097/go-frp-panel/pkg/utils"
	utils2 "github.com/xxl6097/go-service/pkg/utils"
)

func main() {
	face, _ := utils.GetDeviceInfo()
	fmt.Printf("%+v\n", face)
	fmt.Printf("===>%+v\n", face.Ipv4)
	utils2.ExitAnyKey()
}
