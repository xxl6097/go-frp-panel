package frps

import (
	httppkg "github.com/fatedier/frp/pkg/util/http"
	"github.com/xxl6097/go-frp-panel/cmd/proxy/proxy"
	"net/http"
)

func (this *frps) proxyHandlers(helper *httppkg.RouterRegisterHelper) {
	subRouter := helper.Router.NewRoute().Name("admin").Subrouter()
	subRouter.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		proxy.GetInstance().Serve(w, r)
	})
}
