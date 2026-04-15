package api

import (
	"sync"

	"github.com/fatedier/frp/client"
	v1 "github.com/fatedier/frp/pkg/config/v1"
)

type Callback func(svr *client.Service,
	cfg *v1.ClientCommonConfig,
	proxyCfgs []v1.ProxyConfigurer,
	visitorCfgs []v1.VisitorConfigurer) error
type OneInstance struct {
	Cb Callback
}

var instance *OneInstance
var once sync.Once

func GetInstance() *OneInstance {
	once.Do(func() {
		instance = &OneInstance{}
	})
	return instance
}
func (this *OneInstance) SetAdmin(a Callback) {
	this.Cb = a
}
