package frpc

import (
	httppkg "github.com/fatedier/frp/pkg/util/http"
	"github.com/xxl6097/glog/glog"
	"net/http"
)

func (this *frpc) adminHandlers(helper *httppkg.RouterRegisterHelper) {
	subRouter := helper.Router.NewRoute().Name("admin").Subrouter()
	subRouter.Use(helper.AuthMiddleware)
	staticPrefix := "/log/"
	baseDir := glog.GetCrossPlatformDataDir()
	subRouter.PathPrefix(staticPrefix).Handler(http.StripPrefix(staticPrefix, http.FileServer(http.Dir(baseDir))))

	// apis
	subRouter.HandleFunc("/api/version", this.upgrade.ApiVersion).Methods("GET")
	subRouter.HandleFunc("/api/upgrade", this.upgrade.ApiUpdate).Methods("POST")
	subRouter.HandleFunc("/api/upgrade", this.upgrade.ApiUpdate).Methods("PUT")
	subRouter.HandleFunc("/api/restart", this.upgrade.ApiRestart).Methods("GET")
	subRouter.HandleFunc("/api/checkversion", this.upgrade.ApiCheckVersion).Methods("GET")
	subRouter.HandleFunc("/api/uninstall", this.upgrade.ApiUninstall).Methods("GET")

	subRouter.HandleFunc("/api/client/create", this.apiClientCreate).Methods("PUT")
	subRouter.HandleFunc("/api/client/create", this.apiClientCreate).Methods("POST")
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
}
