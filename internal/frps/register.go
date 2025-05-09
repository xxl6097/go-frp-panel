package frps

import (
	httppkg "github.com/fatedier/frp/pkg/util/http"
	"github.com/xxl6097/glog/glog"
	"github.com/xxl6097/go-frp-panel/pkg"
	"net/http"
)

func (this *frps) adminHandlers(helper *httppkg.RouterRegisterHelper) {
	subRouter := helper.Router.NewRoute().Name("admin").Subrouter()
	subRouter.Use(helper.AuthMiddleware)
	staticPrefix := "/log/"
	baseDir := glog.GetCrossPlatformDataDir()
	subRouter.PathPrefix(staticPrefix).Handler(http.StripPrefix(staticPrefix, http.FileServer(http.Dir(baseDir))))

	subRouter.PathPrefix("/fserver/").Handler(http.StripPrefix("/fserver/", http.FileServer(http.Dir("/"))))
	subRouter.HandleFunc("/api/sse-stream", pkg.SseHandler(logQueue))
	subRouter.HandleFunc("/api/files", this.upgrade.ApiFiles).Methods("PUT")

	// apis
	//subRouter.HandleFunc("/api/panelinfo", this.apiPanelinfo).Methods("GET")
	subRouter.HandleFunc("/api/restart", this.upgrade.ApiRestart).Methods("GET")
	subRouter.HandleFunc("/api/checkversion", this.upgrade.ApiCheckVersion).Methods("GET")
	subRouter.HandleFunc("/api/shutdown", this.apiShutdown).Methods("GET")
	subRouter.HandleFunc("/api/uninstall", this.upgrade.ApiUninstall).Methods("GET")
	subRouter.HandleFunc("/api/clear", this.apiClear).Methods("DELETE")
	subRouter.HandleFunc("/api/version", this.upgrade.ApiVersion).Methods("GET")
	subRouter.HandleFunc("/api/upgrade", this.upgrade.ApiUpdate).Methods("POST")
	subRouter.HandleFunc("/api/upgrade", this.upgrade.ApiUpdate).Methods("PUT")
	subRouter.HandleFunc("/api/server/config/get", this.apiServerConfigGet).Methods("GET")
	subRouter.HandleFunc("/api/server/config/set", this.apiServerConfigSet).Methods("PUT")
	subRouter.HandleFunc("/api/proxy/{type}", this.apiProxyByType).Methods("GET")
	subRouter.HandleFunc("/api/bindinfo", this.apiBindInfo).Methods("GET")
}

func (this *frps) userHandlers(helper *httppkg.RouterRegisterHelper) {
	subRouter := helper.Router.NewRoute().Name("user").Subrouter()
	subRouter.Use(helper.AuthMiddleware)
	// apis
	subRouter.HandleFunc("/api/token/add", this.apiUserCreate).Methods("POST")
	subRouter.HandleFunc("/api/token/del", this.apiUserDelete).Methods("POST")
	subRouter.HandleFunc("/api/token/delall", this.apiUserDeleteAll).Methods("POST")
	subRouter.HandleFunc("/api/token/chg", this.apiUserUpdate).Methods("POST")
	subRouter.HandleFunc("/api/token/all", this.apiUserAll).Methods("GET")

	subRouter.HandleFunc("/api/client/get", this.apiClientGet).Methods("GET")
	subRouter.HandleFunc("/api/client/gen", this.apiClientGen).Methods("POST")
	subRouter.HandleFunc("/api/frps/get", this.apiFrpsGet).Methods("GET")
	subRouter.HandleFunc("/api/frps/gen", this.apiFrpsGen).Methods("POST")
	subRouter.HandleFunc("/api/client/toml", this.apiClientToml).Methods("POST")
	subRouter.HandleFunc("/api/client/user/import", this.apiClientUserImport).Methods("POST")
	subRouter.HandleFunc("/api/client/user/export", this.apiClientUserExport).Methods("POST")
	subRouter.HandleFunc("/api/config/upload", this.apiConfigUpload)
	subRouter.HandleFunc("/api/config/upgrade", this.apiConfigUpgrade)
	subRouter.HandleFunc("/api/client/upload", this.apiClientUpload).Methods("POST")
}

func (this *frps) webSocketHandler(helper *httppkg.RouterRegisterHelper) {
	subRouter := helper.Router.NewRoute().Name("frpwebsocket").Subrouter()
	subRouter.Use(helper.AuthMiddleware)
	subRouter.HandleFunc("/frp", this.webSocket.HandleConnections).Methods("GET")
}
