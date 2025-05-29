package utils

import (
	"github.com/xxl6097/glog/glog"
	"github.com/xxl6097/go-service/pkg/utils/util"
)

func ShowUpDirSize() {
	updir := glog.AppHome("temp", "upgrade")
	total, used, free, err := util.GetDiskUsage(updir)
	glog.Printf("Current Working Directory: %s %v\n", updir, err)
	glog.Printf("Total space: %d bytes %v\n", total, ByteCountIEC(total))
	glog.Printf("Used space: %d bytes %v\n\n", used, ByteCountIEC(used))
	glog.Printf("Free space: %d bytes %v\n\n", free, ByteCountIEC(free))
}

func GetAppSpace() (string, string, string) {
	dir := glog.AppHome()
	total, used, free, _ := util.GetDiskUsage(dir)
	//glog.Printf("Current Working Directory: %s %v\n", dir, err)
	//glog.Printf("Total space: %d bytes %v\n", total, ByteCountIEC(total))
	//glog.Printf("Used space: %d bytes %v\n\n", used, ByteCountIEC(used))
	//glog.Printf("Free space: %d bytes %v\n\n", free, ByteCountIEC(free))
	return ByteCountIEC(total), ByteCountIEC(used), ByteCountIEC(free)
}

func HasDiskSpace() bool {
	size := GetSelfSize()
	size *= 16
	dir := glog.AppHome()
	//total, used, free, err := util.GetDiskUsage(dir)
	_, _, free, _ := util.GetDiskUsage(dir)
	//glog.Printf("Current Working Directory: %s %v\n", dir, err)
	//glog.Printf("Total space: %d bytes %v\n", total, ByteCountIEC(total))
	//glog.Printf("Used space: %d bytes %v\n\n", used, ByteCountIEC(used))
	//glog.Printf("Free space: %d bytes %v\n\n", free, ByteCountIEC(free))
	if free > size {
		return true
	}
	return false
}
