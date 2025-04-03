package main

import (
	"fmt"
	"github.com/xxl6097/go-service/gservice/gore/util"
	"os"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current working directory: %v\n", err)
		return
	}
	total, used, free, err := util.GetDiskUsage(path)
	fmt.Printf("Current Working Directory: %s\n", path)
	if err != nil {
		fmt.Printf("Error getting disk usage: %v\n", err)
		return
	}
	fmt.Printf("Total space: %d bytes %v\n", total, float64(total)/1024/1024/1024)
	fmt.Printf("Used space: %d bytes %v\n\n", used, float64(used)/1024/1024/1024)
	fmt.Printf("Free space: %d bytes %v\n\n", free, float64(free)/1024/1024/1024)
}
