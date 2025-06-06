package frp

import (
	"encoding/json"
	"fmt"
	v1 "github.com/fatedier/frp/pkg/config/v1"
	"github.com/xxl6097/glog/glog"
	"github.com/xxl6097/go-frp-panel/internal/com/model"
	"github.com/xxl6097/go-frp-panel/pkg"
	"github.com/xxl6097/go-frp-panel/pkg/utils"
	"github.com/xxl6097/go-service/pkg/github"
	utils2 "github.com/xxl6097/go-service/pkg/utils"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func init() {
	github.Api().SetName(pkg.GithubUser, pkg.GithubRepo)
}
func WriteFrpToml(cfgFilePath string, data any) error {
	if cfgFilePath == "" {
		return fmt.Errorf("cfgFilePath is empty")
	}
	director := filepath.Dir(cfgFilePath)
	err := utils.MakeDir(director)
	if err != nil {
		return fmt.Errorf("make dir err: %v", err)
	}
	switch val := data.(type) {
	case string:
		return utils.WriteToml(cfgFilePath, []byte(val))
	case []byte:
		return utils.WriteToml(cfgFilePath, val)
	case nil:
		return fmt.Errorf("data is nil")
	default:
		if err := os.WriteFile(cfgFilePath, utils.ObjectToTomlText(data), 0o600); err != nil {
			return fmt.Errorf("write content to frpc config file error: %v", err)
		} else {
			glog.Infof("write file success %s", cfgFilePath)
			return nil
		}
	}
}

func ReadFrpToml(cfgFileName string) ([]byte, error) {
	cfgDir, err := GetFrpcTomlDir()
	if err != nil {
		return nil, err
	}
	cfgFilePath := filepath.Join(cfgDir, cfgFileName)
	body, err := utils.ReadToml(cfgFilePath)
	if err != nil {
		return nil, fmt.Errorf("write http body err: %v", err)
	}
	return body, nil
}

func WriteFrpcMainConfig(data any) error {
	dir, err := GetFrpcTomlDir()
	if err != nil {
		glog.Error("os.Executable() error", err)
		return err
	}
	return WriteFrpcMainConfigWithDir(dir, data)
}

func WriteFrpcMainConfigWithDir(dir string, data any) error {
	cfgDir, err := GetFrpcTomlDirByDir(dir)
	if err != nil {
		return err
	}
	err = utils.MakeDir(dir)
	if err != nil {
		return fmt.Errorf("make dir err: %v", err)
	}
	cfgPath := filepath.Join(cfgDir, GetFrpcMainTomlFileName())
	return WriteFrpToml(cfgPath, data)
}

func WriteFrpcMainConfigWithOut(data any) error {
	dir, err := GetFrpcTomlDir()
	if err != nil {
		glog.Error("os.Executable() error", err)
		return err
	}
	cfgPath := filepath.Join(dir, GetFrpcMainTomlFileName())
	fmt.Println(cfgPath)
	if !utils2.FileExists(cfgPath) {
		return WriteFrpToml(cfgPath, data)
	}
	return nil
}

func GetFrpcMainTomlFileName() string {
	return "frpc.toml"
}

func GetFrpcMainTomlFilePath() (string, error) {
	dir, err := GetFrpcTomlDir()
	if err != nil {
		glog.Error("os.Executable() error", err)
		return "", err
	}
	cfgPath := filepath.Join(dir, GetFrpcMainTomlFileName())
	if !utils2.FileExists(cfgPath) {
		return "", fmt.Errorf("config file %s not exists", cfgPath)
	}
	return cfgPath, nil
}

func GetFrpcTomlDir() (string, error) {
	binpath, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf("get executable path err: %v", err)
	}
	return GetFrpcTomlDirByDir(filepath.Dir(binpath))
}

func GetFrpcTomlPath(name string) (string, error) {
	binpath, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf("get executable path err: %v", err)
	}
	cfgDir, err := GetFrpcTomlDirByDir(filepath.Dir(binpath))
	if err != nil {
		return "", fmt.Errorf("get executable path err: %v", err)
	}
	return filepath.Join(cfgDir, name), nil
}

func GetFrpcTomlDirByDir(bindir string) (string, error) {
	cfgDir := filepath.Join(bindir, "config")
	err := utils.MakeDir(cfgDir)
	if err != nil {
		return "", fmt.Errorf("make dir err: %v", err)
	}
	return cfgDir, nil
}

func GetMetadatas(token, id, apiPort, authorization string) map[string]string {
	return map[string]string{
		"token":         token,
		"id":            id,
		"apiPort":       apiPort,
		"authorization": authorization,
	}
}

func GetPort(i interface{}) int {
	switch v := i.(type) {
	case *v1.TCPProxyConfig:
		fmt.Printf("Received an TCPProxyConfig.RemotePort: %d\n", v.RemotePort)
		return v.RemotePort
	default:
		fmt.Println()
	}
	return 0
}

func EncodeSecret(obj *model.FrpcBuffer) (string, error) {
	if obj == nil {
		return "", fmt.Errorf("obj is nil")
	}

	glog.Debugf("EncodeSecret obj: %+v", obj)
	data, err := json.Marshal(obj)
	if err != nil {
		return "", fmt.Errorf("json marshal err: %v", err)
	}
	buffer, err := utils2.GzipCompress(data)
	if err != nil {
		return "", fmt.Errorf("gzip compress err: %v", err)
	}
	//glog.Println(string(data))
	return utils.Encrypt(buffer, nil), nil
}

func DecodeMetas(mapData map[string]string) *model.User {
	v, ok := mapData["secret"]
	if !ok {
		return nil
	}
	buffer := DecodeSecret(v)
	if buffer == nil {
		return nil
	}
	return &buffer.User
}

func DecodeSecret(text string) *model.FrpcBuffer {
	if text == "" {
		return nil
	}
	jsonString := utils.Decrypt(text, nil)
	if jsonString == "" {
		return nil
	}
	buffer, err := utils2.GzipDecompress([]byte(jsonString))
	if err != nil {
		glog.Errorf("gzip decompress err: %v", err)
		return nil
	}
	var obj model.FrpcBuffer
	err = json.Unmarshal(buffer, &obj)
	if err != nil {
		return nil
	}
	//glog.Debugf("%+v", obj)
	return &obj
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

func ParsePorts(ps []any) []int {
	var ports []int
	for _, port := range ps {
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
