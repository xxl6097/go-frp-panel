package main

import (
	"errors"
	"fmt"
	"github.com/avast/retry-go/v4"
	"github.com/xxl6097/glog/glog"
	"time"
)

//	func PingRaw(ip string) bool {
//		conn, _ := icmp.ListenPacket("udp4", "0.0.0.0")
//		defer conn.Close()
//
//		msg := icmp.Message{
//			Type: ipv4.ICMPTypeEcho,
//			Code: 0,
//			Body: &icmp.Echo{ID: os.Getpid() & 0xffff, Seq: 1},
//		}
//		wb, _ := msg.Marshal(nil)
//
//		if _, err := conn.WriteTo(wb, &net.UDPAddr{IP: net.ParseIP(ip)}); err != nil {
//			return false
//		}
//
//		// 设置超时并读取响应
//		conn.SetReadDeadline(time.Now().Add(3 * time.Second))
//		reply := make([]byte, 1500)
//		_, _, err := conn.ReadFrom(reply)
//		return err == nil
//	}
func main() {
	//rawURL := "C://test.txt"
	//
	//// 提取路径部分并获取文件名
	//fileName := filepath.Base(rawURL)
	//fmt.Println("文件名:", fileName) // 输出: document.pdf
	////utils.ScanIP()
	////fmt.Println(utils.PingRaw("192.168.0.10"))
	//a := 30
	//b := 9
	//for i := min(a, b); i < max(a, b); i++ {
	//	fmt.Println(i)
	//}

	e := retry.Do(func() error {
		for i := 0; i < 5; i++ {
			glog.Println("wahaha", i)
			time.Sleep(time.Second)
		}
		return errors.New("error")
	}, retry.DelayType(retry.FixedDelay), retry.Delay(time.Second*2), retry.Attempts(5))
	fmt.Println("-->", e)
}
