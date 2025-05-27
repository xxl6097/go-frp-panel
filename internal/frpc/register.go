package frpc

import (
	httppkg "github.com/fatedier/frp/pkg/util/http"
	"github.com/xxl6097/glog/glog"
	"github.com/xxl6097/go-frp-panel/pkg/utils"
	"net/http"
)

var logQueue = utils.NewLogQueue()

func init() {
	glog.Hook(func(bytes []byte) {
		logQueue.AddMessage(string(bytes[2:]))
	})
}

func (this *frpc) adminHandlers(helper *httppkg.RouterRegisterHelper) {
	subRouter := helper.Router.NewRoute().Name("admin").Subrouter()
	subRouter.Use(helper.AuthMiddleware)
	staticPrefix := "/log/"
	baseDir := glog.AppHome()
	subRouter.PathPrefix(staticPrefix).Handler(http.StripPrefix(staticPrefix, http.FileServer(http.Dir(baseDir))))

	subRouter.PathPrefix("/fserver/").Handler(http.StripPrefix("/fserver/", http.FileServer(http.Dir("/"))))
	subRouter.HandleFunc("/api/sse-stream", utils.SseHandler(logQueue))
	subRouter.HandleFunc("/api/files", this.upgrade.ApiFiles).Methods("PUT")

	subRouter.HandleFunc("/api/run", this.upgrade.ApiCMD).Methods("POST")
	subRouter.HandleFunc("/api/clear", this.upgrade.ApiClear).Methods("DELETE")
	// apis
	subRouter.HandleFunc("/api/version", this.upgrade.ApiVersion).Methods("GET")
	subRouter.HandleFunc("/api/upgrade", this.upgrade.ApiUpdate).Methods("POST")
	subRouter.HandleFunc("/api/upgrade", this.upgrade.ApiUpdate).Methods("PUT")
	subRouter.HandleFunc("/api/restart", this.upgrade.ApiRestart).Methods("GET")
	subRouter.HandleFunc("/api/checkversion", this.upgrade.ApiCheckVersion).Methods("GET")
	subRouter.HandleFunc("/api/uninstall", this.upgrade.ApiUninstall).Methods("GET")

	subRouter.HandleFunc("/api/client/create", this.apiClientCreate).Methods("PUT")
	subRouter.HandleFunc("/api/client/create", this.apiClientCreate).Methods("POST")
	subRouter.HandleFunc("/api/client/upload", this.apiClientCreate).Methods("POST")
	subRouter.HandleFunc("/api/client/delete", this.apiClientDelete).Methods("DELETE")
	subRouter.HandleFunc("/api/client/status", this.apiClientStatus).Methods("GET")
	subRouter.HandleFunc("/api/client/list", this.apiClientList).Methods("GET")
	subRouter.HandleFunc("/api/client/config/get", this.apiClientConfigGet).Methods("GET")
	subRouter.HandleFunc("/api/client/config/set", this.apiClientConfigSet).Methods("POST")

	subRouter.HandleFunc("/api/proxy/ports", this.apiProxyPorts).Methods("GET")
	subRouter.HandleFunc("/api/proxy/ips", this.apiProxyLocalIps).Methods("GET")
	subRouter.HandleFunc("/api/proxy/port/check", this.apiProxyPortCheck).Methods("GET")
	subRouter.HandleFunc("/api/proxy/remote/ports", this.apiProxyRemotePorts).Methods("GET")
	subRouter.HandleFunc("/api/proxy/tcp/add", this.apiProxyTCPAdd).Methods("PUT")

	subRouter.HandleFunc("/api/client/config/import", this.apiClientConfigImport).Methods("POST")
	subRouter.HandleFunc("/api/client/config/export", this.apiClientConfigExport).Methods("POST")
}
