package frps

import (
	"encoding/json"
	"fmt"
	"github.com/xxl6097/glog/glog"
	"github.com/xxl6097/go-frp-panel/pkg/model"
	"github.com/xxl6097/go-frp-panel/pkg/utils"
	utils2 "github.com/xxl6097/go-service/gservice/utils"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func (this *frps) hasNewVersion(newVersion, clientsDir string) bool {
	if utils.CheckDir(clientsDir) {
		entries, err := os.ReadDir(clientsDir)
		if err == nil && len(entries) > 0 {
			entry := entries[0]
			if !entry.IsDir() {
				v1 := utils.GetVersionByFileName(entry.Name())
				isVersion := utils.CompareVersions(newVersion, v1)
				if isVersion <= 0 {
					return false
				}
			}
		}
	}
	_ = utils.EnsureDir(clientsDir)
	return true
}

func (this *frps) downloadFrpc(urls []string, dstDir string) {
	srcFilePath := utils2.DownloadFileWithCancelByUrls(urls)
	//glog.Println("下载完成", srcFilePath, dstDir)
	err := utils.MoveFileToDir(srcFilePath, dstDir)
	if err != nil {
		glog.Error("移动失败", err)
	} else {
		glog.Println("移动成功", srcFilePath, dstDir)
	}
}

func (this *frps) checkFrpc() {
	var baseUrl = "https://api.github.com/repos/xxl6097/go-frp-panel/releases/latest"
	resp, err := http.Get(baseUrl)
	if err != nil {
		glog.Errorf("请求失败:%v\n", err)
		return
	}
	defer resp.Body.Close() // 必须关闭响应体 [1,5,8](@ref)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		glog.Error("github请求失败", err)
		return
	}
	glog.Debug("checkFrpc请求成功")
	var result model.GitHubModel
	err = json.Unmarshal(body, &result)
	clientsDir := filepath.Join(this.binDir, "clients")
	if err == nil {
		if this.hasNewVersion(result.TagName, clientsDir) {
			githubProxys := utils.ParseMarkdownCodeToStringArray(result.Body)
			for _, asset := range result.Assets {
				if strings.Contains(asset.Name, "frpc") {
					newProxy := []string{}
					for _, proxy := range githubProxys {
						newUrl := fmt.Sprintf("%s%s", proxy, asset.BrowserDownloadUrl)
						newProxy = append(newProxy, newUrl)
					}
					glog.Debug("开始下载frpc", asset.BrowserDownloadUrl)
					go this.downloadFrpc(newProxy, clientsDir)
				} else {
					glog.Info("没有找到匹配的frpc客户端链接...")
				}
			}
		} else {
			glog.Info("客户端无需升级...")
		}
	} else {
		glog.Error(err)
	}
}

func (this *frps) check() {
	if !utils.HasDiskSpace() {
		glog.Error("本地磁盘空间不足...")
		return
	}
	glog.Error("开始检测客户端...")
	for {
		this.checkFrpc()
		time.Sleep(time.Hour)
	}
}

func (this *frps) CheckClients() {
	go this.check()
}
