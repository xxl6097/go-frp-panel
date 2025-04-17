package proxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

type Proxy struct {
	targetUrl string
	localUrl  string
}

func NewProxy(targetUrl, localUrl string) *httputil.ReverseProxy {
	this := &Proxy{targetUrl, localUrl}
	target, err := url.Parse(targetUrl)
	if err != nil {
		log.Fatalf("解析目标 URL 时出错: %v", err)
	}

	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.ModifyResponse = this.modifyResponse
	return proxy
}

func (this *Proxy) rewriteLocation(location, targetStr string) string {
	if strings.HasPrefix(location, targetStr) {
		return strings.Replace(location, targetStr, this.localUrl, 1)
	}
	return location
}

func (this *Proxy) modifyResponse(resp *http.Response) error {
	if resp.StatusCode == http.StatusMovedPermanently || resp.StatusCode == http.StatusFound {
		location, err := resp.Location()
		if err == nil {
			newLocation := this.rewriteLocation(location.String(), this.targetUrl)
			resp.Header.Set("Location", newLocation)
		}
	}
	contentType := resp.Header.Get("Content-Type")
	if contentType == "text/html" {
		// 处理 HTML 响应中的静态资源路径
		body, err := httputil.DumpResponse(resp, true)
		if err == nil {
			bodyStr := string(body)
			bodyStr = strings.ReplaceAll(bodyStr, this.targetUrl, this.localUrl)
			resp.Body = http.NoBody
			resp.ContentLength = int64(len(bodyStr))
			resp.Header.Set("Content-Length", string(len(bodyStr)))
			resp.Body = http.MaxBytesReader(nil, http.NoBody, resp.ContentLength)
		}
	}
	return nil
}
