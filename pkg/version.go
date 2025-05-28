package pkg

import (
	"fmt"
	"runtime"
	"strings"
)

func init() {
	OsType = runtime.GOOS
	Arch = runtime.GOARCH
	GoVersion = runtime.Version()
	Compiler = runtime.Compiler
}

var (
	AppName          string // 应用名称
	AppVersion       string // 应用版本
	BuildVersion     string // 编译版本
	BuildTime        string // 编译时间
	GoVersion        string // Golang信息
	DisplayName      string // 服务显示名
	Description      string // 服务描述信息
	OsType           string // 操作系统
	Arch             string // cpu类型
	Compiler         string // 编译器信息
	GitRevision      string // Git版本
	GitBranch        string // Git分支
	GitTreeState     string // state of git tree, either "clean" or "dirty"
	GitCommit        string // sha1 from git, output of 4a2ea0514582c5bdf629ad348341970c5ea8fdc6
	GitVersion       string // semantic version, derived by build scripts
	GitReleaseCommit string
	BinName          string // 运行文件名称，包含平台架构
	GithubUser       string // github用户
	GithubRepo       string // github项目名称
)

// Version 版本信息
func Version() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("%-16s: %-5s\n", "App Name", AppName))
	sb.WriteString(fmt.Sprintf("%-16s: %-5s\n", "App Version", AppVersion))
	sb.WriteString(fmt.Sprintf("%-16s: %-5s\n", "DisplayName", DisplayName))
	sb.WriteString(fmt.Sprintf("%-16s: %-5s\n", "Description", Description))
	sb.WriteString(fmt.Sprintf("%-16s: %-5s\n", "Build version", BuildVersion))
	sb.WriteString(fmt.Sprintf("%-16s: %-5s\n", "Build time", BuildTime))
	sb.WriteString(fmt.Sprintf("%-16s: %-5s\n", "Golang Version", GoVersion))
	sb.WriteString(fmt.Sprintf("%-16s: %-5s\n", "OsType", OsType))
	sb.WriteString(fmt.Sprintf("%-16s: %-5s\n", "Arch", Arch))
	sb.WriteString(fmt.Sprintf("%-16s: %-5s\n", "Compiler", Compiler))
	sb.WriteString(fmt.Sprintf("%-16s: %-5s\n", "Git revision", GitRevision))
	sb.WriteString(fmt.Sprintf("%-16s: %-5s\n", "Git branch", GitBranch))
	sb.WriteString(fmt.Sprintf("%-16s: %-5s\n", "GitTreeState", GitTreeState))
	sb.WriteString(fmt.Sprintf("%-16s: %-5s\n", "GitCommit", GitCommit))
	sb.WriteString(fmt.Sprintf("%-16s: %-5s\n", "GitVersion", GitVersion))
	sb.WriteString(fmt.Sprintf("%-16s: %-5s\n", "GitReleaseCommit", GitReleaseCommit))
	sb.WriteString(fmt.Sprintf("%-16s: %-5s\n", "BinName", BinName))
	sb.WriteString(fmt.Sprintf("%-16s: %-5s\n", "GithubUser", GithubUser))
	sb.WriteString(fmt.Sprintf("%-16s: %-5s\n", "GithubRepo", GithubRepo))
	fmt.Println(sb.String())
	return sb.String()
}
