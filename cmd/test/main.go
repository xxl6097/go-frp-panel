package main

import (
	"fmt"
	"path"
)

func main() {
	rawURL := "C://test.txt"

	// 提取路径部分并获取文件名
	fileName := path.Base(rawURL)
	fmt.Println("文件名:", fileName) // 输出: document.pdf
}
