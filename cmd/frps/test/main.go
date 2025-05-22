package main

import (
	"fmt"
	v1 "github.com/fatedier/frp/pkg/config/v1"
	"github.com/xxl6097/go-frp-panel/cmd"
	frps2 "github.com/xxl6097/go-frp-panel/pkg/frp/frps"
	"os"
	"path/filepath"
)

func main() {
	cmd.Execute(func() error {
		temp := os.TempDir()
		temp = filepath.Join(temp, "frps", "logs")
		err := os.MkdirAll(temp, 0755)
		if err != nil {
			fmt.Println(err)
		}
		cfg := &v1.ServerConfig{
			BindPort: 6000,
			BindAddr: "0.0.0.0",
			WebServer: v1.WebServerConfig{
				Addr:     "0.0.0.0",
				Port:     7200,
				User:     "admin",
				Password: "admin",
			}, HTTPPlugins: []v1.HTTPPluginOptions{
				{
					Name: "frps-panel",
					Addr: fmt.Sprintf("%s:%d", "0.0.0.0", 7200),
					Path: "/handler",
					Ops:  []string{"Login", "NewWorkConn", "NewUserConn", "NewProxy", "Ping"},
				},
			},
			Log: v1.LogConfig{
				To:      filepath.Join(temp, "frps.log"),
				MaxDays: 15,
			},
		}
		//frps.Test(&frps.CfgModel{
		//	Frps: *cfg,
		//})
		//content, _ := json.Marshal(cfg)
		svv, err := frps2.NewFrps(cfg, nil)
		if err != nil {
			return err
		}
		svv.Run()
		return nil
	})
}
