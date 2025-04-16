package utils

import (
	"github.com/xxl6097/glog/glog"
	"github.com/xxl6097/go-service/gservice/gore/util"
	"github.com/xxl6097/go-service/gservice/utils"
)

func ShowUpDirSize() {
	updir := utils.GetUpgradeDir()
	total, used, free, err := util.GetDiskUsage(updir)
	glog.Printf("Current Working Directory: %s %v\n", updir, err)
	glog.Printf("Total space: %d bytes %v\n", total, float64(total)/1024/1024/1024)
	glog.Printf("Used space: %d bytes %v\n\n", used, float64(used)/1024/1024/1024)
	glog.Printf("Free space: %d bytes %v\n\n", free, float64(free)/1024/1024/1024)
}

func HasDiskSpace() bool {
	size := GetSelfSize()
	size *= 16
	dir := glog.GetCrossPlatformDataDir()
	total, used, free, err := util.GetDiskUsage(dir)
	glog.Printf("Current Working Directory: %s %v\n", dir, err)
	glog.Printf("Total space: %d bytes %v\n", total, float64(total)/1024/1024/1024)
	glog.Printf("Used space: %d bytes %v\n\n", used, float64(used)/1024/1024/1024)
	glog.Printf("Free space: %d bytes %v\n\n", free, float64(free)/1024/1024/1024)
	if free > size {
		return true
	}
	return false
}
