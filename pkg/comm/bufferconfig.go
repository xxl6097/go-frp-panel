package comm

import (
	v1 "github.com/fatedier/frp/pkg/config/v1"
	"strconv"
	"strings"
)

type BufferConfig struct {
	Addr       string               `json:"addr"`
	Port       int                  `json:"port"`
	ApiPort    int                  `json:"apiPort"`
	User       string               `json:"user"`
	Token      string               `json:"token"`
	ID         string               `json:"id"`
	Comment    string               `json:"comment,omitempty"`
	Ports      []any                `json:"ports,omitempty"`
	Domains    []string             `json:"domains,omitempty"`
	Subdomains []string             `json:"subdomains,omitempty"`
	Proxy      *v1.TypedProxyConfig `json:"proxy"`
	WebServer  *v1.WebServerConfig  `json:"webserver"`
}

func (u *BufferConfig) ParsePorts() []int {
	ports := []int{}
	for _, port := range u.Ports {
		if str, ok := port.(string); ok {
			if strings.Contains(str, "-") {
				allowedRanges := strings.Split(str, "-")
				if len(allowedRanges) != 2 {
					break
				}
				start, err := strconv.Atoi(strings.TrimSpace(allowedRanges[0]))
				if err != nil {
					break
				}
				end, err := strconv.Atoi(strings.TrimSpace(allowedRanges[1]))
				if err != nil {
					break
				}
				for i := min(start, end); i <= max(start, end); i++ {
					ports = append(ports, i)
				}
			} else {
				if str == "" {
					break
				}
				allowed, err := strconv.Atoi(str)
				if err != nil {
					break
				}
				ports = append(ports, allowed)
			}
		} else {
			num, okk := port.(float64)
			if okk {
				ports = append(ports, int(num))
				break
			}
		}

	}
	return ports
}

func HasProxyes(p *v1.TypedProxyConfig) bool {
	if p == nil {
		return false
	}
	pc := p.ProxyConfigurer
	if pc == nil {
		return false
	}
	switch v := pc.(type) {
	case *v1.TCPProxyConfig:
		if v == nil {
			return false
		}
		if v.RemotePort == 0 {
			return false
		}
	case *v1.UDPProxyConfig:
		if v == nil {
			return false
		}
		if v.RemotePort == 0 {
			return false
		}
	}
	bc := pc.GetBaseConfig()
	if bc == nil {
		return false
	}
	if bc.Name == "" {
		return false
	}
	if bc.Type == "" {
		return false
	}
	if bc.LocalIP == "" {
		return false
	}
	if bc.LocalPort == 0 {
		return false
	}
	return true
}
