package main

import (
	"fmt"
	"github.com/xxl6097/glog/glog"
	"github.com/xxl6097/go-frp-panel/pkg/utils"
)

// env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
func main() {

	body, err := utils.GithubApiReqest()
	glog.Debug(body, err)
	name := "acfrps_v0.1.91_windows_amd64.exe"
	result := utils.SplitLastTwoByUnderscore(name)
	fmt.Printf("%-30s => %v\n", name, result)
	fmt.Printf("%s\n", utils.CleanExt(result[1]))

	proxys := []string{"a", "b", "c", "d", "e", "f"}
	var urls []string
	url := "http://uuxia.cn/file/goodname.ios"
	for _, proxy := range proxys {
		newUrl := fmt.Sprintf("%s%s", proxy, url)
		urls = append(urls, newUrl)
	}
	fmt.Printf("%-30s => %v\n", name, urls)
}
