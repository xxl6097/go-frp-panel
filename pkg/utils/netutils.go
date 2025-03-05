package utils

import (
	"fmt"
	"os/exec"
	"runtime"
)

func GetListeningPorts() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("netstat", "-ano")
	} else {
		cmd = exec.Command("lsof", "-i", "-P", "-n")
	}
	output, _ := cmd.CombinedOutput()
	fmt.Println(string(output)) // 解析输出以提取端口和进程信息
}
