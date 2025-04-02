package utils

import (
	"fmt"
	"github.com/xxl6097/glog/glog"
	"github.com/xxl6097/go-service/gservice/utils"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
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

func DownLoad(url string, args ...string) (string, error) {
	// 要下载的文件的 URL
	// 发送 HTTP GET 请求
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("请求失败，状态码: %d", resp.StatusCode)
	}

	var dstFile string
	if args != nil && len(args) > 0 {
		dstFile = args[0]
	}
	if dstFile == "" {
		dstName := GetFileNameFromUrl(url)
		if dstName == "" {
			dstName = GetFilenameFromHeader(resp.Header)
		}
		if dstName == "" {
			fileName := time.Now().Unix()
			dstName = fmt.Sprintf("%d", fileName)
		}
		if dstName != "" {
			dstFile = filepath.Join(utils.GetUpgradeDir(), dstName)
		}
	}
	glog.Debug("download...", url, dstFile)

	// 获取文件的总大小
	totalSize := resp.ContentLength
	if totalSize == -1 {
		fmt.Println("无法获取文件大小，可能不支持 Content-Length 头信息。")
		return "", fmt.Errorf("无法获取文件大小，可能不支持 Content-Length 头信息。")
	}
	sizeA := float64(resp.ContentLength) / 1024 / 1024
	fmt.Printf("文件大小:%.2fM\n", sizeA)
	// 创建一个本地文件用于保存下载的内容
	file, err := os.Create(dstFile)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// 创建进度写入器实例
	pw := &ProgressWriter{TotalSize: totalSize, Progress: -1, Title: "文件下载："}
	// 将响应体的数据复制到本地文件，并通过 ProgressWriter 跟踪进度
	_, err = io.Copy(io.MultiWriter(file, pw), resp.Body)
	if err != nil {
		return "", fmt.Errorf("下载出错: %v", err)
	}

	fmt.Println("下载完成")
	return dstFile, nil
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
