package frps

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	v1 "github.com/fatedier/frp/pkg/config/v1"
	httppkg "github.com/fatedier/frp/pkg/util/http"
	"github.com/xxl6097/glog/glog"
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

func (this *frps) userHandlers(helper *httppkg.RouterRegisterHelper) {
	subRouter := helper.Router.NewRoute().Name("user").Subrouter()
	subRouter.Use(helper.AuthMiddleware)
	// apis
	subRouter.HandleFunc("/api/token/add", this.apiUserCreate).Methods("POST")
	subRouter.HandleFunc("/api/token/del", this.apiUserDelete).Methods("POST")
	subRouter.HandleFunc("/api/token/chg", this.apiUserUpdate).Methods("POST")
	subRouter.HandleFunc("/api/token/all", this.apiUserAll).Methods("GET")

	subRouter.HandleFunc("/api/client/get", this.apiClientGet).Methods("GET")
	subRouter.HandleFunc("/api/client/gen", this.apiClientGen).Methods("POST")
	subRouter.HandleFunc("/api/frps/get", this.apiFrpsGet).Methods("GET")
	subRouter.HandleFunc("/api/frps/gen", this.apiFrpsGen).Methods("POST")
	subRouter.HandleFunc("/api/client/toml", this.apiClientToml).Methods("POST")
	subRouter.HandleFunc("/api/client/user/import", this.apiClientUserImport).Methods("POST")
	subRouter.HandleFunc("/api/client/user/export", this.apiClientUserExport).Methods("POST")
	subRouter.HandleFunc("/api/config/backup", this.apiConfigBackup)
	subRouter.HandleFunc("/api/client/upload", this.apiClientUpload).Methods("POST")
}

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
	u, err := utils.GetDataByJson[User](r)
	if err != nil {
		res.Err(err)
		glog.Errorf("%v decode user err: %+v", err, u)
		return
	}
	if u == nil {
		res.Error("token is nil")
		return
	}
	err = u.CreateUser()
	if err != nil {
		res.Err(err)
		glog.Errorf("%v create user err: %+v", err, u)
		return
	}
	res.Ok("密钥创建成功")
}

func (this *frps) apiUserDelete(w http.ResponseWriter, r *http.Request) {
	res, f := comm2.Response(r)
	defer f(w)
	users, err := utils.GetDataByJson[[]struct {
		User string `json:"user"`
	}](r)
	if err != nil {
		res.Err(err)
		glog.Error(err)
		return
	}
	if users == nil {
		res.Error("tokens is nil")
		return
	}
	for _, u := range *users {
		err = DeleteUser(u.User)
	}
	//err = this.repo.Delete(u.User)
	//if err != nil {
	//	res.Err(err)
	//	return
	//}
	res.Ok("密钥删除成功")
}

func (this *frps) apiUserUpdate(w http.ResponseWriter, r *http.Request) {
	res, f := comm2.Response(r)
	defer f(w)
	u, err := utils.GetDataByJson[User](r)
	if err != nil {
		res.Err(err)
		return
	}
	if u == nil {
		res.Error("token is nil")
		return
	}
	glog.Printf("%+v\n", u)
	//userFilePath := filepath.Join(this.workDir, "user", fmt.Sprintf("%s.json", u.User))
	//if gore.FileExists(userFilePath) {
	//	os.Remove(userFilePath)
	//}
	//err = u.CreateUser(userFilePath)
	err = u.UpdateUser()

	if err != nil {
		res.Err(err)
		glog.Error(err)
		return
	}
	res.Ok("密钥更新成功")
	a, _ := GetUserAll()
	fmt.Printf("结果：%+v\n", a)
}
func (this *frps) apiUserAll(w http.ResponseWriter, r *http.Request) {
	res, f := comm2.Response(r)
	defer f(w)
	datas, err := GetUserAll()
	if err != nil {
		res.Error("无数据")
		glog.Error(err)
		return
	}
	res.Sucess("全部数据获取成功", datas)
	glog.Infof("%+v\n", datas)
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
	glog.Infof("扫描路径:%s", configPath)
	nodes := utils.GetNodes(configPath)
	if nodes == nil || len(nodes) == 0 {
		nodes = utils.ToTree("", this.frpcGithubDownloadUrls)
	}
	res.Data = nodes
	glog.Infof("扫描结果:%v", res.Data)
}

func (this *frps) apiFrpsGet(w http.ResponseWriter, r *http.Request) {
	res, f := comm2.Response(r)
	defer f(w)
	res.Data = utils.ToTree("", this.frpsGithubDownloadUrls)
	glog.Infof("frpsGithubDownloadUrls:%v", this.frpsGithubDownloadUrls)
	glog.Infof("frps地址扫描:%v", res.Data)
}

func (this *frps) parseUser(data map[string]interface{}) {
	glog.Println(data)
	u := User{
		User:       data["user"].(string),
		Token:      data["token"].(string),
		Comment:    data["comment"].(string),
		Ports:      ToPorts(data["ports"].([]any)),
		Domains:    data["domains"].([]string),
		Subdomains: data["subdomains"].([]string),
		Enable:     data["enable"].(bool),
	}
	glog.Error(u)
}

func (this *frps) apiClientGen(w http.ResponseWriter, r *http.Request) {
	//res := &comm.GeneralResponse{Code: 0}

	//body1, err := io.ReadAll(r.Body)
	//if err != nil {
	//	res.Response(400, fmt.Sprintf("read request body error: %v", err))
	//	glog.Warnf("%s", res.Msg)
	//	return
	//}
	//fmt.Println(string(body1))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	body, err := utils.GetDataByJson[struct {
		BinPath string `json:"binPath"`
		BinUrl  string `json:"binUrl"`
		Addr    string `json:"addr"`
		Port    int    `json:"port"`
		User    User   `json:"user"`
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
	if utils2.IsURL(body.BinPath) {
		if this.githubProxys != nil {
			var urls []string
			for _, proxy := range this.githubProxys {
				newUrl := fmt.Sprintf("%s%s", proxy, body.BinPath)
				urls = append(urls, newUrl)
			}
			dstPath := utils2.DownloadFileWithCancelByUrls(urls)
			body.BinPath = dstPath
		} else {
			dstPath, err := utils2.DownloadFileWithCancel(ctx, body.BinPath)
			if err != nil {
				msg := fmt.Errorf("下载文件失败～%v", err)
				glog.Error(msg)
				http.Error(w, msg.Error(), http.StatusNotImplemented)
				return
			}
			body.BinPath = dstPath
		}
	}
	if utils2.IsURL(body.BinUrl) {
		if this.githubProxys != nil {
			var urls []string
			for _, proxy := range this.githubProxys {
				newUrl := fmt.Sprintf("%s%s", proxy, body.BinUrl)
				urls = append(urls, newUrl)
			}
			dstPath := utils2.DownloadFileWithCancelByUrls(urls)
			body.BinPath = dstPath
		} else {
			dstPath, err := utils2.DownloadFileWithCancel(ctx, body.BinUrl)
			if err != nil {
				msg := fmt.Errorf("下载文件失败～%v", err)
				glog.Error(msg)
				http.Error(w, msg.Error(), http.StatusNotImplemented)
				return
			}
			body.BinPath = dstPath
		}
	}
	if body.User.User == "" {
		msg := fmt.Errorf("用户名空")
		glog.Error(msg)
		http.Error(w, msg.Error(), http.StatusBadGateway)
		return
	}
	binPath := body.BinPath
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

	w.Header().Add("Content-Transfer-Encoding", "binary")
	w.Header().Add("Content-Type", "application/octet-stream")
	if stat, err := tpl.Stat(); err == nil {
		w.Header().Add(`Content-Length`, strconv.FormatInt(stat.Size(), 10))
	}
	w.Header().Add(`Content-Disposition`, fmt.Sprintf("attachment; filename=\"%s\"", filepath.Base(binPath)))
	//cfgBuffer := ukey.GetBuffer()
	bindPort := GetCfgModel().Frps.BindPort
	if body.Port > 0 {
		bindPort = body.Port
	}
	cfgBuffer := bytes.Repeat([]byte{byte(ukey.B)}, len(ukey.GetBuffer()))
	cfg := comm2.BufferConfig{
		Addr:       body.Addr,
		Port:       bindPort,
		User:       body.User.User,
		Token:      body.User.Token,
		Comment:    body.User.Comment,
		Ports:      body.User.Ports,
		Domains:    body.User.Domains,
		Subdomains: body.User.Subdomains,
	}

	cfgNewBytes, err := ukey.GenConfig(cfg, false)
	if err != nil {
		msg := fmt.Errorf("文件签名失败：%v", err)
		glog.Error(msg)
		http.Error(w, msg.Error(), http.StatusHTTPVersionNotSupported)
		return
	}
	prevBuffer := make([]byte, 0)
	for {
		thisBuffer := make([]byte, 1024)
		n, err := tpl.Read(thisBuffer)
		thisBuffer = thisBuffer[:n]
		tempBuffer := append(prevBuffer, thisBuffer...)
		bufIndex := bytes.Index(tempBuffer, cfgBuffer)
		if bufIndex > -1 {
			tempBuffer = bytes.Replace(tempBuffer, cfgBuffer, cfgNewBytes, -1)
		}
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

func (this *frps) OnFrpcConfigExport(fileName string) (error, string) {
	binpath, err := os.Executable()
	if err != nil {
		return err, ""
	}
	userDir := filepath.Join(filepath.Dir(binpath), "user")
	tempDir := filepath.Join(glog.GetCrossPlatformDataDir(), "user")
	_ = utils2.EnsureDir(tempDir)
	zipFilePath := filepath.Join(tempDir, fileName)
	err = utils.Zip(userDir, zipFilePath)
	return err, zipFilePath
}

func (this *frps) apiClientUserExport(w http.ResponseWriter, r *http.Request) {
	res := &comm2.GeneralResponse{Code: 0}
	binpath, err := os.Executable()
	if err != nil {
		res.Err(err)
		glog.Error(err)
		return
	}
	userDir := filepath.Join(filepath.Dir(binpath), "user")

	fileName := fmt.Sprintf("user_%s.zip", utils.GetFileNameByTime())
	tempDir := filepath.Join(glog.GetCrossPlatformDataDir(), "user")
	_ = utils2.EnsureDir(tempDir)
	zipFilePath := filepath.Join(tempDir, fileName)
	err = utils.Zip(userDir, zipFilePath)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		glog.Error("GetDataByJson", err)
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
	binpath, err := os.Executable()
	if err != nil {
		return err
	}
	userDir := filepath.Join(filepath.Dir(binpath), "user")
	if err = utils.DirCheck(userDir); err != nil {
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

	binpath, err := os.Executable()
	if err != nil {
		res.Err(err)
		glog.Error(binpath, err)
		return
	}

	userDir := filepath.Join(filepath.Dir(binpath), "user")

	if err = utils.DirCheck(userDir); err != nil {
		res.Err(fmt.Errorf("check config dir err: %v", err))
		glog.Error(res.Msg)
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

func (this *frps) apiClientToml(w http.ResponseWriter, r *http.Request) {
	res := &comm2.GeneralResponse{Code: 0}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	body, err := utils.GetDataByJson[struct {
		BinPath string `json:"binPath"`
		BinUrl  string `json:"binUrl"`
		Addr    string `json:"addr"`
		User    User   `json:"user"`
	}](r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		glog.Error("GetDataByJson", err)
		return
	}
	if body == nil {
		res.Err(errors.New("body is nil"))
		w.WriteHeader(http.StatusInternalServerError)
		glog.Error("body is nil")
		return
	}
	if body.BinUrl != "" && utils2.IsURL(body.BinUrl) {
		dstPath, err1 := utils2.DownloadFileWithCancel(ctx, body.BinUrl)
		if err1 == nil {
			body.BinPath = dstPath
		}
	}

	fileName := fmt.Sprintf("%s_frpc.toml", body.User.User)

	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("serverAddr = \"%s\"\n", body.Addr))
	sb.WriteString(fmt.Sprintf("serverPort = %d\n", GetCfgModel().Frps.BindPort))
	sb.WriteString(fmt.Sprintf("user = \"%s\"\n", body.User.User))
	sb.WriteString(fmt.Sprintf("metadatas.token = \"%s\"\n", body.User.Token))
	size := sb.Len()

	w.Header().Add("Content-Transfer-Encoding", "binary")
	w.Header().Add("Content-Type", "application/octet-stream")
	w.Header().Add(`Content-Length`, strconv.Itoa(size))
	w.Header().Add(`Content-Disposition`, fmt.Sprintf("attachment; filename=\"%s\"", fileName))
	w.Write([]byte(sb.String()))
}

func (this *frps) apiConfigBackup(w http.ResponseWriter, r *http.Request) {
	res, f := comm2.Response(r)
	defer f(w)
	fpath := filepath.Join(glog.GetCrossPlatformDataDir("obj"), "cloudApi.dat")
	switch r.Method {
	case "GET", "get":
		if !utils2.FileExists(fpath) {
			res.Response(100, "接口设置～")
		} else {
			obj, err := utils.LoadWithGob[model.CloudApi](fpath)
			if err != nil {
				res.Err(err)
			} else {
				glog.Debug("LoadWithGob:", obj)
				err = utils.Import(obj)
				if err == nil {
					err = utils.Export(obj)
					if err == nil {
						res.Ok("同步成功")
						return
					}
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
		}
		res.Ok("哇哈哈")
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

	w.Header().Add("Content-Transfer-Encoding", "binary")
	w.Header().Add("Content-Type", "application/octet-stream")
	if stat, err := tpl.Stat(); err == nil {
		w.Header().Add(`Content-Length`, strconv.FormatInt(stat.Size(), 10))
	}
	w.Header().Add(`Content-Disposition`, fmt.Sprintf("attachment; filename=\"%s\"", filepath.Base(binPath)))

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
	for {
		thisBuffer := make([]byte, 1024)
		n, err := tpl.Read(thisBuffer)
		thisBuffer = thisBuffer[:n]
		tempBuffer := append(prevBuffer, thisBuffer...)
		bufIndex := bytes.Index(tempBuffer, cfgBuffer)
		if bufIndex > -1 {
			tempBuffer = bytes.Replace(tempBuffer, cfgBuffer, cfgNewBytes, -1)
		}
		s, e := w.Write(tempBuffer[:len(prevBuffer)])
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
		s, e := w.Write(prevBuffer)
		if e != nil {
			glog.Errorf("size:%v err:%v", s, e)
		}
		prevBuffer = nil
	}
}
