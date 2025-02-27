package comm

import (
	"encoding/json"
	"fmt"
	"github.com/fatedier/frp/pkg/util/version"
	"github.com/xxl6097/glog/glog"
	"github.com/xxl6097/go-frp-panel/internal/comm/ukey"
	"github.com/xxl6097/go-frp-panel/pkg"
	"github.com/xxl6097/go-frp-panel/pkg/utils"
	"github.com/xxl6097/go-service/gservice/gore"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type CommApi struct {
	Install gore.Install
	Object  any
}

func (this *CommApi) ApiUpdate(w http.ResponseWriter, r *http.Request) {
	res, f := Response(r)
	defer f(w)
	var newFilePath string
	switch r.Method {
	case "PUT", "put":
		body, err := io.ReadAll(r.Body)
		if err != nil {
			res.Response(400, fmt.Sprintf("read request body error: %v", err))
			glog.Warnf("%s", res.Msg)
			return
		}
		if len(body) == 0 {
			res.Response(400, "body can't be empty")
			glog.Warnf("%s", res.Msg)
			return
		}
		newFilePath, err = utils.DownLoad(string(body))
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
		dstFilePath := filepath.Join(os.TempDir(), handler.Filename)
		//dstFilePath 名称为上传文件的原始名称
		err = utils.SaveFile(file, handler.Size, dstFilePath)
		if err != nil {
			res.Error(err.Error())
			return
		}
		newFilePath = dstFilePath
		break
	default:
		res.Error("位置请求方法")
	}
	defer utils.Delete(newFilePath, "更新文件")
	//下载和接收的最新文件 名称为上传文件的原始名称
	newBufferBytes, err := ukey.GenConfig(this.Object, false)
	if err != nil {
		res.Error(fmt.Sprintf("gen config err: %v", err))
		glog.Error(res.Msg)
		return
	}
	signFilePath, err := utils.SignAndInstall(newBufferBytes, ukey.UnInitializeBuffer(), newFilePath)
	glog.Println("签名安装完毕", err, res)
	if err != nil {
		res.Error(err.Error())
		glog.Error(res.Msg)
	} else {
		defer utils.Delete(signFilePath, "签名文件")
		err = this.Install.Upgrade(signFilePath)
		if err != nil {
			res.Error(fmt.Sprintf("更新失败～%v", err))
			return
		}
		res.Ok("升级成功～")
	}
}

func (this *CommApi) ApiRestart(w http.ResponseWriter, r *http.Request) {
	res, f := Response(r)
	defer f(w)
	res.Msg = "restart sucess"
	if res.Code == 0 && this.Install != nil {
		go func() {
			time.Sleep(time.Second)
			err := this.Install.Restart()
			if err != nil {
				glog.Error("重启失败")
			}
			glog.Error("重启ok")
		}()
	}
}

func (this *CommApi) ApiVersion(w http.ResponseWriter, r *http.Request) {
	res, f := Response(r)
	defer f(w)
	v := map[string]interface{}{
		"frpcVersion": version.Full(),
		"appName":     pkg.AppName,
		"appVersion":  pkg.AppVersion,
		"buildTime":   pkg.BuildTime,
		"gitRevision": pkg.GitRevision,
		"gitBranch":   pkg.GitBranch,
		"goVersion":   pkg.GoVersion,
		"displayName": pkg.DisplayName,
		"description": pkg.Description,
	}
	jsonBytes, err := json.Marshal(v)
	if err != nil {
		res.Error("json marshal err")
		glog.Error(res.Msg)
	}
	res.Raw = jsonBytes
}
