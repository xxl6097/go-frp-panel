package main

import (
	"fmt"
	"github.com/xxl6097/go-frp-panel/pkg/utils"
	"os"
	"syscall"
	"time"
)

func test() {
	fmt.Println("程序启动")
	filePath := "/Users/uuxia/Desktop/work/code/go/go-frp-panel/test.txt"
	utils.AppendStringToFile(filePath, "程序启动 "+utils.GetTime()+"\n")
	for {
		ppid := os.Getppid()
		pid := os.Getpid()
		text := fmt.Sprintf("运行中 %s pid:%d ppid:%d", utils.GetTime(), pid, ppid)
		fmt.Println(text)
		utils.AppendStringToFile(filePath, text+"\n")
		time.Sleep(1 * time.Second)
	}
	fmt.Println("退出")
	utils.AppendStringToFile(filePath, "退出 "+utils.GetTime()+"\n")
}

func restart() {
	filePath := "/Users/uuxia/Desktop/work/code/go/go-frp-panel/test.txt"
	// 检查是否需要重启
	if len(os.Args) > 1 && os.Args[1] == "--restart" {
		fmt.Println("程序重启成功")
		utils.AppendStringToFile(filePath, "程序重启成功 "+utils.GetTime()+"\n")
	} else {
		time.Sleep(2 * time.Second)
		fmt.Println("准备重启程序...")
		utils.AppendStringToFile(filePath, "准备重启程序 "+utils.GetTime()+"\n")
		args := append([]string{os.Args[0]}, "--restart")
		env := os.Environ()
		procAttr := &os.ProcAttr{
			Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
			Sys:   &syscall.SysProcAttr{Setpgid: true},
			Env:   env,
		}
		process, err := os.StartProcess(os.Args[0], args, procAttr)
		if err != nil {
			fmt.Println("重启失败:", err)
			return
		}
		utils.AppendStringToFile(filePath, "重启ok "+utils.GetTime()+"\n")
		process.Release()
		time.Sleep(1 * time.Second)
		os.Exit(0)
	}
}

func testRestart() {

	go restart()
	test()
}

func main() {
	a := []byte{'\x11'}
	fmt.Println(a)
}
