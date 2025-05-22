package frps

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	v1 "github.com/fatedier/frp/pkg/config/v1"
	"github.com/xxl6097/glog/glog"
	model2 "github.com/xxl6097/go-frp-panel/internal/com/model"
	"github.com/xxl6097/go-frp-panel/pkg/frp"
	"github.com/xxl6097/go-frp-panel/pkg/utils"
	"github.com/xxl6097/go-service/gservice/ukey"
	utils2 "github.com/xxl6097/go-service/gservice/utils"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func (this *frps) createConfigData(body *model2.ConfigBodyData) error {
	var e error
	if body.ClientConfig == nil || body.ClientConfig.ServerAddr == "" {
		e = fmt.Errorf("body.ClientConfig.ServerAddr is nil")
		return e
	}
	if body.UserConfig == nil || body.UserConfig.User == "" {
		e = fmt.Errorf("body.UserConfig.User is nil")
		return e
	}

	if body.ClientConfig.WebServer.Port <= 0 ||
		body.ClientConfig.WebServer.Addr == "" ||
		body.ClientConfig.WebServer.User == "" ||
		body.ClientConfig.WebServer.Password == "" {
		body.ClientConfig.WebServer = v1.WebServerConfig{}
	}

	if body.ClientConfig.Proxies == nil {
		body.ClientConfig.Proxies = nil
	}

	var proxies []v1.TypedProxyConfig
	for _, v := range body.ClientConfig.Proxies {
		if frp.HasProxyes(&v) {
			proxies = append(proxies, v)
		}
	}
	if proxies == nil || len(proxies) == 0 {
		body.ClientConfig.Proxies = nil
	} else {
		body.ClientConfig.Proxies = proxies
	}

	if body.UserConfig != nil {
		config := model2.FrpcBuffer{
			User:       *body.UserConfig,
			ServerAddr: body.ClientConfig.ServerAddr,
			ServerPort: body.ClientConfig.ServerPort,
			AdminUser:  body.ClientConfig.WebServer.User,
			AdminPass:  body.ClientConfig.WebServer.Password,
		}
		if keycode, e1 := frp.EncodeSecret(&config); e1 == nil {
			if body.ClientConfig.Metadatas == nil {
				body.ClientConfig.Metadatas = make(map[string]string)
			}
			body.ClientConfig.Metadatas["secret"] = keycode
		}
	}
	return nil
}

func (this *frps) getBody(r *http.Request) (*model2.ConfigBodyData, error) {
	body, e := utils.GetDataByJson[model2.ConfigBodyData](r)
	if e != nil {
		e = fmt.Errorf("ConfigBodyData convert error %v", e)
		return nil, e
	}
	if body == nil {
		e = fmt.Errorf("body is nil")
		return nil, e
	}
	e = this.createConfigData(body)
	if e != nil {
		return nil, e
	}
	return body, nil
}

func (this *frps) getData(r *http.Request) (*model2.ConfigBodyData, error) {
	var body *model2.ConfigBodyData = nil
	jsonString := r.FormValue("data")
	e := json.Unmarshal([]byte(jsonString), body)
	if e != nil {
		return nil, e
	}
	glog.Infof("data:%+v", body)
	e = this.createConfigData(body)
	if e != nil {
		return nil, e
	}
	return body, nil
}

func (this *frps) apiCreateFrpcToml(w http.ResponseWriter, r *http.Request) {
	var e error
	defer func() {
		if e != nil {
			glog.Error(e)
			http.Error(w, e.Error(), http.StatusInternalServerError)
		}
	}()
	body, err := this.getBody(r)
	if err != nil {
		e = err
		return
	}
	fileName := fmt.Sprintf("%s.%s.frpc.toml", body.ClientConfig.ServerAddr, body.UserConfig.User)
	buffer := utils.ObjectToTomlText(body.ClientConfig)
	w.Header().Add("Content-Transfer-Encoding", "binary")
	w.Header().Add("Content-Type", "application/octet-stream")
	w.Header().Add(`Content-Length`, strconv.Itoa(len(buffer)))
	w.Header().Add(`Content-Disposition`, fmt.Sprintf("attachment; filename=\"%s\"", fileName))
	_, _ = w.Write(buffer)
}

func (this *frps) apiCreateFrpcByUpload(w http.ResponseWriter, r *http.Request) {
	var e error
	defer func() {
		if e != nil {
			glog.Error(e)
			http.Error(w, e.Error(), http.StatusInternalServerError)
		}
	}()
	body, err := this.getData(r)
	if err != nil {
		e = err
		return
	}
	glog.Infof("body:%+v", body)
	err = r.ParseMultipartForm(32 << 20)
	if err != nil {
		e = fmt.Errorf("ParseMultipartForm error %v", err)
		return
	}
	// 获取上传的文件
	file, handler, err := r.FormFile("file")
	if err != nil {
		e = fmt.Errorf("body no file")
		return
	}
	defer file.Close()

	glog.Info(handler.Filename)

	binPath := filepath.Join(glog.GetCrossPlatformDataDir("temp"), handler.Filename)
	dst, err := os.Create(binPath)
	if err != nil {
		e = fmt.Errorf("create file %s error: %v", handler.Filename, err)
		return
	}
	defer utils2.DeleteAll(binPath, "upload gen file")
	buf := this.upgrade.GetBuffer().Get().([]byte)
	defer this.upgrade.GetBuffer().Put(buf)
	_, err = io.CopyBuffer(dst, file, buf)
	_ = dst.Close()
	if err != nil {
		e = fmt.Errorf("io.CopyBuffer error: %v", err)
		return
	}
	glog.Info("上传成功", binPath)
	e = this.serveFile(binPath, body, w, r)
}

func (this *frps) apiCreateFrpcByUrl(w http.ResponseWriter, r *http.Request) {
	var e error
	defer func() {
		if e != nil {
			glog.Error(e)
			http.Error(w, e.Error(), http.StatusInternalServerError)
		}
	}()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	body, err := this.getBody(r)
	if err != nil {
		e = err
		return
	}
	glog.Infof("body:%+v", body)
	if utils2.IsURL(body.BinAddress) {
		if this.githubProxys != nil {
			var urls []string
			for _, proxy := range this.githubProxys {
				newUrl := fmt.Sprintf("%s%s", proxy, body.BinAddress)
				urls = append(urls, newUrl)
			}
			dstPath := utils2.DownloadFileWithCancelByUrls(urls)
			body.BinAddress = dstPath
		} else {
			dstPath, err1 := utils2.DownloadFileWithCancel(ctx, body.BinAddress)
			if err1 != nil {
				e = fmt.Errorf("下载文件失败～%v", err1)
				return
			}
			body.BinAddress = dstPath
		}
	}
	binPath := body.BinAddress
	if binPath == "" {
		e = fmt.Errorf("bin文件路径空")
		return
	}
	e = this.serveFile(binPath, body, w, r)
}

func (this *frps) serveFile(filePath string, body *model2.ConfigBodyData, w http.ResponseWriter, r *http.Request) error {
	glog.Infof("filePath: %s %+v\n", filePath, body)
	tpl, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("打开文件失败：%v", err)
	}
	defer func() {
		_ = tpl.Close()
	}()

	fileName := filepath.Base(filePath)
	w.Header().Add("Content-Transfer-Encoding", "binary")
	w.Header().Add("Content-Type", "application/octet-stream")
	if stat, err1 := tpl.Stat(); err1 == nil {
		w.Header().Add(`Content-Length`, strconv.FormatInt(stat.Size(), 10))
	}
	w.Header().Add(`Content-Disposition`, fmt.Sprintf("attachment; filename=\"%s\"", fileName))

	cfgBuffer := bytes.Repeat([]byte{byte(ukey.B)}, len(ukey.GetBuffer()))
	glog.Infof("ClientConfig: %+v", body.ClientConfig)
	cfgNewBytes, err := ukey.GenConfig(body.ClientConfig, false)
	if err != nil {
		return fmt.Errorf("文件签名失败：%v", err)

	}
	dstFile := filepath.Join(glog.GetCrossPlatformDataDir("temp", utils2.SecureRandomID()), fileName)
	outFile, err := os.Create(dstFile)
	if err != nil {
		_ = utils2.DeleteAll(dstFile, "创建失败，删除")
		return fmt.Errorf("创建失败：%v", err)
	}
	defer func() {
		_ = outFile.Close()
		_ = utils2.DeleteAll(dstFile, "gen file")
	}()
	prevBuffer := make([]byte, 0)
	for {
		thisBuffer := make([]byte, 1024)
		n, err1 := tpl.Read(thisBuffer)
		thisBuffer = thisBuffer[:n]
		tempBuffer := append(prevBuffer, thisBuffer...)
		bufIndex := bytes.Index(tempBuffer, cfgBuffer)
		if bufIndex > -1 {
			tempBuffer = bytes.Replace(tempBuffer, cfgBuffer, cfgNewBytes, -1)
		}
		//w.Write(tempBuffer[:len(prevBuffer)])
		_, _ = outFile.Write(tempBuffer[:len(prevBuffer)])
		prevBuffer = tempBuffer[len(prevBuffer):]
		if err1 != nil {
			break
		}
	}
	if len(prevBuffer) > 0 {
		//w.Write(prevBuffer)
		_, _ = outFile.Write(prevBuffer)
		prevBuffer = nil
	}
	http.ServeFile(w, r, dstFile)
	return nil
}
