package frp

import (
	"fmt"
	"github.com/xxl6097/glog/glog"
	"github.com/xxl6097/go-frp-panel/pkg/utils"
	utils2 "github.com/xxl6097/go-service/gservice/utils"
	"os"
	"path/filepath"
)

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
	err := utils.MakeDir(dir)
	if err != nil {
		return fmt.Errorf("make dir err: %v", err)
	}
	cfgPath := filepath.Join(dir, GetFrpcMainTomlFileName())
	fmt.Println(cfgPath)
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
