package main

import (
	"nas-file-tool/internal"
	"nas-file-tool/internal/menu"
)

func main() {
	internal.Registe()
	menu.GetInstance().ShowMenu()
}
