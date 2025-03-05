package main

import (
	"fmt"
	"net"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"
)

func GetListeningPorts() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("netstat", "-ano")
	} else {
		cmd = exec.Command("lsof", "-i", "-P", "-n")
	}
	output, _ := cmd.CombinedOutput()
	lines := strings.Split(string(output), "\n")
	fmt.Println(string(output)) // 解析输出以提取端口和进程信息
	// 解析输出，筛选 TCP 监听行
	for _, line := range lines {
		if strings.Contains(line, "tcp") {
			fmt.Println(line) // 输出格式如 "tcp 0 0 0.0.0.0:8080 0.0.0.0:* LISTEN"
		}
	}
}

// 扫描单个端口
func scanPort(ip string, port int, timeout time.Duration, wg *sync.WaitGroup) net.Conn {
	defer wg.Done()
	address := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return nil
	}
	defer conn.Close()
	fmt.Printf("Port %d is open %v\n", port, conn.RemoteAddr())
	return conn
}

func scanPorts(host string, start, end int) {
	var wg sync.WaitGroup
	// 控制并发的 goroutine 数量，避免打开过多文件描述符
	sem := make(chan struct{}, 1000)

	connArray := map[int]net.Conn{}
	for port := start; port <= end; port++ {
		wg.Add(1)
		sem <- struct{}{}
		go func(p int) {
			defer func() { <-sem }()
			conn := scanPort(host, p, 2*time.Second, &wg)
			if conn != nil {
				connArray[port] = conn
			}
		}(port)
	}
	for port := range connArray {
		fmt.Printf("Port %d is open %v\n", port, connArray[port].RemoteAddr())
	}
	wg.Wait()
}
func main() {
	rawURL := "C://test.txt"

	// 提取路径部分并获取文件名
	fileName := filepath.Ext(rawURL)
	fmt.Println("文件名:", fileName) // 输出: document.pdf

	host := "192.168.0.4"
	scanPorts(host, 0, 65535)
}
