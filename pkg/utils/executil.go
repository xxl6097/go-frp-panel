package utils

import (
	"fmt"
	"github.com/xxl6097/go-service/gservice/utils"
	"os"
	"syscall"
	"time"
)

type GeneralResponse struct {
	Pid int
	Msg string
}

var proc *os.Process

// RestartProgram 重启程序
func RestartProgram() error {
	var er error
	// 检查是否需要重启
	if len(os.Args) > 1 && os.Args[1] == "--restart" {
		fmt.Println("程序重启成功")
		return nil
	} else {
		time.Sleep(2 * time.Second)
		fmt.Println("准备重启程序...")
		args := append([]string{os.Args[0]}, "--restart")
		env := os.Environ()
		attr := syscall.SysProcAttr{}
		if !utils.IsWindows() {
			//attr.Setpgid = true
		}
		procAttr := &os.ProcAttr{
			Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
			Sys:   &attr,
			Env:   env,
		}

		process, err := os.StartProcess(os.Args[0], args, procAttr)
		proc = process
		if err != nil {
			fmt.Println("重启失败:", err)
			return err
		}
		er = process.Release()
		os.Exit(0)
	}
	return er
}

func Shutdown() error {
	pid := os.Getpid()
	process, err := os.FindProcess(pid)
	fmt.Printf("pid %v process: %v proc: %v\n", pid, process, proc)
	if err != nil {
		return err
	}
	err = process.Kill()
	if err != nil {
		return err
	}
	return process.Release()
}
