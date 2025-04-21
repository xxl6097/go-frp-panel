package main

import (
	"bytes"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/xxl6097/go-frp-panel/cmd/proxy/proxy"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func testReserve() {
	// 定义服务映射：前缀 -> 目标地址
	services := map[string]string{
		"/api/service1/": "http://localhost:8384",
		"/api/service2/": "http://localhost:6400",
	}

	// 创建多路复用器
	mux := http.NewServeMux()

	// 为每个服务创建反向代理实例
	for prefix, target := range services {
		targetURL, _ := url.Parse(target)
		proxy := httputil.NewSingleHostReverseProxy(targetURL)

		// 重写请求路径（剥离前缀）
		originalDirector := proxy.Director
		proxy.Director = func(req *http.Request) {
			originalDirector(req)
			req.URL.Path = strings.TrimPrefix(req.URL.Path, prefix)
			// 可选：设置Host头（某些后端服务需要）
			req.Host = targetURL.Host
		}

		// 注册路由处理器
		mux.Handle(prefix, http.StripPrefix(prefix, proxy))
	}

	// 启动服务
	http.ListenAndServe(":8080", mux)
}

func rewriteLocation(location, targetStr, localUrl string) string {
	if strings.HasPrefix(location, targetStr) {
		return strings.Replace(location, targetStr, localUrl, 1) //"http://localhost:8080/frpc"
	}
	return location
}

func reserveFrpc() {
	targetStr := "http://192.168.0.3:5666"
	target, err := url.Parse(targetStr)
	if err != nil {
		log.Fatalf("解析目标 URL 时出错: %v", err)
	}

	p := httputil.NewSingleHostReverseProxy(target)

	localUrl := "http://localhost:8080/frpc"
	// 修改响应，处理重定向和静态资源路径
	p.ModifyResponse = func(resp *http.Response) error {
		if resp.StatusCode == http.StatusMovedPermanently || resp.StatusCode == http.StatusFound {
			location, err := resp.Location()
			if err == nil {
				newLocation := rewriteLocation(location.String(), targetStr, localUrl)
				resp.Header.Set("Location", newLocation)
			}
		}
		contentType := resp.Header.Get("Content-Type")
		if contentType == "text/html" {
			// 处理 HTML 响应中的静态资源路径
			body, err := httputil.DumpResponse(resp, true)
			if err == nil {
				bodyStr := string(body)
				bodyStr = strings.ReplaceAll(bodyStr, targetStr, localUrl)
				//resp.Body = http.NoBody
				//resp.ContentLength = int64(len(bodyStr))
				//resp.Header.Set("Content-Length", string(len(bodyStr)))
				//resp.Body = http.MaxBytesReader(nil, http.NoBody, resp.ContentLength)

				bodySize := len(bodyStr)
				fmt.Println(bodySize)
				resp.ContentLength = int64(bodySize)
				resp.Header.Set("Content-Length", fmt.Sprint(bodySize))
				resp.Body = io.NopCloser(bytes.NewReader([]byte(bodyStr)))
			}
		}
		return nil
	}

	//proxyApi := make(map[string]*proxy.Proxy)
	// 处理请求，pattern 改为 /frpc/
	http.HandleFunc("/frpc/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Pattern)
		fmt.Println(r.RequestURI)
		//var p *proxy.Proxy
		//if v, ok := proxyApi[r.Pattern]; ok {
		//	p = v
		//} else {
		//	p = &proxy.Proxy{}
		//	proxyApi[r.Pattern] = p
		//}
		//p.ServeHTTP(r.Pattern, w, r)
		// 去除 /frpc 前缀
		r.URL.Path = strings.TrimPrefix(r.URL.Path, "/frpc")
		r.URL.Host = target.Host
		r.URL.Scheme = target.Scheme
		r.Host = target.Host
		p.ServeHTTP(w, r)
	})

	log.Println("http://localhost:8080/frpc")
	log.Fatal(http.ListenAndServe(":8080", nil))
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
func Server() {
	targetStr := "http://192.168.0.2:8888"
	target, err := url.Parse(targetStr)
	if err != nil {
		log.Fatalf("解析目标 URL 时出错: %v", err)
	}

	p := httputil.NewSingleHostReverseProxy(target)

	localUrl := "http://localhost:8080/frpc"
	// 修改响应，处理重定向和静态资源路径
	p.ModifyResponse = func(resp *http.Response) error {
		if resp.StatusCode == http.StatusMovedPermanently || resp.StatusCode == http.StatusFound {
			location, err := resp.Location()
			if err == nil {
				newLocation := rewriteLocation(location.String(), targetStr, localUrl)
				resp.Header.Set("Location", newLocation)
			}
		}
		contentType := resp.Header.Get("Content-Type")
		if contentType == "text/html" {
			// 处理 HTML 响应中的静态资源路径
			body, err := httputil.DumpResponse(resp, true)
			if err == nil {
				bodyStr := string(body)
				bodyStr = strings.ReplaceAll(bodyStr, targetStr, localUrl)
				resp.Body = http.NoBody
				bodySize := len(bodyStr)
				//fmt.Println(bodySize)
				resp.ContentLength = int64(bodySize)
				resp.Header.Set("Content-Length", fmt.Sprintf("%d", bodySize))
				//resp.Body = http.MaxBytesReader(nil, http.NoBody, resp.ContentLength)
				resp.Body = io.NopCloser(bytes.NewReader([]byte(bodyStr)))
			}
		}
		return nil
	}
	// 创建主路由器
	r := mux.NewRouter()
	// 创建子路由器
	r1 := r.NewRoute().Name("proxy").Subrouter()
	r1.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		firstPath := getFirstPathSegment(r.URL.Path)
		fmt.Println("url", r.URL.String())
		//host, portStr, e := net.SplitHostPort(firstPath)
		//if e == nil {
		//	fmt.Println(host, portStr, err)
		//} else {
		//	fmt.Println(firstPath, r.Host)
		//}
		// 去除 /frpc 前缀
		r.URL.Path = strings.TrimPrefix(r.URL.Path, fmt.Sprintf("/%s", firstPath))
		r.URL.Host = target.Host
		r.URL.Scheme = target.Scheme
		r.Host = target.Host
		p.ServeHTTP(w, r)
	})
	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	log.Println("代理服务器正在监听 http://localhost:8080/frpc")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("服务器启动失败: %v", err)
	}
}

func proxyServer() {
	r := mux.NewRouter()
	frpcSubrouter := r.NewRoute().Name("proxy").Subrouter()
	frpcSubrouter.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		proxy.GetInstance().Serve(w, r)
	})
	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	log.Println("代理服务器正在监听 http://localhost:8080")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
func main() {
	//reserveFrpc()
	//Server()
	proxyServer()

	//targetStr := "http://192.168.0.3:6400"
	//target, err := url.Parse(targetStr)
	//if err != nil {
	//	log.Fatalf("解析目标 URL 时出错: %v", err)
	//}
	//fmt.Println(target.String())
}
