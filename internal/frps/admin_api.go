package frps

import (
	"fmt"
	v1 "github.com/fatedier/frp/pkg/config/v1"
	httppkg "github.com/fatedier/frp/pkg/util/http"
	"github.com/fatedier/frp/pkg/util/log"
	"github.com/gorilla/mux"
	"github.com/xxl6097/glog/glog"
	"github.com/xxl6097/go-frp-panel/internal/comm"
	"github.com/xxl6097/go-frp-panel/pkg"
	"github.com/xxl6097/go-frp-panel/pkg/utils"
	"github.com/xxl6097/go-service/gservice/ukey"
	"io"
	"net/http"
	"os"
	"time"
)

func (this *frps) adminHandlers(helper *httppkg.RouterRegisterHelper) {
	subRouter := helper.Router.NewRoute().Name("admin").Subrouter()
	subRouter.Use(helper.AuthMiddleware)
	staticPrefix := "/log/"
	//baseDir, _ := os.Getwd()
	baseDir := os.TempDir()
	subRouter.PathPrefix(staticPrefix).Handler(http.StripPrefix(staticPrefix, http.FileServer(http.Dir(baseDir))))

	// apis
	subRouter.HandleFunc("/api/panelinfo", this.apiPanelinfo).Methods("GET")
	subRouter.HandleFunc("/api/restart", this.upgrade.ApiRestart).Methods("GET")
	subRouter.HandleFunc("/api/shutdown", this.apiShutdown).Methods("GET")
	subRouter.HandleFunc("/api/version", this.upgrade.ApiVersion).Methods("GET")
	subRouter.HandleFunc("/api/upgrade", this.upgrade.ApiUpdate).Methods("POST")
	subRouter.HandleFunc("/api/upgrade", this.upgrade.ApiUpdate).Methods("PUT")
	subRouter.HandleFunc("/api/server/config/get", this.apiServerConfigGet).Methods("GET")
	subRouter.HandleFunc("/api/server/config/set", this.apiServerConfigSet).Methods("PUT")
	subRouter.HandleFunc("/api/proxy/{type}", this.apiProxyByType).Methods("GET")
}

// /api/shutdown
func (this *frps) apiShutdown(w http.ResponseWriter, r *http.Request) {
	res := comm.GeneralResponse{Code: 0}
	defer func() {
		log.Infof("Http response [%s]: res: %+v", r.URL.Path, res)
		w.WriteHeader(res.Code)
		if len(res.Msg) > 0 {
			_, _ = w.Write([]byte(res.Msg))
		}
	}()

	log.Infof("Http request: [%s]", r.URL.Path)
	res.Msg = "ok"
	//err := this.s.Stop()
	//err := utils.Shutdown()
	//if err != nil {
	//	res.Msg = err.Error()
	//}
}

func (this *frps) apiServerConfigSet(w http.ResponseWriter, r *http.Request) {
	res, f := comm.Response(r)
	defer f(w)
	// 读取请求体
	tomlBytes, err := io.ReadAll(r.Body)
	if err != nil {
		res.Error(fmt.Sprintf("读取body失败%v", err))
		return
	}
	glog.Println(tomlBytes)
	frpsCfg := v1.ServerConfig{}
	err = utils.TomlTextToObject(tomlBytes, &frpsCfg)
	if err != nil {
		res.Error(fmt.Sprintf("配置失败：%v", err))
		return
	}
	cfg := GetCfgModel()
	cfg.Frps = frpsCfg
	filePath, err := os.Executable()
	if err != nil {
		res.Error(fmt.Sprintf("%v", err))
		return
	}
	//下载和接收的最新文件 名称为上传文件的原始名称
	newBufferBytes, err := ukey.GenConfig(GetCfgModel(), false)
	if err != nil {
		res.Error(fmt.Sprintf("gen config err: %v", err))
		glog.Error(res.Msg)
		return
	}
	signFilePath, err := utils.SignAndInstall(newBufferBytes, ukey.GetBuffer(), filePath)
	if err != nil {
		res.Error(err.Error())
	} else {
		defer utils.Delete(signFilePath, "签名文件")
		err = this.install.Upgrade(signFilePath)
		if err != nil {
			res.Error(fmt.Sprintf("更新失败～%v", err))
			return
		}
		res.Ok("配置更新成功～")
	}
}

// /api/restart
func (this *frps) apiRestart(w http.ResponseWriter, r *http.Request) {
	res, f := comm.Response(r)
	defer f(w)
	res.Msg = "restart sucess"
	if res.Code == 0 && this.install != nil {
		go func() {
			time.Sleep(time.Second)
			err := this.install.Restart()
			if err != nil {
				glog.Error("重启失败")
			}
			glog.Error("重启ok")
		}()
	}
}

func (this *frps) apiPanelinfo(w http.ResponseWriter, r *http.Request) {
	res, f := comm.Response(r)
	defer f(w)
	res.Sucess("获取成功", map[string]interface{}{
		"appName":     pkg.AppName,
		"gitRevision": pkg.GitRevision,
		"gitBranch":   pkg.GitBranch,
		"goVersion":   pkg.GoVersion,
		"displayName": pkg.DisplayName,
		"description": pkg.Description,
		"appVersion":  pkg.AppVersion,
		"buildTime":   pkg.BuildTime,
	})
}

// /api/server/config/get
func (this *frps) apiServerConfigGet(w http.ResponseWriter, r *http.Request) {
	res, f := comm.Response(r)
	defer f(w)
	frpsToml := GetCfgModel().Frps
	glog.Println("获取Frps配置:", frpsToml)
	res.Raw = utils.ObjectToTomlText(frpsToml)
}

// /api/proxy/:type
func (svr *frps) apiProxyByType(w http.ResponseWriter, r *http.Request) {
	res := comm.GeneralResponse{Code: 200}
	params := mux.Vars(r)
	proxyType := params["type"]

	defer func() {
		log.Infof("Http response [%s]: code [%d]", r.URL.Path, res.Code)
		w.WriteHeader(res.Code)
		if len(res.Msg) > 0 {
			_, _ = w.Write([]byte(res.Msg))
		}
	}()
	log.Infof("Http request: [%s]", r.URL.Path)

	res.Msg = proxyType
}

//func (this *frps) signAndInstall(oldBufferBytes []byte, cfg interface{}, newFilePath string) error {
//	if !gore.FileExists(newFilePath) {
//		return fmt.Errorf("文件不存在：%s", newFilePath)
//	}
//	//下载和接收的最新文件 名称为上传文件的原始名称
//	defer utils.Delete(newFilePath, "升级文件")
//	newBufferBytes, err := ukey.GenConfig(cfg, false)
//	if err != nil {
//		return err
//	}
//	if newBufferBytes == nil || len(newBufferBytes) == 0 {
//		return fmt.Errorf("加密数据空～")
//	}
//	if oldBufferBytes == nil || len(oldBufferBytes) == 0 {
//		return fmt.Errorf("原始数据buffer空～")
//	}
//	//oldBufferBytes := ukey.UnInitializeBuffer()
//	//config.PrintCfg()
//
//	binFilePath, err := os.Executable()
//	if err != nil {
//		return fmt.Errorf("获取当前可执行文件路径出错: %v\n", err)
//	}
//
//	signFilePath := fmt.Sprintf("%s.sign", binFilePath)
//	glog.Printf("开始签名文件 %s\n", newFilePath)
//	err = utils.GenerateBin(newFilePath, signFilePath, oldBufferBytes, newBufferBytes)
//	if err != nil {
//		glog.Printf("签名失败 %v\n", err)
//		return err
//	}
//	//signFilePath 签名文件
//	defer utils.Delete(signFilePath, "签名文件")
//	err = this.install.CommApi(signFilePath)
//	if err != nil {
//		return fmt.Errorf("更新失败～%v", err)
//	}
//	return err
//}
