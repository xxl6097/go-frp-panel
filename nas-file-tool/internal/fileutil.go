package internal

import (
	"fmt"
	"nas-file-tool/pkg/input"
	"nas-file-tool/pkg/utils"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func scanFiles(root string, re *regexp.Regexp) ([]string, error) {
	var matches []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && re.MatchString(info.Name()) {
			matches = append(matches, path)
		}
		return nil
	})
	return matches, err
}

func filesMove() {
	srcInput := input.InputString("复制目录下的文件(文件名可用通配符)：")
	dstInput := input.InputString("复制目的位置：")
	director := filepath.Dir(srcInput)
	matches, _ := filepath.Glob(srcInput)
	for _, path := range matches {
		fileName := filepath.Base(path)
		srcPath := filepath.Join(director, fileName)
		go func(srcPath, dstInput string) {
			err := utils.CopyFileToDir(srcPath, dstInput)
			if err != nil {
				fmt.Println("复制失败", srcPath, err)
			}
		}(srcPath, dstInput)
	}
}

func copyFilesWithAllChildren() {
	//root := "/Users/uuxia/Desktop/work/code/github/golang/go-frp-panel"
	//pattern := "*.go" // 支持通配符如 *.txt 或 logs/​**​/*.log

	srcInput := input.InputString("请输入路径(文件名可用通配符)：")
	root, pattern := filepath.Split(srcInput)

	// 将通配符转换为正则表达式（支持 * 和 ​**​）
	regexPattern := strings.ReplaceAll(pattern, ".", `\.`)
	regexPattern = strings.ReplaceAll(regexPattern, "*", ".*")
	re := regexp.MustCompile("^" + regexPattern + "$")

	var matches []string
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}
		// 提取相对路径用于匹配（如 data/logs/error.log）
		relPath, _ := filepath.Rel(root, path)
		if re.MatchString(relPath) {
			matches = append(matches, path)
		}
		return nil
	})

	fmt.Println("匹配文件列表:")
	for _, f := range matches {
		fmt.Println("-", f)
	}
}
