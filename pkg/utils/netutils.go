package utils

import (
	"errors"
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
		//fmt.Printf("isURLAccessible: Head response error: %v, %v", err, urlStr)
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

func GetLocalMac() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("获取网络接口失败：", err)
		return ""
	}
	for _, iface := range interfaces {
		if iface.Flags&net.FlagUp != 0 && iface.HardwareAddr != nil {
			devMac := strings.ReplaceAll(iface.HardwareAddr.String(), ":", "")
			fmt.Println(iface.Name, ":", devMac)
			return devMac
		}
	}
	return ""
}
func GetLocalIp() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		return ""
	}
	for _, iface := range interfaces {
		addrs, err := iface.Addrs()
		if err != nil {
			return ""
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip != nil && ip.To4() != nil {
				fmt.Println(ip)
				return ip.String()
			}
		}
	}
	return ""
}

// NetworkInterface 网络接口信息
type NetworkInterface struct {
	Name        string   `json:"name"`        // 接口名称
	DisplayName string   `json:"displayName"` // 显示名称
	MacAddress  string   `json:"macAddress"`  // MAC地址
	Ipv4        string   `json:"ipv4"`        // MAC地址
	IPAddresses []string `json:"ipAddresses"` // IP地址列表
}

// GetNetworkInterfaces 获取所有网络接口信息
func GetNetworkInterfaces() ([]NetworkInterface, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, fmt.Errorf("获取网络接口失败: %w", err)
	}

	var result []NetworkInterface
	for _, iface := range ifaces {
		// 忽略回环接口
		if iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		// 忽略未激活的接口
		if iface.Flags&net.FlagUp == 0 {
			continue
		}

		// 获取接口地址
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}

		var ipAddresses []string
		var ipv4 string
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			// 忽略回环地址
			if ip.IsLoopback() {
				continue
			}

			// 获取IPv4和IPv6地址
			if ip.To4() != nil {
				ipAddresses = append(ipAddresses, ip.String())
				ipv4 = ip.String()
			} else if ip.To16() != nil {
				ipAddresses = append(ipAddresses, "["+ip.String()+"]")
			}
		}

		// 忽略没有IP地址的接口
		if len(ipAddresses) == 0 {
			continue
		}

		result = append(result, NetworkInterface{
			Name:        iface.Name,
			DisplayName: getInterfaceDisplayName(iface.Name),
			MacAddress:  strings.ReplaceAll(iface.HardwareAddr.String(), ":", ""),
			Ipv4:        ipv4,
			IPAddresses: ipAddresses,
		})
	}

	if len(result) == 0 {
		return nil, errors.New("未找到可用的网络接口")
	}

	return result, nil
}

// getInterfaceDisplayName 获取接口的显示名称（跨平台适配）
func getInterfaceDisplayName(name string) string {
	// 为常见接口类型提供友好名称
	switch {
	case strings.HasPrefix(name, "eth"):
		return "以太网"
	case strings.HasPrefix(name, "wlan") || strings.HasPrefix(name, "wifi"):
		return "无线局域网"
	case strings.HasPrefix(name, "en"):
		return "以太网"
	case strings.HasPrefix(name, "wl"):
		return "无线局域网"
	case strings.HasPrefix(name, "vEthernet"):
		return "虚拟以太网"
	default:
		return name
	}
}

// GetPrimaryIP 获取主IP地址（默认网关所在接口的IP）
func GetPrimaryIP() (string, error) {
	ifaces, err := GetNetworkInterfaces()
	if err != nil {
		return "", err
	}

	if len(ifaces) == 0 {
		return "", errors.New("未找到网络接口")
	}

	// 优先选择非虚拟接口
	for _, iface := range ifaces {
		if !strings.Contains(strings.ToLower(iface.Name), "virtual") &&
			!strings.Contains(strings.ToLower(iface.Name), "vmware") &&
			!strings.Contains(strings.ToLower(iface.Name), "docker") {

			for _, ip := range iface.IPAddresses {
				if !strings.Contains(ip, ":") { // 优先返回IPv4
					return ip, nil
				}
			}

			// 如果没有IPv4，返回第一个IP
			if len(iface.IPAddresses) > 0 {
				return iface.IPAddresses[0], nil
			}
		}
	}

	// 如果没有找到非虚拟接口，返回第一个接口的IP
	return ifaces[0].IPAddresses[0], nil
}

// GetDeviceInfo 获取主IP地址、Mac地址（默认网关所在接口的IP）
func GetDeviceInfo() (*NetworkInterface, error) {
	ifaces, err := GetNetworkInterfaces()
	if err != nil {
		return nil, err
	}

	if len(ifaces) == 0 {
		return nil, errors.New("未找到网络接口")
	}

	face := &ifaces[0]
	// 优先选择非虚拟接口
	for _, iface := range ifaces {
		if !strings.Contains(strings.ToLower(iface.Name), "virtual") &&
			!strings.Contains(strings.ToLower(iface.Name), "vmware") &&
			!strings.Contains(strings.ToLower(iface.Name), "docker") {
			//return &iface, nil
			face = &iface
		}
		glog.Debugf("获取设备信息：%+v", iface)
	}

	return face, nil
}
