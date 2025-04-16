package main

import (
	"encoding/json"
	"fmt"
	"github.com/xxl6097/glog/glog"
	"github.com/xxl6097/go-frp-panel/pkg/model"
	"io"
	"net/http"
	"strings"
	"time"
)

func CheckFrpc() error {
	var baseUrl = "https://api.github.com/repos/xxl6097/go-frp-panel/releases/latest"
	resp, err := http.Get(baseUrl)
	if err != nil {
		glog.Errorf("请求失败:%v\n", err)
		return nil
	}
	defer resp.Body.Close() // 必须关闭响应体 [1,5,8](@ref)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		glog.Error("github请求失败", err)
		return nil
	}
	glog.Debug("github请求成功")
	var result model.GitHubModel
	err = json.Unmarshal(body, &result)
	if err == nil {
		frpcUrls := make([]string, 0)
		for _, asset := range result.Assets {
			if strings.Contains(asset.Name, "frpc") {
				frpcUrls = append(frpcUrls, asset.BrowserDownloadUrl)
			}

		}
		fmt.Println(frpcUrls)
	}
	return err
}

// env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
func main() {

	ticker := time.NewTicker(60 * time.Second)
	defer ticker.Stop() // 必须关闭防止资源泄漏

	for {
		select {
		case t := <-ticker.C:
			fmt.Printf("定时任务执行于: %v\n", t.Format("2006-01-02 15:04:05"))
			// 业务逻辑（如数据同步、日志清理）
			_ = CheckFrpc()
		}
	}
}
