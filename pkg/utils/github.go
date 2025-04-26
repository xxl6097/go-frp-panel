package utils

import (
	"fmt"
	"github.com/xxl6097/glog/glog"
	"io"
	"net/http"
)

var githubProxys = []string{"https://ghfast.top/", "https://gh-proxy.com/", "https://ghproxy.1888866.xyz/"}
var indexCount = 0

var GithuApi = "https://api.github.com/repos/xxl6097/go-frp-panel/releases/latest"

func reqestGithubApi(baseUrl string) ([]byte, error) {
	glog.Debug("reqestGithubApi", baseUrl)
	resp, err := http.Get(baseUrl)
	if err != nil {
		glog.Errorf("请求失败:%v\n", err)
		indexCount = (indexCount + 1) % len(githubProxys) // 当 counter 达到 3 时，加 1 后取模结果为 0
		if indexCount == 0 {
			return nil, err
		}
		proxy := githubProxys[indexCount]
		return reqestGithubApi(fmt.Sprintf("%s%s", proxy, GithuApi))
	}
	defer resp.Body.Close() // 必须关闭响应体 [1,5,8](@ref)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		glog.Error("github请求失败", err)
		return nil, err
	}
	return body, nil
}

func GithubApiReqest() ([]byte, error) {
	//indexCount = (indexCount + 1) % len(githubProxys) // 当 counter 达到 3 时，加 1 后取模结果为 0
	//if indexCount == 0 {
	//	return nil, nil
	//}
	//proxy := githubProxys[indexCount]
	//return reqestGithubApi(fmt.Sprintf("%s%s", proxy, GithuApi))
	return reqestGithubApi(GithuApi)
}
