package comm

import (
	v1 "github.com/fatedier/frp/pkg/config/v1"
	"strconv"
	"strings"
)

type BufferConfig struct {
	Addr       string               `json:"addr"`
	Port       int                  `json:"port"`
	User       string               `json:"user"`
	Token      string               `json:"token"`
	ID         string               `json:"id"`
	Comment    string               `json:"comment,omitempty"`
	Ports      []any                `json:"ports,omitempty"`
	Domains    []string             `json:"domains,omitempty"`
	Subdomains []string             `json:"subdomains,omitempty"`
	Proxy      *v1.TypedProxyConfig `json:"proxy"`
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
