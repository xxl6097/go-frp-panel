package comm

import (
	"encoding/json"
	"github.com/xxl6097/glog/glog"
	"github.com/xxl6097/go-frp-panel/pkg/utils"
	"net/http"
)

type Option struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type GeneralResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
	Raw  []byte `json:"-"`
}

func (g *GeneralResponse) response(code int, msg string, data any) *GeneralResponse {
	g.Msg = msg
	g.Code = code
	g.Data = data
	return g
}
func (g *GeneralResponse) Response(code int, msg string) *GeneralResponse {
	return g.response(code, msg, nil)
}
func (g *GeneralResponse) Err(err error) *GeneralResponse {
	return g.Response(-1, err.Error())
}
func (g *GeneralResponse) Error(msg string) *GeneralResponse {
	return g.Response(-1, msg)
}
func (g *GeneralResponse) Any(data any) *GeneralResponse {
	return g.response(0, "ok", data)
}
func (g *GeneralResponse) Ok(msg string) *GeneralResponse {
	return g.Response(0, msg)
}
func (g *GeneralResponse) Sucess(msg string, data any) *GeneralResponse {
	return g.response(0, msg, data)
}

func Response(r *http.Request) (*GeneralResponse, func(w http.ResponseWriter)) {
	res := &GeneralResponse{Code: 0}
	return res, func(w http.ResponseWriter) {
		defer func() {
			if res.Code != 0 {
				glog.Errorf("Http response [%s]: res: %+v", r.URL.Path, res)
			}
		}()

		w.WriteHeader(200)
		var data []byte
		if res.Data == nil {
			res.Data = utils.GetTime()
		}
		if res.Raw != nil {
			data = res.Raw
			glog.Infof("Http response [%s %s]: raw: %s", r.Method, r.URL.Path, string(res.Raw))
		} else {
			//glog.Infof("Http response [%s %s]: res: %v", r.Method, r.URL.Path, res)
			bb, err := json.Marshal(res)
			if err != nil {
				glog.Errorf("marshal result error: %v", err)
				w.WriteHeader(400)
				return
			}
			data = bb
		}
		if len(data) > 0 {
			_, _ = w.Write(data)
		}
	}
}
