package main

import (
	"fmt"
	"github.com/xxl6097/go-frp-panel/pkg/utils"
	"regexp"
)

func test() {
	filename := "acfrps_v1.34.0_windows_amd64.exe"
	// 匹配 v 开头 + 数字组合（支持多级版本号）
	re := regexp.MustCompile(`_v\d+\.\d+\.\d+_`)
	newName := re.ReplaceAllString(filename, "_v0.0.0_") // 替换为单个下划线
	fmt.Println(newName)                                 // 输出 acfrps_windows_amd64.exe
}
func main() {
	test()
	binurl := "https://github.com/xxl6097/go-frp-panel/releases/download/v0.0.47/acfrps_v0.0.47_linux_amd64"
	fmt.Println(utils.IsURLValidAndAccessible(binurl))
}
