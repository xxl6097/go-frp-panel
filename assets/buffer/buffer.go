package buffer

import (
	_ "embed"
	"github.com/xxl6097/go-service/pkg/ukey"
)

//go:embed buffer
var buffer []byte

func init() {
	ukey.Register(buffer)
}
