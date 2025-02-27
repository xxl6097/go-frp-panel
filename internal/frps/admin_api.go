package frps

import (
	"fmt"
	v1 "github.com/fatedier/frp/pkg/config/v1"
	httppkg "github.com/fatedier/frp/pkg/util/http"
	"github.com/fatedier/frp/pkg/util/log"
	"github.com/gorilla/mux"
	"github.com/xxl6097/glog/glog"
	"github.com/xxl6097/go-frp-panel/internal/comm"
	"github.com/xxl6097/go-frp-panel/internal/comm/ukey"
	"github.com/xxl6097/go-frp-panel/pkg"
	"github.com/xxl6097/go-frp-panel/pkg/utils"
	"io"
	"net/http"
	"os"
	"path/filepath"
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
	//glog.Println(tomlBytes)
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
	//newBytes, err := ukey.GenConfig(cfg, false)
	//if err != nil {
	//	res.Msg = fmt.Sprintf("加密失败%v", err)
	//	res.Code = -1
	//	return
	//}
	//err = utils.LocalGenerateBin(ukey.GetBuffer(), newBytes)
	//if err != nil {
	//	res.Msg = err.Error()
	//	res.Code = -1
	//	glog.Error(res.Msg)
	//	return
	//}
	//res.Msg = "配置成功"
	//glog.Println("开始重启服务")
	//if res.Code == 0 && this.install != nil {
	//	go func() {
	//		err = this.install.Restart()
	//		if err != nil {
	//			glog.Error("开始重启失败")
	//		}
	//		glog.Error("开始重启ok")
	//	}()
	//}
}

func (this *frps) apiUpgradePOST(w http.ResponseWriter, r *http.Request) {
	res, f := comm.Response(r)
	defer f(w)
	//ParseMultipartForm将请求的主体作为multipart/form-data解析。请求的整个主体都会被解析，得到的文件记录最多 maxMemery字节保存在内存，其余部分保存在硬盘的temp文件里。如果必要，ParseMultipartForm会自行调用 ParseForm。重复调用本方法是无意义的
	//设置内存大小
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		//httpapi.Error(w, err.Error(), httpapi.StatusInternalServerError)
		res.Msg = err.Error()
		glog.Error(res.Msg)
		res.Code = -1
		return
	}

	// 获取上传的文件
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer file.Close()

	currentPath, err := os.Executable()
	if err != nil {
		res.Msg = fmt.Sprintf("获取当前可执行文件路径出错: %v\n", err)
		glog.Error(res.Msg)
		res.Code = -1
		return
	}

	dstFilePath := filepath.Join(filepath.Dir(currentPath), handler.Filename)
	// 创建本地文件以保存上传的文件
	dst, err := os.Create(dstFilePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()
	pw := &utils.ProgressWriter{TotalSize: handler.Size, Progress: -1, Title: "文件下载："}
	_, err = io.Copy(io.MultiWriter(dst, pw), file)
	// 将上传的文件内容复制到本地文件
	//_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newBufferBytes, err := ukey.GenConfig(GetCfgModel(), false)
	if err != nil {
		glog.Println("加密失败～", err)
		res.Msg = fmt.Sprintf("加密失败～%v", err)
		res.Code = -1
		return
	}
	if newBufferBytes == nil || len(newBufferBytes) == 0 {
		res.Msg = fmt.Sprintf("加密信息空～")
		res.Code = -1
		return
	}
	oldBufferBytes := ukey.UnInitializeBuffer()
	PrintCfg()
	glog.Printf("\n请求地址：%s\n旧签名信息大小：%d\n新签名信息大小：%d\n", r.URL.Path, len(oldBufferBytes), len(newBufferBytes))
	glog.Println("install", this.install)

	newBinPath, err := utils.UpdateByUpload(dstFilePath, oldBufferBytes, newBufferBytes)
	if err != nil {
		res.Msg = fmt.Sprintf("加密失败～%v", err)
		res.Code = -1
		glog.Error(res.Msg)
		return
	}
	if newBinPath == "" {
		res.Msg = fmt.Sprintf("文件不存～%v", err)
		res.Code = -1
		glog.Error(res.Msg)
		return
	}
	err = this.install.Upgrade(newBinPath)
	if err != nil {
		res.Msg = fmt.Sprintf("更新失败～%v", err)
		res.Code = -1
		glog.Error(res.Msg)
		return
	}
	// 返回成功响应
	res.Msg = fmt.Sprintf("文件 %s 上传成功", handler.Filename)
	glog.Info(res.Msg)
	res.Msg = "程序更新成功"
	//if res.Code == 0 && this.s != nil {
	//	go func() {
	//		err = this.s.Restart()
	//		if err != nil {
	//			glog.Error("开始重启失败")
	//		}
	//		glog.Error("开始重启ok")
	//	}()
	//}
	//time.Sleep(time.Second)
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
