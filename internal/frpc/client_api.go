package frpc

import (
	"fmt"
	"github.com/fatedier/frp/pkg/config"
	"github.com/xxl6097/glog/glog"
	"github.com/xxl6097/go-frp-panel/pkg/comm"
	"github.com/xxl6097/go-frp-panel/pkg/frp"
	"github.com/xxl6097/go-frp-panel/pkg/utils"
	utils2 "github.com/xxl6097/go-service/gservice/utils"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func (this *frpc) apiUploadCreate(w http.ResponseWriter, r *http.Request) {
	res, f := comm.Response(r)
	defer f(w)
	res.Ok("")
}

func (this *frpc) apiClientCreate(w http.ResponseWriter, r *http.Request) {
	res, f := comm.Response(r)
	defer f(w)
	var newFilePath string
	cfgDir, err := frp.GetFrpcTomlDir()
	if err != nil {
		res.Err(fmt.Errorf("check config dir err: %v", err))
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
		if body.Name == "" {
			res.Error("文件名空")
			glog.Error(res.Msg)
			return
		}
		if body.Toml == "" {
			res.Error("toml配置空")
			glog.Error(res.Msg)
			return
		}

		if filepath.Ext(body.Name) != ".toml" {
			res.Error("文件必须是toml后缀～")
			glog.Error(res.Msg)
			return
		}
		cfgFilePath := filepath.Join(cfgDir, body.Name)
		if utils2.FileExists(cfgFilePath) {
			res.Err(fmt.Errorf("客户端已经存在"))
			return
		}
		//err = utils.WriteToml(cfgFilePath, []byte(body.Toml))
		err = frp.WriteFrpToml(cfgFilePath, body.Toml)
		if err != nil {
			res.Err(fmt.Errorf("write http body err: %v", err))
			utils.Delete(cfgFilePath)
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
		if filepath.Ext(handler.Filename) != ".toml" {
			res.Error("文件必须是toml后缀～")
			glog.Error(res.Msg)
			return
		}
		dstFilePath := filepath.Join(cfgDir, handler.Filename)
		if utils2.FileExists(dstFilePath) {
			res.Err(fmt.Errorf("客户端已经存在"))
			return
		}
		//dstFilePath 名称为上传文件的原始名称
		dst, err := os.Create(dstFilePath)
		if err != nil {
			res.Error(fmt.Sprintf("create file %s error: %v", handler.Filename, err))
			utils.Delete(dstFilePath)
			return
		}
		buf := this.upgrade.GetBuffer().Get().([]byte)
		defer this.upgrade.GetBuffer().Put(buf)
		_, err = io.CopyBuffer(dst, file, buf)
		dst.Close()
		//err = utils.SaveFile(file, handler.Size, dstFilePath)
		if err != nil {
			res.Error(err.Error())
			utils.Delete(dstFilePath)
			return
		}
		newFilePath = dstFilePath
		break
	default:
		res.Error("位置请求方法")
	}

	if newFilePath != "" {
		_, _, _, _, err = config.LoadClientConfig(newFilePath, true)
		if err != nil {
			res.Err(fmt.Errorf("文件不合法: %v", err))
			utils.Delete(newFilePath)
			return
		}

		//err = retry.Do(func() error {
		//	e := this.runClient(newFilePath)
		//	if e != nil {
		//		glog.Errorf("创建frpc客户端失败: %s %v\n", newFilePath, e)
		//	}
		//	return e
		//}, retry.Delay(time.Second*5), retry.Attempts(10))
		err = this.runClient(newFilePath)
		glog.Error(err)
		if err != nil {
			res.Err(err)
			utils.Delete(newFilePath)
			return
		}
		res.Ok("创建成功～")
	}

}

func (this *frpc) apiClientDelete(w http.ResponseWriter, r *http.Request) {
	res, f := comm.Response(r)
	defer f(w)
	cfgName := r.URL.Query().Get("name")
	if cfgName == "" {
		res.Error("cfg file path is empty")
		return
	}

	cfgDir, err := frp.GetFrpcTomlDir()
	if err != nil {
		res.Err(fmt.Errorf("get executable path err: %v", err))
		return
	}

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

	cfgDir, err := frp.GetFrpcTomlDir()
	if err != nil {
		res.Err(fmt.Errorf("get executable path err: %v", err))
		return
	}

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

	cfgDir, err := frp.GetFrpcTomlDir()
	if err != nil {
		res.Err(fmt.Errorf("get executable path err: %v", err))
		return
	}

	if utils.IsDirectoryExist(cfgDir) {
		files, err := os.ReadDir(cfgDir)
		if err != nil {
			res.Err(fmt.Errorf("read config dir err: %v", err))
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
	} else {
		res.Err(fmt.Errorf("配置目录不存在：%v", cfgDir))
	}

}

func (this *frpc) getClientMainConfig() ([]byte, error) {
	//body, err := utils.ReadToml(this.cls.configFilePath)
	//if err != nil {
	//	return nil, fmt.Errorf("write http body err: %v", err)
	//}
	//return body, nil
	return frp.ReadFrpToml(frp.GetFrpcMainTomlFileName())
}

func (this *frpc) getClientChildConfig(cfgName string) ([]byte, error) {
	return frp.ReadFrpToml(cfgName)
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
	body, err := this.getClientChildConfig(cfgName)
	if err != nil {
		res.Err(fmt.Errorf("get executable path err: %v", err))
		return
	}
	//res.Raw = body
	res.Any(string(body))
}

func (this *frpc) upgradeMainTomlContent(content string) error {
	//err := utils.WriteToml(this.cls.configFilePath, []byte(content))
	err := frp.WriteFrpToml(this.cls.configFilePath, content)
	if err != nil {
		return fmt.Errorf("write http body err: %v", err)
	}
	err = this.upgradeMainConfig()
	if err != nil {
		return fmt.Errorf("run client err: %v", err)
	}
	return nil
}

func (this *frpc) upgradeTomlContent(name, content string) error {
	//err := utils.WriteToml(this.cls.configFilePath, []byte(content))
	cfgPath, err := frp.GetFrpcTomlPath(name)
	if err != nil {
		return fmt.Errorf("get executable path err: %v", err)
	}
	err = frp.WriteFrpToml(cfgPath, content)
	if err != nil {
		return fmt.Errorf("write http body err: %v", err)
	}
	err = this.updateClient(cfgPath)
	if err != nil {
		return fmt.Errorf("run client err: %v", err)
	}
	return nil
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

	cfgDir, err := frp.GetFrpcTomlDir()
	if err != nil {
		res.Err(fmt.Errorf("get executable path err: %v", err))
		return
	}
	cfgFilePath := filepath.Join(cfgDir, body.Name)
	if !utils2.FileExists(cfgFilePath) {
		res.Err(fmt.Errorf("客户端不存在: %v", err))
		return
	}
	//err = utils.WriteToml(cfgFilePath, []byte(body.Toml))
	err = frp.WriteFrpToml(cfgFilePath, body.Toml)
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

func (this *frpc) apiClientConfigExport(w http.ResponseWriter, r *http.Request) {
	res, f := comm.Response(r)
	defer f(w)
	cfgDir, err := frp.GetFrpcTomlDir()
	if err != nil {
		res.Err(err)
		return
	}
	var zipFilePath string
	fileName := fmt.Sprintf("config_%s.zip", utils.GetFileNameByTime())
	tempDir := glog.GetCrossPlatformDataDir("config")
	zipFilePath = filepath.Join(tempDir, fileName)
	err = utils.Zip(cfgDir, zipFilePath)
	if err != nil {
		res.Err(err)
		return
	}

	defer utils2.Delete(zipFilePath, "frpc配置")
	tpl, err := os.Open(zipFilePath)
	if err != nil {
		res.Err(err)
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
func (this *frpc) apiClientConfigImport(w http.ResponseWriter, r *http.Request) {
	res := &comm.GeneralResponse{Code: 0}
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

	cfgDir, err := frp.GetFrpcTomlDir()
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

		if err != nil {
			res.Error(err.Error())
			return
		}
		err = utils.UnzipToRoot(dstFilePath, cfgDir, true)
		if err == nil {
			utils.Delete(dstFilePath, "用户文件")
			glog.Info("解压成功", cfgDir)
		}
		break
	default:
		res.Error("file type not support")
	}
}
