package main

import (
	"fmt"
	v1 "github.com/fatedier/frp/pkg/config/v1"
	"github.com/xxl6097/go-frp-panel/internal/frps"
	"github.com/xxl6097/go-service/pkg/ukey"
)

const B = '\x18'

func main() {
	cfg := &frps.CfgModel{
		Frps: v1.ServerConfig{
			BindPort: 7777,
			Auth: v1.AuthServerConfig{
				Token: "xiaxiaoli1",
			},
			WebServer: v1.WebServerConfig{
				User:     "admin",
				Password: "admin",
				Port:     7200,
				Addr:     "0.0.0.0",
			},
		},
	}
	cfg.Frps.Complete()
	newBytes, err := ukey.GenConfig(cfg.Bytes(), false)
	if err != nil {
		panic(err)
	}
	//http://uuxia.cn:8087/soft/acfrps/0.2.0/acfrps_0.2.0_windows_amd64.exe
	scrFilePath := "./dist/acfrps_0.2.7_windows_amd64.exe"
	dstFilePath := "./dist/acfrps-1.exe"
	//err = utils1.GenerateBin(scrFilePath, dstFilePath, ukey.GetBuffer(), newBytes)
	//if err != nil {
	//	panic(err)
	//}
	fmt.Println(scrFilePath, dstFilePath, ukey.GetBuffer(), newBytes)

	//byteArray, err := ukey.GenConfig(cfg, true)
	//if err != nil {
	//	panic(err)
	//}
	//size := len(byteArray)
	//fmt.Println("配置长度", size)
	//temp := bytes.Repeat([]byte{0x18}, size)
	//buffer := utils.PrintByteArrayAsConstant(temp)
	//fmt.Printf("var buffer = %s\n", buffer)
	//aa := bytes.Trim(temp, string(byte(0x18)))
	//fmt.Printf("aa = %v %d\n", aa, len(aa))

	//raw := "你是我的宝贝，我的阿胖宝～"
	//rec := "我是你的宝贝，你的大福宝～"
	//rawBytes := []byte(raw)
	//recBytes := []byte(rec)
	//ukey.GenSign(rawBytes, recBytes)

	////
	//rawKey, err := ukey.GetRawKey()
	//if err != nil {
	//	log.Fatal("构建原始key错误：", err)
	//}
	//fmt.Println("原始key", rawKey)
	//fmt.Println("buffer", ukey.GetBuffer())
	//fmt.Println("buffer", string(ukey.GetBuffer()))

	//err := util.Find("./dist/ac", buffer1)
	//fmt.Println("======>", err)
	//find("./dist/ac", []byte(buffer))
	//fmt.Println(len([]byte(buffer)), buffer)

	//rawKey, _ := ukey.GetRawKey()
	//fmt.Println(len(ukey.GetBuffer()), len(rawKey), len(ukey.GetKey()))

	//fmt.Println("buffer size", len(buffer))
	//fmt.Println("key size", len(key))
}
