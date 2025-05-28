package comm

import (
	"encoding/json"
	"fmt"
	"github.com/xxl6097/glog/glog"
	"github.com/xxl6097/go-frp-panel/pkg"
	iface2 "github.com/xxl6097/go-frp-panel/pkg/comm/iface"
	"github.com/xxl6097/go-frp-panel/pkg/comm/ws"
	"github.com/xxl6097/go-frp-panel/pkg/model"
	utils2 "github.com/xxl6097/go-frp-panel/pkg/utils"
	"github.com/xxl6097/go-service/pkg/github"
	"github.com/xxl6097/go-service/pkg/gs/igs"
	"github.com/xxl6097/go-service/pkg/utils"
	"github.com/xxl6097/go-service/pkg/utils/util"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type commapi struct {
	igs  igs.Service
	pool *sync.Pool // use sync.Pool caching buf to reduce gc ratio
}

func (this *commapi) ApiCMD(w http.ResponseWriter, r *http.Request) {
	res, f := Response(r)
	defer f(w)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		glog.Error("body读取失败", err)
		res.Err(err)
		return
	}
	if body == nil {
		msg := "body is nil"
		glog.Error(msg)
		res.Err(fmt.Errorf(msg))
		return
	}
	var msg iface2.Message[any]
	err = json.Unmarshal(body, &msg)
	if err != nil {
		glog.Error("解析Json对象失败", err)
		res.Err(err)
		return
	}
	switch msg.Action {
	case ws.CLIENT_NETWORLD:
		arr, e := utils2.GetNetworkInterfaces()
		if e != nil {
			res.Err(e)
		} else {
			res.Any(arr)
		}
		break
	case ws.CMD:
		data, ok := msg.Data.(map[string]interface{})
		if ok {
			glog.Infof("data %+v", data)
			d := data["data"]
			if d == nil {
				glog.Errorf("data is nil %+v", msg.Data)
				break
			}
			v, okk := d.(string)
			if !okk {
				glog.Infof("string err %+v", d)
				break
			}
			arrData := strings.Split(v, " ")
			var cmd *exec.Cmd
			if len(arrData) >= 2 {
				cmd = exec.Command(arrData[0], arrData[1:]...)
			} else {
				cmd = exec.Command(arrData[0])
			}
			output, err := cmd.CombinedOutput()
			if err != nil {
				res.Err(err)
				return
			}
			res.Any(string(output))
		} else {
			res.Err(fmt.Errorf("cmd err %+v", msg.Data))
		}
		break
	}
}

func NewCommApi(install igs.Service) *commapi {
	return &commapi{
		igs: install,
		pool: &sync.Pool{
			New: func() interface{} { return make([]byte, 32*1024) },
		},
	}
}

func (this *commapi) GetBuffer() *sync.Pool {
	return this.pool
}

func (this *commapi) ApiFiles(w http.ResponseWriter, r *http.Request) {
	res, f := Response(r)
	defer f(w)
	params, err := utils2.GetDataByJson[struct {
		Path string `json:"path"`
	}](r)
	if err != nil {
		res.Error(fmt.Errorf("read param err: %v", err).Error())
		glog.Error(res.Msg)
		return
	}
	if params == nil {
		res.Error("params is empty")
		glog.Error(res.Msg)
		return
	}
	path := params.Path
	isFile := strings.HasSuffix(path, "/")

	if !isFile {
		w.Header().Set("File-Type", "text")
		http.ServeFile(w, r, path) //r.URL.Path
	} else {
		dirs, err := os.ReadDir(path)
		if err != nil {
			return
		}

		var files []model.TreeData
		for _, dir := range dirs {
			f := model.TreeData{
				Id:    dir.Name(),
				Label: dir.Name(),
			}
			if dir.IsDir() {
				f.Label = dir.Name() + "/"
			}
			files = append(files, f)
		}
		res.Any(files)
	}

}

func (this *commapi) ApiUpdate1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	select {
	case <-time.After(20 * time.Second):
		fmt.Println("Operation completed")
		w.Write([]byte("Operation completed"))
	case <-ctx.Done():
		// 客户端断开或超时
		//if ctx.Err() == context.Canceled {
		//}
		fmt.Println("Client disconnected", ctx.Err())
	}
}

func (this *commapi) ApiUpdate(w http.ResponseWriter, r *http.Request) {
	res, f := Response(r)
	defer f(w)
	ctx := r.Context()
	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()
	updir := glog.AppHome()
	_, _, free, _ := util.GetDiskUsage(updir)
	if free < utils2.GetSelfSize()*2 {
		if err := utils2.ClearTmpDir(); err != nil {
			fmt.Println("/tmp清空失败:", err)
		} else {
			fmt.Println("/tmp清空完成")
		}
	}

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
			res.Response(400, "升级URL空的哦～")
			glog.Warnf("%s", res.Msg)
			return
		}
		binUrl := string(body)
		glog.Debugf("upgrade by url: %s", newFilePath)
		newUrl := utils.DownloadFileWithCancelByUrls(github.Api().GetProxyUrls(binUrl))
		newFilePath = newUrl
		break
	case "POST", "post":
		// 获取上传的文件
		file, handler, err := r.FormFile("file")
		if err != nil {
			res.Error("body no file")
			return
		}
		defer file.Close()
		dstFilePath := filepath.Join(glog.AppHome("upgrade"), handler.Filename)
		//dstFilePath 名称为上传文件的原始名称
		dst, err := os.Create(dstFilePath)
		if err != nil {
			res.Error(fmt.Sprintf("create file %s error: %v", handler.Filename, err))
			return
		}
		buf := this.pool.Get().([]byte)
		defer this.pool.Put(buf)
		_, err = io.CopyBuffer(dst, file, buf)
		dst.Close()
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
		glog.Debugf("开始升级 %s", newFilePath)
		err := this.igs.Upgrade(ctx, newFilePath)
		glog.Debug("---->升级", err)
		if err == nil {
			res.Ok("升级成功～")
		} else {
			res.Error(fmt.Sprintf("更新失败～%v", err))
		}

	}
}

func (this *commapi) ApiRestart(w http.ResponseWriter, r *http.Request) {
	res, f := Response(r)
	defer f(w)
	res.Msg = "restart sucess"
	if res.Code == 0 && this.igs != nil {
		go func() {
			time.Sleep(time.Second)
			var err error
			err = this.igs.Restart()
			//if utils.IsOpenWRT() {
			//	err = this.igs.RunCmd("restart")
			//} else {
			//	err = this.igs.Restart()
			//}
			if err != nil {
				glog.Error("重启失败")
			}
			glog.Error("重启ok")
		}()
	}
}

func (this *commapi) ApiCheckVersion(w http.ResponseWriter, r *http.Request) {
	res, f := Response(r)
	defer f(w)
	data, err := github.Api().DefaultRequest().CheckUpgrade(pkg.BinName, nil).Result()
	if err != nil {
		res.Err(err)
	} else {
		glog.Debug("version:", data)
		res.Any(data)
	}
	//args, err := utils2.CheckVersionFromGithub()
	//if err != nil {
	//	res.Err(err)
	//	return
	//}
	//if args != nil && len(args) > 0 {
	//	res.response(1, args[1], args[0])
	//} else {
	//	res.Ok("已经是最新版本～")
	//}
}

// /api/shutdown
func (this *commapi) ApiClear(w http.ResponseWriter, r *http.Request) {
	res, f := Response(r)
	defer f(w)
	glog.Infof("Http request: [%s]", r.URL.Path)
	binPath, err := os.Executable()
	if err != nil {
		res.Error(fmt.Sprintf("获取当前可执行文件路径出错: %v\n", err))
		glog.Error(res.Msg)
		return
	}
	binDir := filepath.Dir(binPath)
	clientsDir := filepath.Join(binDir, "clients")
	err = utils.DeleteAllDirector(clientsDir)
	if this.igs != nil {
		err = this.igs.ClearCache()
	}
	if err != nil {
		res.Err(err)
	} else {
		res.Msg = "删除成功"
	}
}
func (this *commapi) ApiUninstall(w http.ResponseWriter, r *http.Request) {
	res, f := Response(r)
	defer f(w)
	res.Msg = "uninstall sucess"
	if res.Code == 0 && this.igs != nil {
		go func() {
			time.Sleep(time.Second)
			var err error
			//if utils.IsOpenWRT() {
			//	err = this.igs.RunCmd("uninstall")
			//} else {
			//	err = this.igs.Uninstall()
			//}
			//err = this.igs.RunCmd("uninstall")
			err = this.igs.UnInstall()
			if err != nil {
				glog.Error("uninstall 失败", err)
			} else {
				glog.Error("uninstall ok")
			}
		}()
	}
}
func (this *commapi) ApiVersion(w http.ResponseWriter, r *http.Request) {
	res, f := Response(r)
	defer f(w)
	res.Sucess("获取成功", utils2.GetVersion())
	//glog.Println("操作系统:", runtime.GOOS)     // 如 "linux", "windows"
	//glog.Println("CPU 架构:", runtime.GOARCH) // 如 "amd64", "arm64"
	//glog.Println("CPU 核心数:", runtime.NumCPU())
}
