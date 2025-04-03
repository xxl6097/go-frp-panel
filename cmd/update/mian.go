package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/xxl6097/go-frp-panel/pkg/utils"
	utils2 "github.com/xxl6097/go-service/gservice/utils"
	"strings"
)

func extractCodeBlocks(markdown string) []string {
	var codeBlocks []string
	inCodeBlock := false
	var currentCodeBlock strings.Builder

	scanner := bufio.NewScanner(strings.NewReader(markdown))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "```") {
			if inCodeBlock {
				codeBlocks = append(codeBlocks, currentCodeBlock.String())
				currentCodeBlock.Reset()
			}
			inCodeBlock = !inCodeBlock
		} else if inCodeBlock {
			currentCodeBlock.WriteString(line)
			currentCodeBlock.WriteRune('\n')
		}
	}

	return codeBlocks
}
func main() {

	//var baseUrl = "https://api.github.com/repos/xxl6097/go-frp-panel/releases/latest"
	//r, err := http.Get(baseUrl)
	//if err != nil {
	//	glog.Fatal(err)
	//}
	//b, _ := io.ReadAll(r.Body)
	//var res map[string]interface{}
	//json.Unmarshal(b, &res)
	//str := res["body"].(string)
	//index := strings.Index(str, "---")
	//fmt.Println(index, str[:index])

	//codeBlocks := extractCodeBlocks(res["body"].(string))
	//for _, block := range codeBlocks {
	//	var r []string
	//	json.Unmarshal([]byte(block), &r)
	//	fmt.Println(r)
	//}

	newProxy := []string{"https://ghfast.top/https://github.com/xxl6097/go-frp-panel/releases/download/v0.1.60/acfrps_v0.1.60_linux_amd64",
		"https://gh-proxy.com/https://github.com/xxl6097/go-frp-panel/releases/download/v0.1.60/acfrps_v0.1.60_linux_amd64",
		"https://ghproxy.1888866.xyz/https://github.com/xxl6097/go-frp-panel/releases/download/v0.1.60/acfrps_v0.1.60_linux_amd64",
		"https://github.com/xxl6097/go-frp-panel/releases/download/v0.1.60/acfrps_v0.1.60_linux_amd64",
	}
	ctx, cancel := context.WithCancel(context.Background())
	newUrl := utils.DynamicSelect[string](newProxy, func(i int, s string) string {
		var dst string
		for {
			fmt.Println("通道 ", i, s)
			dstFilePath, err := utils2.DownloadFileWithCancel(ctx, s)
			if err == nil {
				dst = dstFilePath
				break
			}
		}
		return dst
	})
	cancel()
	fmt.Println("下载完成", newUrl)

	//testFilePath := filepath.Join(os.TempDir(), "go-frp-panel-update-test.txt")
	//tempFolder := fmt.Sprintf("%d", time.Now().Unix())
	//dir, f := filepath.Split(testFilePath)
	//testFilePath = filepath.Join(dir, tempFolder, f)
	//fmt.Println(testFilePath)
}
