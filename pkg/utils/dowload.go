package utils

import (
	"fmt"
	"github.com/xxl6097/glog/glog"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
)

// ProgressWriter 自定义进度写入器结构体
type ProgressWriter struct {
	TotalSize int64
	Written   int64
	Progress  float64
	Title     string
}

// Write 实现 io.Writer 接口的 Write 方法
func (pw *ProgressWriter) Write(p []byte) (int, error) {
	n := len(p)
	pw.Written += int64(n)
	// 计算下载进度百分比
	progress := float64(pw.Written) / float64(pw.TotalSize) * 100
	// 使用 \r 覆盖当前行，实现进度动态更新
	if progress >= pw.Progress {
		glog.Printf("%s %.2f%%\n", pw.Title, progress)
		pw.Progress = progress
		pw.Progress += 5
	}
	return n, nil
}

func GetFileNameFromUrl(rawURL string) string {
	parsedURL, _ := url.Parse(rawURL)

	// 提取路径部分并获取文件名
	fileName := path.Base(parsedURL.Path)
	fmt.Println("文件名:", fileName) // 输出: document.pdf
	return fileName
}

func GetFilenameFromHeader(header http.Header) string {
	contentDisposition := header.Get("Content-Disposition")
	parts := strings.Split(contentDisposition, ";")
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if strings.HasPrefix(part, "filename=") {
			fileName := strings.TrimPrefix(part, "filename=")
			fileName = strings.Trim(fileName, `"`) // 去除双引号
			return fileName
		}
	}
	return ""
}

func SaveFile(file multipart.File, fileSize int64, saveFilePath string) error {
	dst, err := os.Create(saveFilePath)
	if err != nil {
		return fmt.Errorf("create file: %v", err)
	}
	defer dst.Close()
	pw := &ProgressWriter{TotalSize: fileSize, Progress: -1, Title: "文件保存："}
	_, err = io.Copy(io.MultiWriter(dst, pw), file)
	if err != nil {
		return fmt.Errorf("write file: %v", err)
	}
	return nil
}
