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
