package main

import (
	"fmt"
	model2 "github.com/xxl6097/go-frp-panel/internal/com/model"
	"github.com/xxl6097/go-frp-panel/pkg/frp"
)

func main() {
	cfg := model2.FrpcBuffer{
		User: model2.User{
			User: "1234567890",
		},
		ServerAddr: "192.168.0.2",
		ServerPort: 6500,
		AdminUser:  "admin",
		AdminPass:  "admin",
	}

	//data, err := json.Marshal(cfg)
	//if err != nil {
	//	return
	//}
	////data, _ := json.Marshal(cfg)
	//fmt.Println(string(data))
	//code := utils.Encrypt(data, nil)
	////fmt.Println(hex.EncodeToString(code))
	//fmt.Println(code)
	//code = utils.Decrypt(code, nil)
	//var ccg model2.FrpcBuffer
	//json.Unmarshal([]byte(code), &ccg)
	//fmt.Println(fmt.Sprintf("%+v", ccg))

	secret, err := frp.EncodeSecret(&cfg)
	if err != nil {
		return
	}
	fmt.Println(secret)
	obj := frp.DecodeSecret(secret)
	fmt.Println(fmt.Sprintf("%+v", obj))
}
