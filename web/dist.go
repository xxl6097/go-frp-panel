package web

import (
	"embed"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

//go:embed dist/*
var files embed.FS

func UnPack() {
	// 目标目录
	destDir := "./clients"
	// 创建目标目录
	if err := os.MkdirAll(destDir, os.ModePerm); err != nil {
		log.Fatalf("创建目录失败: %v", err)
	}
	// 遍历嵌入的目录
	entries, err := files.ReadDir("dist")
	if err != nil {
		fmt.Printf("读取目录时出错: %v\n", err)
		return
	}
	for _, entry := range entries {
		if entry.IsDir() {
			fmt.Printf("目录: %s\n", entry.Name())
		} else {
			fmt.Printf("文件: %s\n", entry.Name())
			fileName := entry.Name()
			srcData, err := files.ReadFile(filepath.Join("dist", entry.Name()))
			if err != nil {
				log.Fatalf("读取文件 %s 失败: %v", fileName, err)
			}
			// 构建目标路径
			destPath := filepath.Join(destDir, fileName)
			// 写入文件（保留元数据）
			if err := os.WriteFile(destPath, srcData, 0755); err != nil {
				log.Fatalf("写入文件 %s 失败: %v", destPath, err)
			}
			fmt.Printf("已复制: %s → %s\n", fileName, destPath)
		}
	}
}
