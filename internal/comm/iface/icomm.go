package iface

import "net/http"

type IComm interface {
	ApiUpdate(w http.ResponseWriter, r *http.Request)
	ApiRestart(w http.ResponseWriter, r *http.Request)
	ApiVersion(w http.ResponseWriter, r *http.Request)
}
