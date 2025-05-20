package comm

import (
	"fmt"
	"github.com/fatedier/frp/pkg/util/version"
	"github.com/xxl6097/glog/glog"
	"github.com/xxl6097/go-frp-panel/pkg"
	"github.com/xxl6097/go-frp-panel/pkg/model"
	utils2 "github.com/xxl6097/go-frp-panel/pkg/utils"
	"github.com/xxl6097/go-service/gservice/gore"
	"github.com/xxl6097/go-service/gservice/gore/util"
	"github.com/xxl6097/go-service/gservice/utils"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type commapi struct {
	igs  gore.IGService
	obj  any
	pool *sync.Pool // use sync.Pool caching buf to reduce gc ratio
}

func NewCommApi(install gore.IGService, obj any) *commapi {
	return &commapi{
		igs: install,
		obj: obj,
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
		newFilePath = string(body)
		//newFilePath, err = utils.DownLoad()
		//if err != nil {
		//	res.Error(fmt.Sprintf("down load error: %v", err))
		//	glog.Warnf("%s\n", res.Msg)
		//	return
		//}
		glog.Debugf("upgrade by url: %s", newFilePath)
		urls := strings.Split(newFilePath, ",")

		updir := utils.GetUpgradeDir()
		_, _, free, err := util.GetDiskUsage(updir)
		//glog.Printf("Current Working Directory: %s\n", updir)
		//glog.Printf("Total space: %d bytes %v\n", total, utils2.ByteCountIEC(total))
		//glog.Printf("Used space: %d bytes %v\n\n", used, utils2.ByteCountIEC(used))
		//glog.Printf("Free space: %d bytes %v\n\n", free, utils2.ByteCountIEC(free))

		if free < utils2.GetSelfSize()*2 && urls != nil && len(urls) > 0 {
			urls = []string{urls[0]}
			if err := utils2.ClearTmpDir(); err != nil {
				fmt.Println("/tmp清空失败:", err)
			} else {
				fmt.Println("/tmp清空完成")
			}
		}

		newUrl := utils.DownloadFileWithCancelByUrls(urls)
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
		dstFilePath := filepath.Join(utils.GetUpgradeDir(), handler.Filename)
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
	//defer utils.Delete(newFilePath, "更新文件")
	if newFilePath != "" {
		glog.Debugf("开始升级 %s", newFilePath)
		var ch chan error
		go func(ch chan<- error) {
			err := this.igs.Upgrade(ctx, newFilePath)
			glog.Debug("---->升级", err)
			if err == nil {
				res.Ok("升级成功～")
			} else {
				res.Error(fmt.Sprintf("更新失败～%v", err))
			}
			f(w)
			time.Sleep(time.Second)
			ch <- err
			if err != nil {
				res.Error(fmt.Sprintf("更新失败～%v", err))
				return
			}
		}(ch)

		select {
		case <-ctx.Done():
			glog.Error("请求断开", newFilePath)
			break
		case err := <-ch:
			glog.Error("升级成功", err, newFilePath)
			if err != nil {
				res.Error(fmt.Sprintf("更新失败～%v", err))
				return
			} else {
				res.Ok("升级成功～")
			}
		}

		//err := this.igs.Upgrade(ctx, newFilePath)
		//if err != nil {
		//	res.Error(fmt.Sprintf("更新失败～%v", err))
		//	return
		//}
		//res.Ok("升级成功～")
	}
	//下载和接收的最新文件 名称为上传文件的原始名称
	//newBufferBytes, err := ukey.GenConfig(this.obj, false)
	//if err != nil {
	//	res.Error(fmt.Sprintf("gen config err: %v", err))
	//	glog.Error(res.Msg)
	//	return
	//}
	//signFilePath, err := utils.SignAndInstall(newBufferBytes, ukey.UnInitializeBuffer(), newFilePath)
	//glog.Println("签名安装完毕", err, res)
	//if err != nil {
	//	res.Error(err.Error())
	//	glog.Error(res.Msg)
	//} else {
	//	defer utils.Delete(signFilePath, "签名文件")
	//	err = this.igs.Upgrade(signFilePath)
	//	if err != nil {
	//		res.Error(fmt.Sprintf("更新失败～%v", err))
	//		return
	//	}
	//	res.Ok("升级成功～")
	//}
}

func (this *commapi) ApiRestart(w http.ResponseWriter, r *http.Request) {
	res, f := Response(r)
	defer f(w)
	res.Msg = "restart sucess"
	if res.Code == 0 && this.igs != nil {
		go func() {
			time.Sleep(time.Second)
			var err error
			if utils.IsOpenWRT() {
				err = this.igs.RunCmd("restart")
			} else {
				err = this.igs.Restart()
			}
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
	args := utils2.CheckVersionFromGithub()
	if args != nil && len(args) > 0 {
		res.response(1, args[1], args[0])
	} else {
		res.Ok("已经是最新版本～")
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
			err = this.igs.RunCmd("uninstall")
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
	total, used, free := utils2.GetAppSpace()
	hostName, _ := os.Hostname()
	itel, _ := utils2.GetDeviceInfo()
	var ip, netName, netDisplayName, macAddress string
	if itel != nil {
		ip = itel.Ipv4
		netName = itel.Name
		netDisplayName = itel.DisplayName
		macAddress = itel.MacAddress
	}
	res.Sucess("获取成功", map[string]interface{}{
		"frpcVersion":    version.Full(),
		"ipv4":           ip,
		"netDisplayName": netDisplayName,
		"macAddress":     macAddress,
		"netName":        netName,
		"hostName":       hostName,
		"appName":        pkg.AppName,
		"appVersion":     pkg.AppVersion,
		"buildTime":      pkg.BuildTime,
		"gitRevision":    pkg.GitRevision,
		"gitBranch":      pkg.GitBranch,
		"goVersion":      pkg.GoVersion,
		"displayName":    pkg.DisplayName,
		"description":    pkg.Description,
		"osType":         pkg.OsType,
		"arch":           pkg.Arch,
		"totalSize":      total,
		"usedSize":       used,
		"freeSize":       free,
	})
	//glog.Println("操作系统:", runtime.GOOS)     // 如 "linux", "windows"
	//glog.Println("CPU 架构:", runtime.GOARCH) // 如 "amd64", "arm64"
	//glog.Println("CPU 核心数:", runtime.NumCPU())
}
