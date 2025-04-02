package utils

import (
	"fmt"
	"github.com/xxl6097/glog/glog"
	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"
)

func isURLFormatValid(urlStr string) bool {
	// 严格解析绝对 URL
	if _, err := url.ParseRequestURI(urlStr); err != nil {
		glog.Errorf("isURLFormatValid: ParseRequestURI error: %v, %v", err, urlStr)
		return false
	}
	// 提取协议和主机名
	parsed, err := url.Parse(urlStr)
	if err != nil {
		glog.Errorf("isURLFormatValid: Parse error: %v, %v", err, urlStr)
	}
	return err == nil && parsed.Scheme != "" && parsed.Host != ""
}
func isURLAccessible(urlStr string) bool {
	client := &http.Client{
		Timeout: 10 * time.Second, // 超时控制
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse // 禁用重定向
		},
	}
	resp, err := client.Head(urlStr)
	if err != nil || resp.StatusCode >= 400 {
		fmt.Printf("isURLAccessible: Head response error: %v, %v", err, urlStr)
		return false
	}
	defer resp.Body.Close()
	return true
}

// IsURLValidAndAccessible 检查 URL 是否有效并且可访问
func IsURLValidAndAccessible(rawURL string) bool {
	// 阶段1：格式校验
	if !isURLFormatValid(rawURL) {
		return false
	}
	// 阶段2：网络可达性检测
	return isURLAccessible(rawURL)
}
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

func PingRaw(ip string) bool {
	conn, _ := icmp.ListenPacket("udp4", "0.0.0.0")
	defer conn.Close()

	msg := icmp.Message{
		Type: ipv4.ICMPTypeEcho,
		Code: 0,
		Body: &icmp.Echo{ID: os.Getpid() & 0xffff, Seq: 1},
	}
	wb, _ := msg.Marshal(nil)

	if _, err := conn.WriteTo(wb, &net.UDPAddr{IP: net.ParseIP(ip)}); err != nil {
		return false
	}

	// 设置超时并读取响应
	conn.SetReadDeadline(time.Now().Add(500 * time.Millisecond))
	reply := make([]byte, 1500)
	_, _, err := conn.ReadFrom(reply)
	return err == nil
}

func ping(ip string) bool {

	// 创建 ICMP 连接
	conn, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil {
		//fmt.Printf("[%s] 错误: %v\n", ip, err)
		return false
	}
	defer conn.Close()

	// 构造 ICMP 消息
	msg := icmp.Message{
		Type: ipv4.ICMPTypeEcho, Code: 0,
		Body: &icmp.Echo{
			ID:   os.Getpid() & 0xffff,
			Seq:  1,
			Data: []byte("HELLO"),
		},
	}
	msgBytes, _ := msg.Marshal(nil)

	// 发送请求
	if _, err := conn.WriteTo(msgBytes, &net.IPAddr{IP: net.ParseIP(ip)}); err != nil {
		//fmt.Printf("[%s] 发送失败: %v\n", ip, err)
		return false
	}

	// 设置超时并等待响应
	reply := make([]byte, 1500)
	conn.SetReadDeadline(time.Now().Add(1 * time.Second))
	_, _, err = conn.ReadFrom(reply)
	if err != nil {
		return false
	}
	//fmt.Printf("[+] %s 在线\n", ip)
	return true
}

// 判断 ping 输出是否表示成功
func isPingSuccessful(output string) bool {
	// 不同系统 ping 成功的输出关键字不同
	// Windows 包含 "Reply from"
	// Linux 和 macOS 包含 "bytes from"
	return strings.Contains(strings.ToLower(output), "reply from") ||
		strings.Contains(strings.ToLower(output), "bytes from")
}

// 扫描指定 IP 是否活跃
func scanIP(ip string) bool {
	var cmd *exec.Cmd
	// 根据不同操作系统选择不同的 ping 命令参数
	// Windows 系统使用 -n 1 表示只发送一个数据包，-w 1000 表示超时时间为 1 秒
	// Linux 和 macOS 使用 -c 1 表示只发送一个数据包，-W 1 表示超时时间为 1 秒
	switch runtime.GOOS {
	case "windows":
		args := []string{"-n", "1", "-w", "10000", ip}
		fmt.Println("ping", args)
		cmd = exec.Command("ping", args...)
	default:
		args := []string{"-c", "1", "-W", "10", ip}
		//fmt.Println("ping", args)
		cmd = exec.Command("ping", args...)
	}
	output, err := cmd.CombinedOutput()
	if err != nil {
		return false
	}
	//fmt.Println(ip, string(output))
	// 检查输出中是否包含表示成功的关键字
	if isPingSuccessful(string(output)) {
		//fmt.Printf("Active host: %s\n", ip)
		return true
	}
	return false
}

func IsPortOpen(host string, port int, timeout time.Duration) bool {
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return false
	}
	defer conn.Close()
	fmt.Printf("port %d is open %v\n", port, conn.RemoteAddr())
	return true
}

// 扫描单个端口
func scanPort(ip string, port int, timeout time.Duration, wg *sync.WaitGroup) bool {
	defer wg.Done()
	return IsPortOpen(ip, port, timeout)
}

func ScanPort(host string, duration time.Duration, start, end int) []int {
	var wg sync.WaitGroup
	// 控制并发的 goroutine 数量，避免打开过多文件描述符
	sem := make(chan struct{}, 1000)

	connArray := []int{}
	for port := start; port <= end; port++ {
		wg.Add(1)
		sem <- struct{}{}
		go func(p int) {
			defer func() { <-sem }()
			conn := scanPort(host, p, duration, &wg)
			if conn {
				connArray = append(connArray, p)
			}
		}(port)
	}
	sort.Ints(connArray) // 升序
	//for port := range connArray {
	//	fmt.Printf("Port %d is open %v\n", port, connArray[port].RemoteAddr())
	//}
	wg.Wait()
	return connArray
}

func ScanPorts(host string, start, end int) []int {
	return ScanPort(host, time.Millisecond*200, start, end)
}

func ScanIP() []string {
	ips := []string{}
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Error getting interfaces:", err)
		return nil
	}

	for _, iface := range ifaces {
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if ip == nil || ip.IsLoopback() {
				continue
			}

			ip = ip.To4()
			if ip == nil {
				continue
			}

			// 生成局域网内的 IP 地址
			network := ip.Mask(net.CIDRMask(24, 32))
			var wg sync.WaitGroup
			for i := 1; i < 255; i++ {
				ip := net.IPv4(network[0], network[1], network[2], byte(i))
				wg.Add(1)
				go func() {
					defer wg.Done()
					tempIp := ip.String()
					ok := scanIP(tempIp)
					//ok := ping(tempIp)
					if ok {
						fmt.Println(ok, tempIp)
						ips = append(ips, tempIp)
					}
				}()
			}
			wg.Wait()
			//sort.Strings(ips)
			//fmt.Println("IPS:", ips)
		}
	}
	return ips
}
