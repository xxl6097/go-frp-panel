package main

import (
	"github.com/xxl6097/glog/glog"
	"github.com/xxl6097/go-frp-panel/pkg/utils"
)

func main() {
	//buffer := model.FrpcBuffer{}
	//key := make([]byte, 16)
	//data, _ := json.Marshal(buffer)
	//fmt.Println(string(data))
	//code := utils.Encrypt(data, key)
	////fmt.Println(hex.EncodeToString(code))
	//fmt.Println(code)
	//code = utils.Decrypt(code, key)
	//fmt.Println(string(code))

	net, e := utils.GetDeviceInfo()
	if e != nil {
		return
	}
	glog.Infof("Get device info: %+v", net)
}
