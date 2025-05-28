package main

import (
	"fmt"
	v1 "github.com/fatedier/frp/pkg/config/v1"
	"github.com/xxl6097/glog/glog"
	"github.com/xxl6097/go-frp-panel/internal/frpc"
	"github.com/xxl6097/go-frp-panel/pkg"
	"github.com/xxl6097/go-frp-panel/pkg/frp"
	frpc2 "github.com/xxl6097/go-frp-panel/pkg/frp/frpc"
	"github.com/xxl6097/go-frp-panel/pkg/utils"
	"github.com/xxl6097/go-service/pkg/github"
)

func init() {
	pkg.BinName = "acfrpc_v0.4.15_linux_amd64"
	pkg.GithubUser = "xxl6097"
	pkg.GithubRepo = "go-frp-panel"
	github.Api().SetName(pkg.GithubUser, pkg.GithubRepo)
}
func main() {
	ccc := &v1.ClientCommonConfig{
		ServerAddr: "192.168.0.3",
		ServerPort: 6000,
		User:       "clife-fnos",
		Metadatas: map[string]string{
			"token": "clife-fnos",
		},
		Log: v1.LogConfig{
			To: "console",
		},
		WebServer: v1.WebServerConfig{
			Addr:     "0.0.0.0",
			Port:     6401,
			User:     "admin",
			Password: "admin",
		},
	}

	tcpProxy := v1.TypedProxyConfig{
		ProxyConfigurer: &v1.TCPProxyConfig{
			RemotePort: 6503,
			ProxyBaseConfig: v1.ProxyBaseConfig{
				Name: "test-001",
				Type: "tcp",
				ProxyBackend: v1.ProxyBackend{
					LocalIP:   "0.0.0.0",
					LocalPort: 6401,
				},
			},
		},
	}

	glog.Infof("tcpProxy:%+v", tcpProxy)
	var proxies []v1.TypedProxyConfig
	proxies = append(proxies, tcpProxy)

	cfg := &v1.ClientConfig{
		ClientCommonConfig: *ccc,
		Proxies:            proxies,
	}
	glog.Infof("ClientConfig:%+v", cfg)
	frpc.SetCfgModel(&frpc.CfgModel{Frpc: *cfg})

	err := frp.WriteFrpcMainConfigWithOut(cfg)
	//err = frp.WriteFrpcMainConfig(cfg)
	if err != nil {
		glog.Warnf("write content to frpc config file error: %v", err)
	}

	_ = ReadFrcMainConfigWithOut()
	//fmt.Println(cfgPath)
	//fmt.Println(string(utils.ObjectToTomlText(cfg)))
	cls, err := frpc2.NewFrpc(nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("http://localhost:%d\n", cfg.WebServer.Port)
	cls.Run()
}

func ReadFrcMainConfigWithOut() error {
	content, err := frp.ReadFrpToml(frp.GetFrpcMainTomlFileName())
	if err != nil {
		return err
	}
	fmt.Println(string(content))
	cfg := v1.ClientConfig{}
	err = utils.TomlTextToObject(content, &cfg)
	if err != nil {
		return err
	}
	fmt.Println(cfg)
	return nil
}
