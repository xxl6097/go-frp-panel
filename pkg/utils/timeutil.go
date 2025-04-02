package utils

import (
	"time"
)

func GetTime() string {
	// 加载时区信息，这里以 Asia/Shanghai 为例
	//loc, err := time.LoadLocation("Asia/Shanghai")
	//if err != nil {
	//	glog.Error("加载时区信息出错:", err)
	//	return ""
	//}
	// 获取当前时间并设置时区
	//now := time.Now().In(loc)
	now := time.Now()
	// 定义格式化模板，包含毫秒
	format := "2006-01-02 15:04:05.000"
	// 格式化时间
	formattedTime := now.Format(format)
	return formattedTime
}

func GetFileNameByTime() string {
	// 加载时区信息，这里以 Asia/Shanghai 为例
	//loc, err := time.LoadLocation("Asia/Shanghai")
	//if err != nil {
	//	glog.Error("加载时区信息出错:", err)
	//	return ""
	//}
	// 获取当前时间并设置时区
	now := time.Now()
	// 定义格式化模板，包含毫秒
	format := "20060102150405"
	// 格式化时间
	formattedTime := now.Format(format)
	return formattedTime
}
