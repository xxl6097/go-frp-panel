package proxy

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"strings"
	"sync"
)

type Proxy struct {
	Url   *url.URL
	Proxy *httputil.ReverseProxy
}

// singleton 单例对象类型
type singleton struct {
	proxies       map[int]*Proxy
	proxiesRemote map[string]*Proxy
}

var (
	instance *singleton
	once     sync.Once
)

// GetInstance 返回单例实例
func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{proxies: make(map[int]*Proxy), proxiesRemote: make(map[string]*Proxy)} // 初始化逻辑
		fmt.Println("Singleton instance created")
	})
	return instance
}

//type Proxy struct {
//	targetUrl string
//	localUrl  string
//	proxy     *httputil.ReverseProxy
//	target    *url.URL
//}

//func NewProxy1(targetUrl, localUrl string) *Proxy {
//	target, err := url.Parse(targetUrl)
//	if err != nil {
//		log.Fatalf("解析目标 URL 时出错: %v", err)
//	}
//	proxy := httputil.NewSingleHostReverseProxy(target)
//	this := &Proxy{targetUrl, localUrl, proxy, target}
//	proxy.ModifyResponse = this.modifyResponse
//	return this
//}
//func NewProxy(localUrl string) *Proxy {
//	target, err := url.Parse(targetUrl)
//	if err != nil {
//		log.Fatalf("解析目标 URL 时出错: %v", err)
//	}
//	proxy := httputil.NewSingleHostReverseProxy(target)
//	this := &Proxy{targetUrl, localUrl, proxy, target}
//	proxy.ModifyResponse = this.modifyResponse
//	return this
//}

//	func (instance *singleton) AddTarget(targetUrl string) error {
//		targetStr := "http://192.168.0.3:6400"
//		target, err := url.Parse(targetStr)
//		if err != nil {
//			return err
//		}
//		return err
//	}

//func (this *singleton) ServeHTTP(pattern string, w http.ResponseWriter, r *http.Request) {
//	r.URL.Path = strings.TrimPrefix(r.URL.Path, pattern)
//	r.URL.Host = this.target.Host
//	r.URL.Scheme = this.target.Scheme
//	r.Host = this.target.Host
//	this.proxy.ServeHTTP(w, r)
//}

func (p *singleton) modifyResponse(targetUrl, localUrl string, o *httputil.ReverseProxy) {
	o.ModifyResponse = func(resp *http.Response) error {
		if resp.StatusCode == http.StatusMovedPermanently || resp.StatusCode == http.StatusFound {
			location, err := resp.Location()
			if err == nil {
				newLocation := rewriteLocation(location.String(), targetUrl, localUrl)
				resp.Header.Set("Location", newLocation)
			}
		}
		//contentType := resp.Header.Get("Content-Type")
		//if contentType == "text/html" {
		//	// 处理 HTML 响应中的静态资源路径
		//	body, err := httputil.DumpResponse(resp, true)
		//	if err == nil {
		//		bodyStr := string(body)
		//		bodyStr = strings.ReplaceAll(bodyStr, targetUrl, localUrl)
		//		//resp.Body = http.NoBody
		//		//resp.ContentLength = int64(len(bodyStr))
		//		//resp.Header.Set("Content-Length", string(len(bodyStr)))
		//		//resp.Body = http.MaxBytesReader(nil, http.NoBody, resp.ContentLength)
		//
		//		bodySize := len(bodyStr)
		//		resp.Body = io.NopCloser(bytes.NewReader([]byte(bodyStr)))
		//		resp.ContentLength = int64(bodySize)
		//		resp.Header.Set("Content-Length", fmt.Sprintf("%d", bodySize))
		//	}
		//}
		return nil
	}
}
func (this *singleton) serveLocal(firstPath string, w http.ResponseWriter, r *http.Request) {
	port, err := strconv.Atoi(firstPath)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	p := this.proxies[port]
	if p == nil {
		target, e := url.Parse(fmt.Sprintf("http://localhost:%d", port))
		if e != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		p = &Proxy{Url: target, Proxy: httputil.NewSingleHostReverseProxy(target)}
		this.modifyResponse(target.String(), fmt.Sprintf("http://%s/%s", r.Host, firstPath), p.Proxy)
		this.proxies[port] = p
	}

	r.URL.Path = strings.TrimPrefix(r.URL.Path, fmt.Sprintf("/%s", firstPath))
	r.URL.Host = p.Url.Host
	r.URL.Scheme = p.Url.Scheme
	r.Host = p.Url.Host
	p.Proxy.ServeHTTP(w, r)
}

func (this *singleton) serveRemote(hosts, ports string, w http.ResponseWriter, r *http.Request) {
	hostname := fmt.Sprintf("%s:%s", hosts, ports)
	targetUrl := fmt.Sprintf("http://%s", hostname)
	p := this.proxiesRemote[targetUrl]
	if p == nil {
		target, e := url.Parse(targetUrl)
		if e != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		p = &Proxy{Url: target, Proxy: httputil.NewSingleHostReverseProxy(target)}
		this.modifyResponse(target.String(), fmt.Sprintf("http://%s", r.Host), p.Proxy)
		this.modifyResponse(target.String(), fmt.Sprintf("http://%s/%s", r.Host, hostname), p.Proxy)
		this.proxiesRemote[targetUrl] = p
	}

	r.URL.Path = strings.TrimPrefix(r.URL.Path, fmt.Sprintf("/%s", hostname))
	r.URL.Host = p.Url.Host
	r.URL.Scheme = p.Url.Scheme
	r.Host = p.Url.Host
	p.Proxy.ServeHTTP(w, r)
}

func (this *singleton) Serve(w http.ResponseWriter, r *http.Request) {
	firstPath := getFirstPathSegment(r.URL.Path)
	host, portStr, e := net.SplitHostPort(firstPath)
	if e == nil {
		fmt.Println(host, portStr, e)
		this.serveRemote(host, portStr, w, r)
	} else {
		fmt.Println(firstPath, r.Host)
		this.serveLocal(firstPath, w, r)
	}
}

func rewriteLocation(location, targetStr, localUrl string) string {
	if strings.HasPrefix(location, targetStr) {
		return strings.Replace(location, targetStr, localUrl, 1)
	}
	return location
}

// 获取路径第一个有效段（优化实现）
func getFirstPathSegment(path string) string {
	// 清理路径并分割
	trimmed := strings.Trim(path, "/")
	if len(trimmed) == 0 {
		return ""
	}

	// 分割并返回第一个非空段
	if idx := strings.IndexByte(trimmed, '/'); idx >= 0 {
		return trimmed[:idx]
	}
	return trimmed
}
