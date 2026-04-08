package frps

import (
	"embed"

	"github.com/fatedier/frp/assets"
)

//go:embed dist/*
var content embed.FS

func init() {
	assets.Register(content)
}
