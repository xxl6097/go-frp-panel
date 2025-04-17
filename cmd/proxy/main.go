package main

import (
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

func rewriteLocation(location, targetStr string) string {
	if strings.HasPrefix(location, targetStr) {
		return strings.Replace(location, targetStr, "http://localhost:8080/frpc", 1)
	}
	return location
}

func reserveFrpc() {
	targetStr := "http://127.0.0.1:6400"
	target, err := url.Parse(targetStr)
	if err != nil {
		log.Fatalf("解析目标 URL 时出错: %v", err)
	}

	proxy := httputil.NewSingleHostReverseProxy(target)

	// 修改响应，处理重定向和静态资源路径
	proxy.ModifyResponse = func(resp *http.Response) error {
		if resp.StatusCode == http.StatusMovedPermanently || resp.StatusCode == http.StatusFound {
			location, err := resp.Location()
			if err == nil {
				newLocation := rewriteLocation(location.String(), targetStr)
				resp.Header.Set("Location", newLocation)
			}
		}
		contentType := resp.Header.Get("Content-Type")
		if contentType == "text/html" {
			// 处理 HTML 响应中的静态资源路径
			body, err := httputil.DumpResponse(resp, true)
			if err == nil {
				bodyStr := string(body)
				bodyStr = strings.ReplaceAll(bodyStr, targetStr, "http://localhost:8080/frpc")
				resp.Body = http.NoBody
				resp.ContentLength = int64(len(bodyStr))
				resp.Header.Set("Content-Length", string(len(bodyStr)))
				resp.Body = http.MaxBytesReader(nil, http.NoBody, resp.ContentLength)
			}
		}
		return nil
	}

	// 处理请求，pattern 改为 /frpc/
	http.HandleFunc("/frpc/", func(w http.ResponseWriter, r *http.Request) {
		// 去除 /frpc 前缀
		r.URL.Path = strings.TrimPrefix(r.URL.Path, "/frpc")
		r.URL.Host = target.Host
		r.URL.Scheme = target.Scheme
		r.Host = target.Host
		proxy.ServeHTTP(w, r)
	})

	log.Println("代理服务器正在监听 :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	reserveFrpc()
}
