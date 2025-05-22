package main

import (
	"encoding/json"
	"fmt"
	"github.com/xxl6097/go-frp-panel/pkg/comm"
	"github.com/xxl6097/go-frp-panel/pkg/utils"
)

func main() {
	buffer := comm.BufferConfig{
		Addr:          "127.0.0.1",
		Port:          8080,
		ApiPort:       8081,
		Authorization: "xiaxiaoli1",
		User:          "uuxia",
	}
	key := make([]byte, 16)
	data, _ := json.Marshal(buffer)
	fmt.Println(string(data))
	code := utils.Encrypt(data, key)
	//fmt.Println(hex.EncodeToString(code))
	fmt.Println(code)
	code = utils.Decrypt(code, key)
	fmt.Println(string(code))
}
