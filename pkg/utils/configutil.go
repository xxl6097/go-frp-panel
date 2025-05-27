package utils

import (
	"fmt"
	"github.com/xxl6097/glog/glog"
	"github.com/xxl6097/go-frp-panel/pkg/model"
	"github.com/xxl6097/go-service/pkg/utils"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func Export(obj model.CloudApi) error {
	userDir, err := GetUserDir()
	if err != nil {
		return err
	}
	if ok, e := IsDirEmpty(userDir); ok || e != nil {
		glog.Error("IsDirEmpty", userDir)
		return err
	}
	fileName := fmt.Sprintf("user_%s.zip", GetFileNameByTime())
	tempDir := filepath.Join(glog.AppHome(), "user")
	_ = utils.ResetDirector(tempDir)
	zipFilePath := filepath.Join(tempDir, fileName)
	err = Zip(userDir, zipFilePath)
	if err != nil {
		glog.Error("GetDataByJson", err)
		return err
	}
	defer utils.DeleteAllDirector(zipFilePath)
	envType := os.Getenv("ENV_TYPE")
	if envType == "" {
		envType = "uuxia"
	}
	if strings.Contains(obj.Addr, "coding.net") {
		baseUrl := fmt.Sprintf("%s/%s_frps_config.zip?version=latest", obj.Addr, envType)
		err = UploadGeneric(baseUrl, "PUT", zipFilePath, obj.User, obj.Pass)
		version := time.Now().Format("2006.01.02.15.04.05")
		baseUrl = fmt.Sprintf("%s/%s_frps_config.zip?version=%s", obj.Addr, envType, version)
		err = UploadGeneric(baseUrl, "PUT", zipFilePath, obj.User, obj.Pass)
	} else {
		obj.Addr = fmt.Sprintf("%s/frp/config/%s_frps_config.zip", obj.Addr, envType)
		err = UploadGeneric(obj.Addr, "POST", zipFilePath, obj.User, obj.Pass)
	}
	if err != nil {
		return err
	}
	return nil
}

func Import(obj model.CloudApi) error {
	dstFilePath := filepath.Join(glog.AppHome("temp"), "user_import.zip")
	envType := os.Getenv("ENV_TYPE")
	if envType == "" {
		envType = "uuxia"
	}
	if strings.Contains(obj.Addr, "coding.net") {
		obj.Addr = fmt.Sprintf("%s/%s_frps_config.zip?version=latest", obj.Addr, envType)
	} else {
		obj.Addr = fmt.Sprintf("%s/frp/config/%s_frps_config.zip", obj.Addr, envType)
	}
	err := DownLoadGeneric(obj.Addr, obj.User, obj.Pass, dstFilePath)
	if err != nil {
		return err
	}
	defer Delete(dstFilePath, "用户文件")
	userDir, err := GetUserDir()
	if err != nil {
		return err
	}
	err = UnzipToRoot(dstFilePath, userDir, true)
	if err != nil {
		return err
	}
	glog.Info("解压成功", userDir)
	return nil
}
