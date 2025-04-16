package main

import (
	"fmt"
	"github.com/xxl6097/go-frp-panel/pkg/utils"
)

// env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
func main() {
	name := "acfrps_v0.1.91_windows_amd64.exe"
	result := utils.SplitLastTwoByUnderscore(name)
	fmt.Printf("%-30s => %v\n", name, result)
	fmt.Printf("%s\n", utils.CleanExt(result[1]))
}
