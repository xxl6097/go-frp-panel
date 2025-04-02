package pkg

import (
	"fmt"
	"runtime"
	"strings"
)

func init() {
	OsType = runtime.GOOS
	Arch = runtime.GOARCH
}

var (
	AppName      string // 应用名称
	AppVersion   string // 应用版本
	BuildVersion string // 编译版本
	BuildTime    string // 编译时间
	GitRevision  string // Git版本
	GitBranch    string // Git分支
	GoVersion    string // Golang信息
	DisplayName  string // 服务显示名
	Description  string // 服务描述信息
	OsType       string // 操作系统
	Arch         string // cpu类型
	BinName      string // 运行文件名称，包含平台架构
)

// Version 版本信息
func Version() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("App Name:\t%s\n", AppName))
	sb.WriteString(fmt.Sprintf("App Version:\t%s\n", AppVersion))
	sb.WriteString(fmt.Sprintf("Build version:\t%s\n", BuildVersion))
	sb.WriteString(fmt.Sprintf("Build time:\t%s\n", BuildTime))
	sb.WriteString(fmt.Sprintf("Git revision:\t%s\n", GitRevision))
	sb.WriteString(fmt.Sprintf("Git branch:\t%s\n", GitBranch))
	sb.WriteString(fmt.Sprintf("Golang Version: %s\n", GoVersion))
	sb.WriteString(fmt.Sprintf("DisplayName:\t%s\n", DisplayName))
	sb.WriteString(fmt.Sprintf("Description:\t%s\n", Description))
	sb.WriteString(fmt.Sprintf("OsType:\t%s\n", OsType))
	sb.WriteString(fmt.Sprintf("Arch:\t%s\n", Arch))
	sb.WriteString(fmt.Sprintf("BinName:\t%s\n", BinName))
	fmt.Println(sb.String())
	return sb.String()
}
