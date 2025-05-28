package frps

import (
	"github.com/xxl6097/glog/glog"
	"github.com/xxl6097/go-frp-panel/pkg"
	"github.com/xxl6097/go-service/pkg/github"
	"os"
	"time"
)

func (this *frps) checkFrpc() {
	glog.Debug("checkFrpc请求")
	github.Api().Request(pkg.GithubUser, pkg.GithubRepo)
}

func (this *frps) check() {
	glog.Error("开始检测客户端...")
	for {
		this.checkFrpc()
		time.Sleep(time.Hour * 8)
	}
}

func (this *frps) CheckVersion() {
	checks := os.Getenv("CHECK_CLIENTS")
	if checks != "" {
		return
	}
	go this.check()
}
