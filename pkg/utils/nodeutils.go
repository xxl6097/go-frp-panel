package utils

import (
	"fmt"
	"github.com/xxl6097/go-service/gservice/utils"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type Node struct {
	Label    string `json:"label"`
	Value    string `json:"value"`
	FilePath string `json:"filePath"`
}
type Option struct {
	Label    string `json:"label"`
	Value    string `json:"value"`
	Children []Node `json:"children"`
}

func SplitLastTwoByUnderscore(s string) []string {
	// 过滤空元素
	parts := strings.FieldsFunc(s, func(r rune) bool {
		return r == '_'
	})
	if len(parts) < 2 {
		return []string{}
	}
	return parts[len(parts)-2:]
}

func CleanExt(name string) string {
	filename := path.Base(name) // 获取文件名"app.log"
	nameOnly := filename[:len(filename)-len(path.Ext(filename))]
	return nameOnly
}

func ToTree(dir string, entries []string) []Option {
	maps := make(map[string][]Node)
	for _, uri := range entries {
		var name string
		var fileUri string
		if utils.IsURL(uri) {
			name = path.Base(uri) // 输出：report.pdf
			fileUri = uri
		} else {
			name = uri
			fileUri = filepath.Join(dir, name)
		}
		result := SplitLastTwoByUnderscore(name)
		fmt.Printf("%-30s => %v\n", name, result)
		if len(result) == 2 {
			nodeArray := maps[result[0]]
			if nodeArray == nil {
				nodeArray = make([]Node, 0)
			}
			nameOnly := CleanExt(result[1])

			nodeArray = append(nodeArray, Node{
				Label:    nameOnly,
				Value:    nameOnly,
				FilePath: fileUri,
			})
			maps[result[0]] = nodeArray
		}
	}

	var options []Option
	for k, v := range maps {
		options = append(options, Option{
			Label:    ToUpperFirst(k),
			Value:    k,
			Children: v,
		})
	}
	return options
}

func GetNodes(dir string) []Option {
	files := make([]string, 0)
	entries, _ := os.ReadDir(dir) // 读取当前目录
	for _, entry := range entries {
		name := entry.Name()
		files = append(files, name)
	}
	return ToTree(dir, files)
}

func GetNodes1(dir string) []Option {
	maps := make(map[string][]Node)
	entries, _ := os.ReadDir(dir) // 读取当前目录
	for _, entry := range entries {
		name := entry.Name()
		result := SplitLastTwoByUnderscore(name)
		fmt.Printf("%-30s => %v\n", name, result)
		if len(result) == 2 {
			nodeArray := maps[result[0]]
			if nodeArray == nil {
				nodeArray = make([]Node, 0)
			}
			nameOnly := CleanExt(result[1])

			nodeArray = append(nodeArray, Node{
				Label:    nameOnly,
				Value:    nameOnly,
				FilePath: filepath.Join(dir, name),
			})
			maps[result[0]] = nodeArray
		}
	}
	var options []Option
	for k, v := range maps {
		options = append(options, Option{
			Label:    ToUpperFirst(k),
			Value:    k,
			Children: v,
		})
	}
	return options
}
