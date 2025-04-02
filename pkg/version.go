package pkg
import (
	"fmt"
	"strings"
	"runtime"
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
	sb.WriteString(fmt.Sprintf("%-15s: %-5s\n", "App Name", AppName))
	sb.WriteString(fmt.Sprintf("%-15s: %-5s\n", "App Version", AppVersion))
	sb.WriteString(fmt.Sprintf("%-15s: %-5s\n", "Build version", BuildVersion))
	sb.WriteString(fmt.Sprintf("%-15s: %-5s\n", "Build time", BuildTime))
	sb.WriteString(fmt.Sprintf("%-15s: %-5s\n", "Git revision", GitRevision))
	sb.WriteString(fmt.Sprintf("%-15s: %-5s\n", "Git branch", GitBranch))
	sb.WriteString(fmt.Sprintf("%-15s: %-5s\n", "Golang Version", GoVersion))
	sb.WriteString(fmt.Sprintf("%-15s: %-5s\n", "DisplayName", DisplayName))
	sb.WriteString(fmt.Sprintf("%-15s: %-5s\n", "Description", Description))
	sb.WriteString(fmt.Sprintf("%-15s: %-5s\n", "OsType", OsType))
	sb.WriteString(fmt.Sprintf("%-15s: %-5s\n", "Arch", Arch))
	sb.WriteString(fmt.Sprintf("%-15s: %-5s\n", "BinName", BinName))
	fmt.Println(sb.String())
	return sb.String()
}
