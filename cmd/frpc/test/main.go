package main

import (
	"fmt"
	v1 "github.com/fatedier/frp/pkg/config/v1"
	"github.com/xxl6097/glog/glog"
	"github.com/xxl6097/go-frp-panel/internal/frpc"
	"github.com/xxl6097/go-frp-panel/pkg/comm"
	"github.com/xxl6097/go-frp-panel/pkg/utils"
	utils2 "github.com/xxl6097/go-service/gservice/utils"
	"os"
	"path/filepath"
)

func main() {
	cfg := &v1.ClientCommonConfig{
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
			Port:     6400,
			User:     "admin",
			Password: "admin",
		},
	}

	cfgBuffer := &comm.BufferConfig{
		Addr:  cfg.ServerAddr,
		Port:  cfg.ServerPort,
		User:  cfg.User,
		Token: cfg.User,
		Ports: []any{8089, "8200-9000"},
	}

	frpc.SetCfgModel(&frpc.CfgModel{Frpc: *cfg, Cfg: cfgBuffer})

	binPath, err := os.Executable()
	if err != nil {
		glog.Fatal("os.Executable() error", err)
	}
	cfgPath := filepath.Join(filepath.Dir(binPath), "config.toml")

	if !utils2.FileExists(cfgPath) {
		if err := os.WriteFile(cfgPath, utils.ObjectToTomlText(cfg), 0o600); err != nil {
			glog.Warnf("write content to frpc config file error: %v", err)
		}
	}

	fmt.Println(cfgPath)
	fmt.Println(string(utils.ObjectToTomlText(cfg)))
	cls, err := frpc.NewFrpc(nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("http://localhost:%d\n", cfg.WebServer.Port)
	cls.Run()
}
