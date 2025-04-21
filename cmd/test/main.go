package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

type ProxyData struct {
	Proxy *httputil.ReverseProxy
	URL   *url.URL
	Host  string
	Port  string
}

func createTargets(host, port string) *ProxyData {
	data := ProxyData{
		Host: host,
		Port: port,
	}
	target, err := url.Parse(fmt.Sprintf("http://%s:%s", host, port))
	if err != nil {
		log.Fatalf("解析目标 URL 时出错: %v", err)
	}
	data.URL = target
	data.Proxy = httputil.NewSingleHostReverseProxy(target)
	return &data
}

func main() {
	target, err := url.Parse("http://192.168.0.2:8888")
	if err != nil {
		log.Fatalf("解析目标 URL 时出错: %v", err)
	}
	proxy := httputil.NewSingleHostReverseProxy(target)
	apis := make(map[string]*ProxyData)
	apis["8888"] = createTargets("192.168.0.2", "8888")
	apis["6500"] = createTargets("192.168.0.3", "6500")

	r := mux.NewRouter()
	r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Pattern:", r.Pattern)
		fmt.Println("RequestURI:", r.RequestURI)
		fmt.Println("Path:", r.URL.Path)
		fmt.Println("Scheme:", r.URL.Scheme)
		fmt.Println("Upgrade:", r.Header.Get("Upgrade"))
		//prefix := utils.GetFirstPathSegment(r.URL.Path)
		//fmt.Println("--->prefix", prefix)
		//if prefix == "" {
		//	w.WriteHeader(404)
		//	return
		//}
		//p := apis[prefix]
		//if p == nil {
		//	w.WriteHeader(404)
		//	return
		//}

		Serve(w, r, target, "fnos", proxy)

	})

	log.Println("http://localhost:8080/fnos")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func Serve(w http.ResponseWriter, r *http.Request, target *url.URL, prefix string, proxy *httputil.ReverseProxy) {
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/"+prefix)
	r.URL.Host = target.Host
	r.URL.Scheme = target.Scheme
	r.Host = target.Host
	fmt.Println("修改后:", r.URL.String())
	proxy.ServeHTTP(w, r)
}
