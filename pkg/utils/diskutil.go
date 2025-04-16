package utils

import (
	"fmt"
	"github.com/xxl6097/glog/glog"
	"github.com/xxl6097/go-service/gservice/gore/util"
	"github.com/xxl6097/go-service/gservice/utils"
	"os"
	"path/filepath"
)

func ShowUpDirSize() {
	updir := utils.GetUpgradeDir()
	total, used, free, err := util.GetDiskUsage(updir)
	glog.Printf("Current Working Directory: %s %v\n", updir, err)
	glog.Printf("Total space: %d bytes %v\n", total, ByteCountIEC(total))
	glog.Printf("Used space: %d bytes %v\n\n", used, ByteCountIEC(used))
	glog.Printf("Free space: %d bytes %v\n\n", free, ByteCountIEC(free))
}

func HasDiskSpace() bool {
	size := GetSelfSize()
	size *= 16
	dir := glog.GetCrossPlatformDataDir()
	total, used, free, err := util.GetDiskUsage(dir)
	glog.Printf("Current Working Directory: %s %v\n", dir, err)
	glog.Printf("Total space: %d bytes %v\n", total, ByteCountIEC(total))
	glog.Printf("Used space: %d bytes %v\n\n", used, ByteCountIEC(used))
	glog.Printf("Free space: %d bytes %v\n\n", free, ByteCountIEC(free))
	if free > size {
		return true
	}
	return false
}

func ClearDir(tmpPath string) error {
	entries, err := os.ReadDir(tmpPath)
	if err != nil {
		return fmt.Errorf("读取目录失败: %v", err)
	}

	for _, entry := range entries {
		fullPath := filepath.Join(tmpPath, entry.Name())
		err = os.RemoveAll(fullPath)
		if err != nil {
			return fmt.Errorf("删除 %s 失败: %v", fullPath, err)
		}
	}
	return nil
}

func ClearTmpDir() error {
	return ClearDir("/tmp")
}
