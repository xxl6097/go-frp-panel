package frpc

import (
	"fmt"
	httppkg "github.com/fatedier/frp/pkg/util/http"
	"github.com/xxl6097/glog/glog"
	"github.com/xxl6097/go-frp-panel/internal/comm"
	"github.com/xxl6097/go-frp-panel/pkg/utils"
	"github.com/xxl6097/go-service/gservice/gore"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func (this *frpc) adminHandlers(helper *httppkg.RouterRegisterHelper) {
	subRouter := helper.Router.NewRoute().Name("admin").Subrouter()
	subRouter.Use(helper.AuthMiddleware)
	staticPrefix := "/log/"
	baseDir := os.TempDir()
	subRouter.PathPrefix(staticPrefix).Handler(http.StripPrefix(staticPrefix, http.FileServer(http.Dir(baseDir))))

	// apis
	subRouter.HandleFunc("/api/version", this.upgrade.ApiVersion).Methods("GET")
	subRouter.HandleFunc("/api/upgrade", this.upgrade.ApiUpdate).Methods("POST")
	subRouter.HandleFunc("/api/upgrade", this.upgrade.ApiUpdate).Methods("PUT")
	subRouter.HandleFunc("/api/restart", this.upgrade.ApiRestart).Methods("GET")

	subRouter.HandleFunc("/api/client/create", this.apiClientCreate).Methods("PUT")
	subRouter.HandleFunc("/api/client/create", this.apiClientCreate).Methods("POST")
	subRouter.HandleFunc("/api/client/delete", this.apiClientDelete).Methods("GET")
	subRouter.HandleFunc("/api/client/status", this.apiClientStatus).Methods("GET")
	subRouter.HandleFunc("/api/client/list", this.apiClientList).Methods("GET")
	subRouter.HandleFunc("/api/client/config/get", this.apiClientConfigGet).Methods("GET")
	subRouter.HandleFunc("/api/client/config/set", this.apiClientConfigSet).Methods("POST")
}

func (this *frpc) apiClientCreate(w http.ResponseWriter, r *http.Request) {
	res, f := comm.Response(r)
	defer f(w)
	var newFilePath string
	binpath, err := os.Executable()
	if err != nil {
		res.Err(fmt.Errorf("get executable path err: %v", err))
		glog.Error(res.Msg)
		return
	}
	cfgDir := filepath.Join(filepath.Dir(binpath), "config")
	if err = utils.DirCheck(cfgDir); err != nil {
		res.Err(fmt.Errorf("check config dir err: %v", err))
		glog.Error(res.Msg)
		return
	}

	switch r.Method {
	case "PUT", "put":
		body, err := utils.GetDataByJson[struct {
			Name string `json:"name"`
			Toml string `json:"toml"`
		}](r)
		if body == nil {
			res.Error("body is empty")
			glog.Error(res.Msg)
			return
		}
		cfgFilePath := filepath.Join(cfgDir, body.Name)
		if gore.FileExists(cfgFilePath) {
			res.Err(fmt.Errorf("客户端已经存在"))
			glog.Error(res.Msg)
			return
		}
		err = utils.WriteToml(cfgFilePath, []byte(body.Toml))
		if err != nil {
			res.Err(fmt.Errorf("write http body err: %v", err))
			glog.Error(res.Msg)
			return
		}
		newFilePath = cfgFilePath
		break
	case "POST", "post":
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
		dstFilePath := filepath.Join(cfgDir, handler.Filename)
		if gore.FileExists(dstFilePath) {
			res.Err(fmt.Errorf("客户端已经存在"))
			glog.Error(res.Msg)
			return
		}
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
		//err = utils.SaveFile(file, handler.Size, dstFilePath)
		if err != nil {
			res.Error(err.Error())
			return
		}
		newFilePath = dstFilePath
		break
	default:
		res.Error("位置请求方法")
	}

	if newFilePath != "" {
		err := this.runClient(newFilePath)
		if err != nil {
			res.Err(fmt.Errorf("run client err: %v", err))
			glog.Error(res.Msg)
			return
		}
		res.Ok("创建成功～")
	}

}

func (this *frpc) apiClientCreatePOST(w http.ResponseWriter, r *http.Request) {
	res, f := comm.Response(r)
	defer f(w)
	body, err := utils.GetDataByJson[struct {
		Name string `json:"name"`
		Toml string `json:"toml"`
	}](r)
	if body == nil {
		res.Error("body is empty")
		glog.Error(res.Msg)
		return
	}
	binpath, err := os.Executable()
	if err != nil {
		res.Err(fmt.Errorf("get executable path err: %v", err))
		glog.Error(res.Msg)
		return
	}
	cfgDir := filepath.Join(filepath.Dir(binpath), "config")
	if err = utils.DirCheck(cfgDir); err != nil {
		res.Err(fmt.Errorf("check config dir err: %v", err))
		glog.Error(res.Msg)
		return
	}
	cfgFilePath := filepath.Join(cfgDir, body.Name)
	if gore.FileExists(cfgFilePath) {
		res.Err(fmt.Errorf("客户端已经存在"))
		glog.Error(res.Msg)
		return
	}
	err = utils.WriteToml(cfgFilePath, []byte(body.Toml))
	if err != nil {
		res.Err(fmt.Errorf("write http body err: %v", err))
		glog.Error(res.Msg)
		return
	}
	err = this.runClient(cfgFilePath)
	if err != nil {
		res.Err(fmt.Errorf("run client err: %v", err))
		glog.Error(res.Msg)
		return
	}
	res.Ok("创建成功～")
}

func (this *frpc) apiClientDelete(w http.ResponseWriter, r *http.Request) {
	res, f := comm.Response(r)
	defer f(w)
	cfgName := r.URL.Query().Get("name")
	if cfgName == "" {
		res.Error("cfg file path is empty")
		return
	}
	binpath, err := os.Executable()
	if err != nil {
		res.Err(fmt.Errorf("get executable path err: %v", err))
		return
	}
	cfgDir := filepath.Join(filepath.Dir(binpath), "config")
	cfgFilePath := filepath.Join(cfgDir, cfgName)
	err = os.Remove(cfgFilePath)
	if err != nil {
		res.Err(fmt.Errorf("delete config file err: %v", err))
		return
	}
	err = this.deleteClient(cfgFilePath)
	if err != nil {
		res.Err(fmt.Errorf("delete client err: %v", err))
		return
	}
	res.Ok("删除成功～")
}

func (this *frpc) apiClientStatus(w http.ResponseWriter, r *http.Request) {
	res, f := comm.Response(r)
	defer f(w)
	cfgName := r.URL.Query().Get("name")
	if cfgName == "" {
		res.Error("cfg file path is empty")
		return
	}
	binpath, err := os.Executable()
	if err != nil {
		res.Err(fmt.Errorf("get executable path err: %v", err))
		return
	}
	cfgDir := filepath.Join(filepath.Dir(binpath), "config")
	cfgFilePath := filepath.Join(cfgDir, cfgName)
	buf, err := this.statusClient(cfgFilePath)
	if err != nil {
		res.Err(fmt.Errorf("get status client err: %v", err))
		return
	}
	res.Raw = buf
}

func (this *frpc) apiClientList(w http.ResponseWriter, r *http.Request) {
	res, f := comm.Response(r)
	defer f(w)
	binpath, err := os.Executable()
	if err != nil {
		res.Err(fmt.Errorf("get executable path err: %v", err))
		return
	}
	cfgDir := filepath.Join(filepath.Dir(binpath), "config")
	files, err := os.ReadDir(cfgDir)
	if err != nil {
		res.Err(fmt.Errorf("read config dir err: %v", err))
		glog.Error(res.Msg)
		return
	}

	var names []comm.Option
	for _, f := range files {
		ext := strings.ToLower(filepath.Ext(f.Name()))
		if !f.IsDir() && ext == ".toml" {
			names = append(names, comm.Option{
				Label: f.Name(),
				Value: f.Name(),
			})
		}
	}
	res.Sucess("客户端列表获取成功", names)
}

func (this *frpc) apiClientConfigGet(w http.ResponseWriter, r *http.Request) {
	res, f := comm.Response(r)
	defer f(w)
	cfgName := r.URL.Query().Get("name")
	if cfgName == "" {
		res.Error("cfg file path is empty")
		glog.Error(res.Msg)
		return
	}
	binpath, err := os.Executable()
	if err != nil {
		res.Err(fmt.Errorf("get executable path err: %v", err))
		glog.Error(res.Msg)
		return
	}
	cfgDir := filepath.Join(filepath.Dir(binpath), "config")
	cfgFilePath := filepath.Join(cfgDir, cfgName)
	body, err := utils.ReadToml(cfgFilePath)
	if err != nil {
		res.Err(fmt.Errorf("write http body err: %v", err))
		glog.Error(res.Msg)
		return
	}
	res.Raw = body
}

func (this *frpc) apiClientConfigSet(w http.ResponseWriter, r *http.Request) {
	res, f := comm.Response(r)
	defer f(w)
	body, err := utils.GetDataByJson[struct {
		Name string `json:"name"`
		Toml string `json:"toml"`
	}](r)
	if body == nil {
		res.Error("body is empty")
		glog.Error(res.Msg)
		return
	}
	binpath, err := os.Executable()
	if err != nil {
		res.Err(fmt.Errorf("get executable path err: %v", err))
		glog.Error(res.Msg)
		return
	}
	cfgDir := filepath.Join(filepath.Dir(binpath), "config")
	cfgFilePath := filepath.Join(cfgDir, body.Name)
	if !gore.FileExists(cfgFilePath) {
		res.Err(fmt.Errorf("客户端不存在: %v", err))
		glog.Error(res.Msg)
		return
	}
	err = utils.WriteToml(cfgFilePath, []byte(body.Toml))
	if err != nil {
		res.Err(fmt.Errorf("write http body err: %v", err))
		return
	}
	err = this.updateClient(cfgFilePath)
	if err != nil {
		res.Err(fmt.Errorf("run client err: %v", err))
		return
	}
	res.Ok("更新成功～")
}
