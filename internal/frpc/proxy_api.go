package frpc

import (
	"fmt"
	"github.com/xxl6097/glog/glog"
	"github.com/xxl6097/go-frp-panel/pkg/comm"
	"github.com/xxl6097/go-frp-panel/pkg/utils"
	utils2 "github.com/xxl6097/go-service/pkg/utils"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func (this *frpc) apiProxyTCPAdd(w http.ResponseWriter, r *http.Request) {
	res, f := comm.Response(r)
	defer f(w)
	cfgFileName := r.URL.Query().Get("name")
	body, err := utils.GetDataByJson[struct {
		Type       string `json:"type"`
		Name       string `json:"name"`
		LocalIP    string `json:"localIP"`
		LocalPort  int    `json:"localPort"`
		RemotePort int    `json:"remotePort"`
	}](r)

	if err != nil {
		res.Err(fmt.Errorf("解析数据失败: %v", err))
		return
	}
	if body == nil {
		res.Error("数据nil")
		return
	}
	binpath, err := os.Executable()
	if err != nil {
		res.Err(fmt.Errorf("get executable path err: %v", err))
		return
	}

	cfgDir := filepath.Dir(binpath)
	if cfgFileName == "" {
		cfgFileName = "config.toml"
	} else {
		cfgDir = filepath.Join(cfgDir, "config")
	}
	cfgFilePath := filepath.Join(cfgDir, cfgFileName)
	if !utils2.FileExists(cfgFilePath) {
		res.Err(fmt.Errorf("客户端不存在: %v", err))
		return
	}
	sb := strings.Builder{}
	sb.WriteString("\r\n[[proxies]]\n")
	sb.WriteString(fmt.Sprintf("name = \"%s\"\n", body.Name))
	sb.WriteString(fmt.Sprintf("type = \"%s\"\n", body.Type))
	sb.WriteString(fmt.Sprintf("localIP = \"%s\"\n", body.LocalIP))
	sb.WriteString(fmt.Sprintf("localPort = %d\n", body.LocalPort))
	sb.WriteString(fmt.Sprintf("remotePort = %d\n", body.RemotePort))
	err = utils.WriteAppend(cfgFilePath, []byte(sb.String()))
	if err != nil {
		res.Err(err)
	}
	res.Ok("sucess")
}

func (this *frpc) apiProxyRemotePorts(w http.ResponseWriter, r *http.Request) {
	res, f := comm.Response(r)
	defer f(w)
	name := r.URL.Query().Get("name")
	//获取当前正在代理的端口
	array := this.getTcpProxyArray(name)
	var ports []comm.Option
	for _, port := range array {
		p := fmt.Sprintf("%d", port)
		ports = append(ports, comm.Option{
			Label: p,
			Value: p,
		})
	}
	res.Any(ports)
}

func (this *frpc) apiProxyPortCheck(w http.ResponseWriter, r *http.Request) {
	res, f := comm.Response(r)
	defer f(w)
	name := r.URL.Query().Get("name")
	port := r.URL.Query().Get("port")
	if port == "" {
		res.Error("port is empty")
		return
	}
	num, err := strconv.Atoi(port)
	if err != nil {
		res.Error("port转换错误")
		return
	}
	var host string
	var cls *frpClient
	if name == "" {
		cls = this.cls
	} else {
		if v, ok := this.svrs[name]; ok {
			cls = v
		}
	}
	if cls == nil {
		res.Err(fmt.Errorf("[%s] cls is nil", name))
		return
	}
	if cls.cfg == nil {
		res.Err(fmt.Errorf("[%s] cls.cfg is nil", name))
		return
	}
	host = cls.cfg.ServerAddr
	if utils.IsPortOpen(host, num, time.Second*3) {
		res.Error("端口被占用")
		return
	}
	res.Ok("端口未被占用")
}

func (this *frpc) apiProxyLocalIps(w http.ResponseWriter, r *http.Request) {
	res, f := comm.Response(r)
	defer f(w)
	array := utils.ScanIP()
	var ports []comm.Option
	for _, ip := range array {
		ports = append(ports, comm.Option{
			Label: ip,
			Value: ip,
		})
	}
	res.Any(ports)
}

func (this *frpc) apiProxyPorts(w http.ResponseWriter, r *http.Request) {
	res, f := comm.Response(r)
	defer f(w)
	localIP := r.URL.Query().Get("localIP")
	if localIP == "" {
		res.Error("localIP is empty")
		return
	}

	arr, _ := utils2.BlockingFunction[[]int](r.Context(), time.Second*20, func() []int {
		start := time.Now()
		data := utils.ScanPorts(localIP, 0, 65535)
		end := time.Now()
		glog.Println("扫描耗时", end.Sub(start))
		glog.Println("端口数量", len(data))
		return data
	})
	if len(arr) == 0 {
		res.Error("localIP is empty")
		return
	}

	var ports []comm.Option
	for _, port := range arr {
		ports = append(ports, comm.Option{
			Label: fmt.Sprintf("%d", port),
			Value: fmt.Sprintf("%d", port),
		})
	}
	res.Any(ports)
}

func (this *frpc) apiProxyGithubApi(w http.ResponseWriter, r *http.Request) {
	res, f := comm.Response(r)
	defer f(w)
	body, err := utils.GetDataByJson[struct {
		ProxyUrl string `json:"proxyUrl"`
	}](r)

	if err != nil {
		res.Err(fmt.Errorf("解析数据失败: %v", err))
		return
	}
	if body == nil {
		res.Error("数据nil")
		return
	}
	err = os.Setenv("GITHUB_API_PROXY", body.ProxyUrl)
	if err != nil {
		res.Err(err)
	} else {
		res.Ok("设置成功～")
	}
	glog.Debug("设置GITHUB_API_PROXY", err)
}
