package utils

import (
	"fmt"
	"github.com/xxl6097/glog/glog"
	"github.com/xxl6097/go-frp-panel/pkg/model"
	utils2 "github.com/xxl6097/go-service/gservice/utils"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func Export(obj model.CloudApi) error {
	binpath, err := os.Executable()
	if err != nil {
		glog.Error(err)
		return err
	}
	userDir := filepath.Join(filepath.Dir(binpath), "user")
	fileName := fmt.Sprintf("user_%s.zip", GetFileNameByTime())
	tempDir := filepath.Join(glog.GetCrossPlatformDataDir(), "user")
	_ = utils2.EnsureDir(tempDir)
	zipFilePath := filepath.Join(tempDir, fileName)
	err = Zip(userDir, zipFilePath)
	if err != nil {
		glog.Error("GetDataByJson", err)
		return err
	}
	defer utils2.Delete(zipFilePath, "用户配置")
	err = UploadGeneric(obj.Addr, "PUT", zipFilePath, obj.User, obj.Pass)
	if strings.Contains(obj.Addr, "latest") {
		version := time.Now().Format("2006.01.02.15.04.05")
		obj.Addr = strings.ReplaceAll(obj.Addr, "latest", version)
		err = UploadGeneric(obj.Addr, "PUT", zipFilePath, obj.User, obj.Pass)
	}
	if err != nil {
		return err
	}
	return nil
}

func Import(obj model.CloudApi) error {
	dstFilePath := filepath.Join(glog.GetCrossPlatformDataDir("temp"), "user_import.zip")
	err := DownLoadGeneric(obj.Addr, obj.User, obj.Pass, dstFilePath)
	if err != nil {
		return err
	}
	defer Delete(dstFilePath, "用户文件")
	binpath, err := os.Executable()
	if err != nil {
		glog.Error(binpath, err)
		return err
	}

	userDir := filepath.Join(filepath.Dir(binpath), "user")
	err = UnzipToRoot(dstFilePath, userDir, true)
	if err != nil {
		return err
	}
	glog.Info("解压成功", userDir)
	return nil
}
