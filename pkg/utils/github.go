package utils

import (
	"github.com/xxl6097/glog/glog"
	"github.com/xxl6097/go-frp-panel/pkg/model"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func init() {
	LoadGithubKey()
}

func LoadGithubKey() {
	fpath := filepath.Join(glog.AppHome("obj"), "githubKey.dat")
	obj, err := LoadWithGob[model.GithubKey](fpath)
	if err == nil && obj.ClientId != "" && obj.ClientSecret != "" {
		os.Setenv("GITHUB_CLIENT_ID", obj.ClientId)
		os.Setenv("GITHUB_CLIENT_SECRET", obj.ClientSecret)
	}
}

var githubProxys = []string{"https://ghfast.top/", "https://gh-proxy.com/", "https://ghproxy.1888866.xyz/"}
var indexCount = 0

var GithuApi = "https://api.github.com/repos/xxl6097/go-frp-panel/releases/latest"

func reqestGithubApi(baseUrl string) ([]byte, error) {
	glog.Debug("reqestGithubApi", baseUrl)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", baseUrl, nil)
	clientId := os.Getenv("GITHUB_CLIENT_ID")
	clientSecret := os.Getenv("GITHUB_CLIENT_SECRET")
	if clientId != "" || clientSecret != "" {
		req.SetBasicAuth(clientId, clientSecret) // 自动 Base64 编码
	}
	resp, err := client.Do(req)
	//resp, err := http.Get(baseUrl)
	if err != nil {
		glog.Errorf("请求失败:%v\n", err)
		//indexCount = (indexCount + 1) % len(githubProxys) // 当 counter 达到 3 时，加 1 后取模结果为 0
		//if indexCount == 0 {
		//	return nil, err
		//}
		//proxy := githubProxys[indexCount]
		//return reqestGithubApi(fmt.Sprintf("%s%s", proxy, GithuApi))
		return nil, err
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
