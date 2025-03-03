package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	rawURL := "C://test.txt"

	// 提取路径部分并获取文件名
	fileName := filepath.Base(rawURL)
	fmt.Println("文件名:", fileName) // 输出: document.pdf
}
