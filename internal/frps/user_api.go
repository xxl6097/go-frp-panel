package frps

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	v1 "github.com/fatedier/frp/pkg/config/v1"
	"github.com/xxl6097/glog/glog"
	model2 "github.com/xxl6097/go-frp-panel/internal/com/model"
	comm2 "github.com/xxl6097/go-frp-panel/pkg/comm"
	"github.com/xxl6097/go-frp-panel/pkg/model"
	"github.com/xxl6097/go-frp-panel/pkg/utils"
	"github.com/xxl6097/go-service/gservice/ukey"
	utils2 "github.com/xxl6097/go-service/gservice/utils"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func (this *frps) apiUserCreate(w http.ResponseWriter, r *http.Request) {
	res, f := comm2.Response(r)
	defer f(w)
	//body, err := io.ReadAll(r.Body)
	//if err != nil {
	//	res.Response(400, fmt.Sprintf("read request body error: %v", err))
	//	glog.Warnf("%s", res.Msg)
	//	return
	//}
	//fmt.Println(string(body))
	u, err := utils.GetDataByJson[model2.User](r)
	if err != nil {
		res.Err(err)
		return
	}
	if u == nil {
		res.Error("token is nil")
		return
	}
	err = u.CreateUserByID()
	if err != nil {
		res.Err(err)
		return
	}
	res.Ok("密钥创建成功")
}

func (this *frps) apiUserDelete(w http.ResponseWriter, r *http.Request) {
	res, f := comm2.Response(r)
	defer f(w)
	users, err := utils.GetDataByJson[[]struct {
		User string `json:"user"`
		ID   string `json:"id"`
	}](r)
	if err != nil {
		res.Err(err)
		return
	}
	if users == nil {
		res.Error("tokens is nil")
		return
	}
	for _, u := range *users {
		err = model2.DeleteUser(u.ID)
	}
	//err = this.repo.Delete(u.User)
	//if err != nil {
	//	res.Err(err)
	//	return
	//}
	res.Ok("密钥删除成功")
}

func (this *frps) apiUserDeleteAll(w http.ResponseWriter, r *http.Request) {
	res, f := comm2.Response(r)
	defer f(w)
	userDir, err := utils.GetUserDir()
	if err != nil {
		res.Err(err)
		return
	}
	err = utils2.DeleteAll(userDir, "apiUserDeleteAll")
	if err != nil {
		res.Err(err)
		return
	}

	//err = this.repo.Delete(u.User)
	//if err != nil {
	//	res.Err(err)
	//	return
	//}
	res.Ok("删除成功")
}

func (this *frps) apiUserUpdate(w http.ResponseWriter, r *http.Request) {
	res, f := comm2.Response(r)
	defer f(w)
	u, err := utils.GetDataByJson[model2.User](r)
	if err != nil {
		res.Err(err)
		return
	}
	if u == nil {
		res.Error("token is nil")
		return
	}
	glog.Printf("%+v\n", u)
	err = u.UpdateUser()

	if err != nil {
		res.Err(err)
		return
	}
	res.Ok("密钥更新成功")
	a, _ := this.GetUserAll()
	fmt.Printf("结果：%+v\n", a)
}
func (this *frps) apiUserAll(w http.ResponseWriter, r *http.Request) {
	res, f := comm2.Response(r)
	defer f(w)
	datas, err := this.GetUserAll()
	if err != nil {
		res.Error("无数据")
		glog.Error(err)
		return
	}
	res.Sucess("全部数据获取成功", datas)
	//glog.Infof("%+v\n", datas)
}

func (this *frps) apiClientGet(w http.ResponseWriter, r *http.Request) {
	res, f := comm2.Response(r)
	defer f(w)
	binPath, err := os.Executable()
	if err != nil {
		res.Error(fmt.Sprintf("获取当前可执行文件路径出错: %v\n", err))
		glog.Error(res.Msg)
		return
	}
	configDir := filepath.Dir(binPath)
	configPath := filepath.Join(configDir, "clients")
	//glog.Infof("扫描路径:%s", configPath)
	nodes := utils.GetNodes(configPath)
	if nodes == nil || len(nodes) == 0 {
		nodes = utils.ToTree("", this.frpcGithubDownloadUrls)
	}
	res.Data = nodes
	//glog.Infof("扫描结果:%v", res.Data)
}

func (this *frps) apiClientListGet(w http.ResponseWriter, r *http.Request) {
	res, f := comm2.Response(r)
	defer f(w)
	if this.webSocketApi == nil {
		res.Error("webSocketApi is nil")
		return
	}
	timeObj, err := utils.GetDataByJson[struct {
		FrpID string `json:"frpId"`
	}](r)
	if err != nil {
		res.Err(err)
		return
	}
	if timeObj == nil {
		res.Error("timeObj is nil")
		return
	}
	sessions := this.webSocketApi.GetList(timeObj.FrpID)
	res.Data = sessions
}

func (this *frps) apiFrpsGet(w http.ResponseWriter, r *http.Request) {
	res, f := comm2.Response(r)
	defer f(w)
	res.Data = utils.ToTree("", this.frpsGithubDownloadUrls)
	//glog.Infof("frpsGithubDownloadUrls:%v", this.frpsGithubDownloadUrls)
	//glog.Infof("frps地址扫描:%v", res.Data)
}

//func (this *frps) parseUser(data map[string]interface{}) {
//	glog.Println(data)
//	u := User{
//		User:       data["user"].(string),
//		Token:      data["token"].(string),
//		SseId:         data["id"].(string),
//		Comment:    data["comment"].(string),
//		Ports:      ToPorts(data["ports"].([]any)),
//		Domains:    data["domains"].([]string),
//		Subdomains: data["subdomains"].([]string),
//		Enable:     data["enable"].(bool),
//	}
//	glog.Error(u)
//}

//func (this *frps) apiClientGenPut(w http.ResponseWriter, r *http.Request) {
//	res := &comm2.GeneralResponse{Code: 0}
//
//	var body struct {
//		Addr      string               `json:"addr"`
//		Port      int                  `json:"port"`
//		ApiPort   int                  `json:"apiPort"`
//		User      model2.User          `json:"user"`
//		Proxy     *v1.TypedProxyConfig `json:"proxy"`
//		WebServer *v1.WebServerConfig  `json:"webserver"`
//	}
//	jstr := r.FormValue("data")
//	err := json.Unmarshal([]byte(jstr), &body)
//	glog.Infof("data:%+v", body)
//
//	err = r.ParseMultipartForm(32 << 20)
//	if err != nil {
//		res.Error("body can't be empty")
//		glog.Error(res.Msg)
//		return
//	}
//	// 获取上传的文件
//	file, handler, err := r.FormFile("file")
//	if err != nil {
//		res.Error("body no file")
//		return
//	}
//	defer file.Close()
//
//	glog.Info(handler.Filename)
//
//	binPath := filepath.Join(glog.GetCrossPlatformDataDir("temp"), handler.Filename)
//	dst, err := os.Create(binPath)
//	if err != nil {
//		res.Error(fmt.Sprintf("create file %s error: %v", handler.Filename, err))
//		return
//	}
//	defer utils2.DeleteAll(binPath, "upload gen file")
//	buf := this.upgrade.GetBuffer().Get().([]byte)
//	defer this.upgrade.GetBuffer().Put(buf)
//	_, err = io.CopyBuffer(dst, file, buf)
//	dst.Close()
//	if err != nil {
//		res.Error(err.Error())
//		return
//	}
//	glog.Info("上传成功", binPath)
//
//	tpl, err := os.Open(binPath)
//	if err != nil {
//		msg := fmt.Errorf("打开文件失败：%v", err)
//		glog.Error(msg)
//		http.Error(w, msg.Error(), http.StatusGatewayTimeout)
//		return
//	}
//	defer tpl.Close()
//
//	fileName := filepath.Base(binPath)
//	w.Header().Add("Content-Transfer-Encoding", "binary")
//	w.Header().Add("Content-Type", "application/octet-stream")
//	if stat, err := tpl.Stat(); err == nil {
//		w.Header().Add(`Content-Length`, strconv.FormatInt(stat.Size(), 10))
//	}
//	w.Header().Add(`Content-Disposition`, fmt.Sprintf("attachment; filename=\"%s\"", fileName))
//	//cfgBuffer := ukey.GetBuffer()
//	if GetCfgModel() == nil {
//		msg := fmt.Errorf("GetCfgModel is nil")
//		glog.Error(msg)
//		http.Error(w, msg.Error(), http.StatusGatewayTimeout)
//		return
//	}
//	authorization := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", GetCfgModel().Frps.WebServer.User, GetCfgModel().Frps.WebServer.Password)))
//	bindPort := GetCfgModel().Frps.BindPort
//	if body.Port > 0 {
//		bindPort = body.Port
//	}
//	cfgBuffer := bytes.Repeat([]byte{byte(ukey.B)}, len(ukey.GetBuffer()))
//	cfg := comm2.BufferConfig{
//		Addr:          body.Addr,
//		Port:          bindPort,
//		Authorization: authorization,
//		ApiPort:       body.ApiPort,
//		ID:            body.User.ID,
//		User:          body.User.User,
//		Token:         body.User.Token,
//		Comment:       body.User.Comment,
//		Ports:         body.User.Ports,
//		Domains:       body.User.Domains,
//		Subdomains:    body.User.Subdomains,
//		Proxy:         body.Proxy,
//		WebServer:     body.WebServer,
//	}
//
//	glog.Infof("BufferConfig: %+v", cfg)
//	cfgNewBytes, err := ukey.GenConfig(cfg, false)
//	if err != nil {
//		msg := fmt.Errorf("文件签名失败：%v", err)
//		glog.Error(msg)
//		http.Error(w, msg.Error(), http.StatusHTTPVersionNotSupported)
//		return
//	}
//
//	//err = frpc.TestLoadBuffer(cfgNewBytes)
//	//glog.Infof("TestLoadBuffer: %+v\n", err)
//
//	dstFile := filepath.Join(glog.GetCrossPlatformDataDir("temp", utils2.SecureRandomID()), fileName)
//	outFile, err := os.Create(dstFile)
//	if err != nil {
//		_ = utils2.DeleteAll(dstFile, "创建失败，删除")
//		http.Error(w, fmt.Errorf("创建失败：%v", err).Error(), http.StatusHTTPVersionNotSupported)
//		return
//	}
//	defer outFile.Close()
//	defer utils2.DeleteAll(dstFile, "gen file")
//
//	prevBuffer := make([]byte, 0)
//	for {
//		thisBuffer := make([]byte, 1024)
//		n, err := tpl.Read(thisBuffer)
//		thisBuffer = thisBuffer[:n]
//		tempBuffer := append(prevBuffer, thisBuffer...)
//		bufIndex := bytes.Index(tempBuffer, cfgBuffer)
//		if bufIndex > -1 {
//			tempBuffer = bytes.Replace(tempBuffer, cfgBuffer, cfgNewBytes, -1)
//		}
//		//w.Write(tempBuffer[:len(prevBuffer)])
//		outFile.Write(tempBuffer[:len(prevBuffer)])
//		prevBuffer = tempBuffer[len(prevBuffer):]
//		if err != nil {
//			break
//		}
//	}
//	if len(prevBuffer) > 0 {
//		//w.Write(prevBuffer)
//		outFile.Write(prevBuffer)
//		prevBuffer = nil
//	}
//	http.ServeFile(w, r, dstFile)
//}

//
//func (this *frps) apiClientGen(w http.ResponseWriter, r *http.Request) {
//	ctx, cancel := context.WithCancel(context.Background())
//	defer cancel()
//	body, err := utils.GetDataByJson[struct {
//		BinPath   string               `json:"binPath"`
//		BinUrl    string               `json:"binUrl"`
//		Addr      string               `json:"addr"`
//		Port      int                  `json:"port"`
//		ApiPort   int                  `json:"apiPort"`
//		User      model2.User          `json:"user"`
//		Proxy     *v1.TypedProxyConfig `json:"proxy"`
//		WebServer *v1.WebServerConfig  `json:"webserver"`
//	}](r)
//	if err != nil {
//		glog.Error("解析Json对象失败", err)
//		return
//	}
//	if body == nil {
//		msg := "json对象nil"
//		glog.Error(msg)
//		http.Error(w, "json对象nil", http.StatusInternalServerError)
//		return
//	}
//	glog.Debugf("客户端生成参数:%+v", body)
//	if utils2.IsURL(body.BinPath) {
//		if this.githubProxys != nil {
//			var urls []string
//			for _, proxy := range this.githubProxys {
//				newUrl := fmt.Sprintf("%s%s", proxy, body.BinPath)
//				urls = append(urls, newUrl)
//			}
//			dstPath := utils2.DownloadFileWithCancelByUrls(urls)
//			body.BinPath = dstPath
//		} else {
//			dstPath, err := utils2.DownloadFileWithCancel(ctx, body.BinPath)
//			if err != nil {
//				msg := fmt.Errorf("下载文件失败～%v", err)
//				glog.Error(msg)
//				http.Error(w, msg.Error(), http.StatusNotImplemented)
//				return
//			}
//			body.BinPath = dstPath
//		}
//	}
//	if utils2.IsURL(body.BinUrl) {
//		if this.githubProxys != nil {
//			var urls []string
//			for _, proxy := range this.githubProxys {
//				newUrl := fmt.Sprintf("%s%s", proxy, body.BinUrl)
//				urls = append(urls, newUrl)
//			}
//			dstPath := utils2.DownloadFileWithCancelByUrls(urls)
//			body.BinPath = dstPath
//		} else {
//			dstPath, err := utils2.DownloadFileWithCancel(ctx, body.BinUrl)
//			if err != nil {
//				msg := fmt.Errorf("下载文件失败～%v", err)
//				glog.Error(msg)
//				http.Error(w, msg.Error(), http.StatusNotImplemented)
//				return
//			}
//			body.BinPath = dstPath
//		}
//	}
//	if body.User.User == "" {
//		msg := fmt.Errorf("用户名空")
//		glog.Error(msg)
//		http.Error(w, msg.Error(), http.StatusBadGateway)
//		return
//	}
//	binPath := body.BinPath
//	if binPath == "" {
//		msg := fmt.Errorf("bin文件路径空")
//		glog.Error(msg)
//		http.Error(w, msg.Error(), http.StatusServiceUnavailable)
//		return
//	}
//	glog.Infof("binPath: %s %+v\n", binPath, body)
//	tpl, err := os.Open(binPath)
//	if err != nil {
//		msg := fmt.Errorf("打开文件失败：%v", err)
//		glog.Error(msg)
//		http.Error(w, msg.Error(), http.StatusGatewayTimeout)
//		return
//	}
//	defer tpl.Close()
//
//	fileName := filepath.Base(binPath)
//	w.Header().Add("Content-Transfer-Encoding", "binary")
//	w.Header().Add("Content-Type", "application/octet-stream")
//	if stat, err := tpl.Stat(); err == nil {
//		w.Header().Add(`Content-Length`, strconv.FormatInt(stat.Size(), 10))
//	}
//	w.Header().Add(`Content-Disposition`, fmt.Sprintf("attachment; filename=\"%s\"", fileName))
//	//cfgBuffer := ukey.GetBuffer()
//	if GetCfgModel() == nil {
//		msg := fmt.Errorf("GetCfgModel is nil")
//		glog.Error(msg)
//		http.Error(w, msg.Error(), http.StatusGatewayTimeout)
//		return
//	}
//	authorization := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", GetCfgModel().Frps.WebServer.User, GetCfgModel().Frps.WebServer.Password)))
//	bindPort := GetCfgModel().Frps.BindPort
//	if body.Port > 0 {
//		bindPort = body.Port
//	}
//
//	cfgBuffer := bytes.Repeat([]byte{byte(ukey.B)}, len(ukey.GetBuffer()))
//	cfg := comm2.BufferConfig{
//		Addr:          body.Addr,
//		ApiPort:       body.ApiPort,
//		Authorization: authorization,
//		Port:          bindPort,
//		ID:            body.User.ID,
//		User:          body.User.User,
//		Token:         body.User.Token,
//		Comment:       body.User.Comment,
//		Ports:         body.User.Ports,
//		Domains:       body.User.Domains,
//		Subdomains:    body.User.Subdomains,
//		Proxy:         body.Proxy,
//		WebServer:     body.WebServer,
//	}
//
//	glog.Infof("BufferConfig: %+v", cfg)
//	cfgNewBytes, err := ukey.GenConfig(cfg, false)
//	if err != nil {
//		msg := fmt.Errorf("文件签名失败：%v", err)
//		glog.Error(msg)
//		http.Error(w, msg.Error(), http.StatusHTTPVersionNotSupported)
//		return
//	}
//
//	//err = frpc.TestLoadBuffer(cfgNewBytes)
//	//glog.Infof("TestLoadBuffer: %+v\n", err)
//
//	dstFile := filepath.Join(glog.GetCrossPlatformDataDir("temp", utils2.SecureRandomID()), fileName)
//	outFile, err := os.Create(dstFile)
//	if err != nil {
//		_ = utils2.DeleteAll(dstFile, "创建失败，删除")
//		http.Error(w, fmt.Errorf("创建失败：%v", err).Error(), http.StatusHTTPVersionNotSupported)
//		return
//	}
//	defer outFile.Close()
//	defer utils2.DeleteAll(dstFile, "gen file")
//
//	prevBuffer := make([]byte, 0)
//	for {
//		thisBuffer := make([]byte, 1024)
//		n, err := tpl.Read(thisBuffer)
//		thisBuffer = thisBuffer[:n]
//		tempBuffer := append(prevBuffer, thisBuffer...)
//		bufIndex := bytes.Index(tempBuffer, cfgBuffer)
//		if bufIndex > -1 {
//			tempBuffer = bytes.Replace(tempBuffer, cfgBuffer, cfgNewBytes, -1)
//		}
//		//w.Write(tempBuffer[:len(prevBuffer)])
//		outFile.Write(tempBuffer[:len(prevBuffer)])
//		prevBuffer = tempBuffer[len(prevBuffer):]
//		if err != nil {
//			break
//		}
//	}
//	if len(prevBuffer) > 0 {
//		//w.Write(prevBuffer)
//		outFile.Write(prevBuffer)
//		prevBuffer = nil
//	}
//	http.ServeFile(w, r, dstFile)
//}

func (this *frps) OnFrpcConfigExport(fileName string) (error, string) {
	userDir, err := utils.GetUserDir()
	if err != nil {
		return err, ""
	}
	tempDir := filepath.Join(glog.GetCrossPlatformDataDir(), "user")
	_ = utils2.EnsureDir(tempDir)
	zipFilePath := filepath.Join(tempDir, fileName)
	err = utils.Zip(userDir, zipFilePath)
	return err, zipFilePath
}

func (this *frps) apiClientUserExport(w http.ResponseWriter, r *http.Request) {
	res, f := comm2.Response(r)
	defer f(w)
	users, err := utils.GetDataByJson[[]struct {
		User string `json:"user"`
		ID   string `json:"id"`
	}](r)
	if err != nil {
		res.Err(err)
		return
	}
	userDir, err := utils.GetUserDir()
	if err != nil {
		res.Err(err)
		return
	}
	var zipFilePath string
	fileName := fmt.Sprintf("user_%s.zip", utils.GetFileNameByTime())
	tempDir := glog.GetCrossPlatformDataDir("user")
	_ = utils2.EnsureDir(tempDir)
	zipFilePath = filepath.Join(tempDir, fileName)
	if users != nil && len(*users) > 0 {
		var ids []string
		for _, u := range *users {
			ids = append(ids, filepath.Join(userDir, u.ID+".json"))
		}
		err = utils.ZipFiles(zipFilePath, ids)
	} else {
		err = utils.Zip(userDir, zipFilePath)
	}

	if err != nil {
		res.Err(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer utils2.Delete(zipFilePath, "用户配置")
	tpl, err := os.Open(zipFilePath)
	if err != nil {
		res.Err(fmt.Errorf("打开文件失败：%v", err))
		bb, err := json.Marshal(res)
		if err != nil {
			glog.Errorf("marshal result error: %v", err)
			w.WriteHeader(400)
			return
		}
		w.Write(bb)
		return
	}
	defer tpl.Close()

	w.Header().Add("Content-Transfer-Encoding", "binary")
	w.Header().Add("Content-Type", "application/octet-stream")
	if stat, err := tpl.Stat(); err == nil {
		w.Header().Add(`Content-Length`, strconv.FormatInt(stat.Size(), 10))
	}
	w.Header().Add(`Content-Disposition`, fmt.Sprintf("attachment; filename=\"%s\"", fileName))

	prevBuffer := make([]byte, 0)
	for {
		thisBuffer := make([]byte, 1024)
		n, err := tpl.Read(thisBuffer)
		thisBuffer = thisBuffer[:n]
		tempBuffer := append(prevBuffer, thisBuffer...)
		w.Write(tempBuffer[:len(prevBuffer)])
		prevBuffer = tempBuffer[len(prevBuffer):]
		if err != nil {
			break
		}
	}
	if len(prevBuffer) > 0 {
		w.Write(prevBuffer)
		prevBuffer = nil
	}
}

func (this *frps) OnFrpcConfigImport(dstFilePath string) error {
	userDir, err := utils.GetUserDir()
	if err != nil {
		glog.Error(err)
		return err
	}
	err = utils.UnzipToRoot(dstFilePath, userDir, true)
	if err == nil {
		utils.Delete(dstFilePath, "用户文件")
		glog.Info("解压成功", userDir)
	}
	return err
}

func (this *frps) apiClientUserImport(w http.ResponseWriter, r *http.Request) {
	res := &comm2.GeneralResponse{Code: 0}
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		res.Error("body can't be empty")
		glog.Error(res.Msg)
		return
	}
	// 获取上传的文件
	file, handler, err := r.FormFile("file")
	if err != nil {
		res.Error("body no file")
		return
	}
	defer file.Close()

	userDir, err := utils.GetUserDir()
	if err != nil {
		res.Err(err)
		return
	}

	glog.Info(handler.Filename)
	ext := strings.ToLower(filepath.Ext(handler.Filename)) // 统一转为小写
	switch ext {
	case ".zip":
		dstFilePath := filepath.Join(os.TempDir(), handler.Filename)
		dst, err := os.Create(dstFilePath)
		if err != nil {
			res.Error(fmt.Sprintf("create file %s error: %v", handler.Filename, err))
			return
		}
		buf := this.upgrade.GetBuffer().Get().([]byte)
		defer this.upgrade.GetBuffer().Put(buf)
		_, err = io.CopyBuffer(dst, file, buf)
		dst.Close()

		//err = utils.SaveFile(file, handler.Size, dstFilePath)
		if err != nil {
			res.Error(err.Error())
			return
		}
		err = utils.UnzipToRoot(dstFilePath, userDir, true)
		if err == nil {
			utils.Delete(dstFilePath, "用户文件")
			glog.Info("解压成功", userDir)
		}
		break
	case ".json":
		dstFilePath := filepath.Join(userDir, handler.Filename)
		dst, err := os.Create(dstFilePath)
		if err != nil {
			res.Error(fmt.Sprintf("create file %s error: %v", handler.Filename, err))
			return
		}
		buf := this.upgrade.GetBuffer().Get().([]byte)
		defer this.upgrade.GetBuffer().Put(buf)
		_, err = io.CopyBuffer(dst, file, buf)
		dst.Close()
		//err = utils.SaveFile(file, handler.Size, dstFilePath)
		if err != nil {
			res.Error(err.Error())
			return
		}
		glog.Info("导入成功", dstFilePath)
		break
	}
}

//func (this *frps) apiClientToml(w http.ResponseWriter, r *http.Request) {
//	res := &comm2.GeneralResponse{Code: 0}
//
//	ctx, cancel := context.WithCancel(context.Background())
//	defer cancel()
//
//	body, err := utils.GetDataByJson[struct {
//		BinPath   string               `json:"binPath"`
//		BinUrl    string               `json:"binUrl"`
//		Addr      string               `json:"addr"`
//		Port      int                  `json:"port"`
//		ApiPort   int                  `json:"apiPort"`
//		User      model2.User          `json:"user"`
//		Proxy     *v1.TypedProxyConfig `json:"proxy"`
//		WebServer *v1.WebServerConfig  `json:"webserver"`
//	}](r)
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		glog.Error("GetDataByJson", err)
//		return
//	}
//	if body == nil {
//		res.Err(errors.New("body is nil"))
//		w.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//	if body.BinUrl != "" && utils2.IsURL(body.BinUrl) {
//		dstPath, err1 := utils2.DownloadFileWithCancel(ctx, body.BinUrl)
//		if err1 == nil {
//			body.BinPath = dstPath
//		}
//	}
//
//	fileName := fmt.Sprintf("%s.%s.frpc.toml", body.Addr, body.User.User)
//	if GetCfgModel() == nil {
//		msg := fmt.Errorf("GetCfgModel is nil")
//		glog.Error(msg)
//		http.Error(w, msg.Error(), http.StatusGatewayTimeout)
//		return
//	}
//	authorization := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", GetCfgModel().Frps.WebServer.User, GetCfgModel().Frps.WebServer.Password)))
//	bindPort := GetCfgModel().Frps.BindPort
//	if body.Port > 0 {
//		bindPort = body.Port
//	}
//	ccc := v1.ClientCommonConfig{
//		ServerAddr: body.Addr,
//		ServerPort: bindPort,
//		User:       body.User.User,
//		//frps 网页后台，生成frpc toml配置信息
//		Metadatas: frp.GetMetadatas(body.User.Token, body.User.ID, fmt.Sprintf("%d", body.ApiPort), authorization),
//	}
//	if body.WebServer != nil && body.WebServer.Port != 0 && body.WebServer.User != "" && body.WebServer.Password != "" && body.WebServer.Addr != "" {
//		ccc.WebServer = *body.WebServer
//	}
//
//	var proxies []v1.TypedProxyConfig
//	if body.Proxy != nil && comm2.HasProxyes(body.Proxy) {
//		proxies = append(proxies, *body.Proxy)
//	}
//
//	cc := v1.ClientConfig{
//		ClientCommonConfig: ccc,
//		Proxies:            proxies,
//	}
//	cfg := &frpc.CfgModel{
//		Frpc: cc,
//	}
//	buffer := utils.ObjectToTomlText(cfg.Frpc)
//
//	//sb := strings.Builder{}
//	//sb.WriteString(fmt.Sprintf("serverAddr = \"%s\"\n", body.Addr))
//	//sb.WriteString(fmt.Sprintf("serverPort = %d\n", bindPort))
//	//sb.WriteString(fmt.Sprintf("user = \"%s\"\n", body.User.User))
//	//sb.WriteString(fmt.Sprintf("metadatas.token = \"%s\"\n", body.User.Token))
//	//sb.WriteString(fmt.Sprintf("metadatas.id = \"%s\"\n", body.User.SseId))
//	//size := sb.Len()
//	//
//	w.Header().Add("Content-Transfer-Encoding", "binary")
//	w.Header().Add("Content-Type", "application/octet-stream")
//	w.Header().Add(`Content-Length`, strconv.Itoa(len(buffer)))
//	w.Header().Add(`Content-Disposition`, fmt.Sprintf("attachment; filename=\"%s\"", fileName))
//	//w.Write([]byte(sb.String()))
//	_, _ = w.Write(buffer)
//}

func (this *frps) apiConfigUpload(w http.ResponseWriter, r *http.Request) {
	res, f := comm2.Response(r)
	defer f(w)
	fpath := filepath.Join(glog.GetCrossPlatformDataDir("obj"), "cloudApi.dat")
	switch r.Method {
	case "GET", "get":
		if !utils2.FileExists(fpath) {
			res.Result(100, "接口设置～", this.cloudApi)
		} else {
			obj, err := utils.LoadWithGob[model.CloudApi](fpath)
			if err != nil {
				res.Err(err)
			} else {
				this.cloudApi = &obj
				glog.Debug("LoadWithGob:", obj)
				err = utils.Export(obj)
				if err == nil {
					res.Ok("上传成功")
					return
				}
				res.Err(err)
			}
		}
		break
	case "POST", "post":
		body, err := utils.GetDataByJson[model.CloudApi](r)
		if err != nil {
			res.Err(err)
			return
		}
		glog.Debugf("参数：%+v", body)
		if body.Addr != "" {
			err = utils.SaveWithGob[model.CloudApi](*body, fpath)
			if err != nil {
				res.Err(err)
				return
			}
			glog.Debug("SaveWithGob", fpath)
			err = utils.Export(*body)
			if err == nil {
				res.Ok("上传成功")
				return
			}
		} else {
			res.Error("cloud api无效")
		}
		break
	default:
		break
	}
}

func (this *frps) apiConfigUpgrade(w http.ResponseWriter, r *http.Request) {
	res, f := comm2.Response(r)
	defer f(w)
	fpath := filepath.Join(glog.GetCrossPlatformDataDir("obj"), "cloudApi.dat")
	switch r.Method {
	case "GET", "get":
		glog.Debug("同步配置...get")
		if !utils2.FileExists(fpath) {
			res.Result(100, "接口设置～", this.cloudApi)
		} else {
			obj, err := utils.LoadWithGob[model.CloudApi](fpath)
			if err != nil {
				res.Err(err)
			} else {
				this.cloudApi = &obj
				glog.Debug("LoadWithGob:", obj)
				err = utils.Import(obj)
				if err == nil {
					res.Ok("更新成功")
					return
				}
				res.Err(err)
			}
		}
		break
	case "POST", "post":
		glog.Debug("同步配置...post")
		body, err := utils.GetDataByJson[model.CloudApi](r)
		if err != nil {
			res.Err(err)
			return
		}
		glog.Debugf("参数：%+v", body)
		if body.Addr != "" {
			err = utils.SaveWithGob[model.CloudApi](*body, fpath)
			if err != nil {
				res.Err(err)
				return
			}
			glog.Debug("SaveWithGob", fpath)
			err = utils.Import(*body)
			if err == nil {
				res.Ok("更新成功")
				return
			}
		} else {
			res.Error("cloud api无效")
		}
		break
	default:
		break
	}
}

func (this *frps) apiClientUpload(w http.ResponseWriter, r *http.Request) {
	res, f := comm2.Response(r)
	defer f(w)
	//err := r.ParseMultipartForm(32 << 20)
	//if err != nil {
	//	res.Error("body can't be empty")
	//	glog.Error(res.Msg)
	//	return
	//}
	// 获取上传的文件
	file, handler, err := r.FormFile("file")
	if err != nil {
		res.Error("body no file")
		return
	}
	defer file.Close()
	binPath, err := os.Executable()
	if err != nil {
		res.Error(fmt.Sprintf("获取当前可执行文件路径出错: %v\n", err))
		glog.Error(res.Msg)
		return
	}
	binDir := filepath.Dir(binPath)
	clientsDir := filepath.Join(binDir, "clients")
	err = utils.EnsureDir(clientsDir)
	if err != nil {
		res.Error(fmt.Sprintf("文件夹创建失败: %v\n", err))
		glog.Error(res.Msg)
		return
	}

	dstFilePath := filepath.Join(clientsDir, handler.Filename)
	//dstFilePath 名称为上传文件的原始名称
	dst, err := os.Create(dstFilePath)
	if err != nil {
		res.Error(fmt.Sprintf("create file %s error: %v", handler.Filename, err))
		return
	}
	buf := this.upgrade.GetBuffer().Get().([]byte)
	defer this.upgrade.GetBuffer().Put(buf)
	_, err = io.CopyBuffer(dst, file, buf)
	dst.Close()
	if err != nil {
		res.Error(err.Error())
		glog.Error(res.Msg)
		return
	}
	glog.Println("客户端路径", clientsDir)
	glog.Println("文件上传成功", dstFilePath)
	err = utils.UnzipToRoot(dstFilePath, clientsDir, true)
	if err != nil {
		res.Error(err.Error())
		glog.Error(res.Msg)
		return
	} else {
		utils.Delete(dstFilePath)
	}
	res.Ok("文件上传成功～")
}

func (this *frps) apiFrpsGen(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	body, err := utils.GetDataByJson[struct {
		BindPort  int      `json:"bindPort"`
		AdminAddr string   `json:"adminAddr"`
		AdminPort int      `json:"adminPort"`
		User      string   `json:"user"`
		Pass      string   `json:"pass"`
		Ops       []string `json:"ops"`
	}](r)
	if err != nil {
		glog.Error("解析Json对象失败", err)
		return
	}
	if body == nil {
		msg := "json对象nil"
		glog.Error(msg)
		http.Error(w, "json对象nil", http.StatusInternalServerError)
		return
	}
	glog.Debugf("body:%+v\n", body)

	if body.Ops == nil {
		msg := "body.Ops nil"
		glog.Error(msg)
		http.Error(w, "body.Ops nil", http.StatusInternalServerError)
		return
	}
	binUrl := this.getFrpsDownloadUrls(body.Ops[0], body.Ops[1])

	if binUrl == "" {
		msg := "frps download url is nil"
		glog.Error(msg)
		http.Error(w, "frps download url is nil", http.StatusInternalServerError)
		return
	}
	var binPath string
	if this.githubProxys != nil {
		var urls []string
		for _, proxy := range this.githubProxys {
			newUrl := fmt.Sprintf("%s%s", proxy, binUrl)
			urls = append(urls, newUrl)
		}
		binPath = utils2.DownloadFileWithCancelByUrls(urls)
	} else {
		dstPath, e := utils2.DownloadFileWithCancel(ctx, binUrl)
		if e != nil {
			msg := fmt.Errorf("下载文件失败～%v", e)
			glog.Error(msg)
			http.Error(w, msg.Error(), http.StatusNotImplemented)
			return
		}
		binPath = dstPath
	}
	if binPath == "" {
		msg := fmt.Errorf("bin文件路径空")
		glog.Error(msg)
		http.Error(w, msg.Error(), http.StatusServiceUnavailable)
		return
	}
	glog.Infof("binPath: %s %+v\n", binPath, body)
	tpl, err := os.Open(binPath)
	if err != nil {
		msg := fmt.Errorf("打开文件失败：%v", err)
		glog.Error(msg)
		http.Error(w, msg.Error(), http.StatusGatewayTimeout)
		return
	}
	defer tpl.Close()

	fileName := filepath.Base(binPath)
	w.Header().Add("Content-Transfer-Encoding", "binary")
	w.Header().Add("Content-Type", "application/octet-stream")
	if stat, err := tpl.Stat(); err == nil {
		w.Header().Add(`Content-Length`, strconv.FormatInt(stat.Size(), 10))
	}
	w.Header().Add(`Content-Disposition`, fmt.Sprintf("attachment; filename=\"%s\"", fileName))

	cfg := &CfgModel{
		Frps: v1.ServerConfig{
			BindPort: body.BindPort,
			WebServer: v1.WebServerConfig{
				User:     body.User,
				Password: body.Pass,
				Port:     body.AdminPort,
				Addr:     body.AdminAddr,
			},
			Log: v1.LogConfig{
				To:      filepath.Join(glog.GetCrossPlatformDataDir("log"), "frps.log"),
				MaxDays: 3,
				Level:   "error",
			},
		},
	}
	cfgNewBytes, err := ukey.GenConfig(cfg, false)
	if err != nil {
		msg := fmt.Errorf("文件签名失败：%v", err)
		glog.Error(msg)
		http.Error(w, msg.Error(), http.StatusHTTPVersionNotSupported)
		return
	}
	glog.Debugf("配置信息:%+v", cfgNewBytes)
	cfgBuffer := bytes.Repeat([]byte{byte(ukey.B)}, len(ukey.GetBuffer()))
	prevBuffer := make([]byte, 0)

	dstFile := filepath.Join(glog.GetCrossPlatformDataDir("temp", utils2.SecureRandomID()), fileName)
	outFile, err := os.Create(dstFile)
	if err != nil {
		_ = utils2.DeleteAll(dstFile, "创建失败，删除")
		http.Error(w, fmt.Errorf("创建失败：%v", err).Error(), http.StatusHTTPVersionNotSupported)
		return
	}
	defer outFile.Close()
	defer utils2.DeleteAll(dstFile, "gen file")

	for {
		thisBuffer := make([]byte, 1024)
		n, err := tpl.Read(thisBuffer)
		thisBuffer = thisBuffer[:n]
		tempBuffer := append(prevBuffer, thisBuffer...)
		bufIndex := bytes.Index(tempBuffer, cfgBuffer)
		if bufIndex > -1 {
			tempBuffer = bytes.Replace(tempBuffer, cfgBuffer, cfgNewBytes, -1)
		}
		//s, e := w.Write(tempBuffer[:len(prevBuffer)])
		s, e := outFile.Write(tempBuffer[:len(prevBuffer)])
		if e != nil {
			glog.Errorf("size:%v err:%v", s, e)
		}
		prevBuffer = tempBuffer[len(prevBuffer):]
		if err != nil {
			glog.Errorf("tpl.Read err:%v", err)
			break
		}
	}
	if len(prevBuffer) > 0 {
		//s, e := w.Write(prevBuffer)
		s, e := outFile.Write(prevBuffer)
		if e != nil {
			glog.Errorf("size:%v err:%v", s, e)
		}
		prevBuffer = nil
	}

	http.ServeFile(w, r, dstFile)
}
