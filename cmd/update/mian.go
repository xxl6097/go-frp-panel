package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/xxl6097/glog/glog"
	"io"
	"net/http"
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

	var baseUrl = "https://api.github.com/repos/xxl6097/go-frp-panel/releases/latest"
	r, err := http.Get(baseUrl)
	if err != nil {
		glog.Fatal(err)
	}
	b, _ := io.ReadAll(r.Body)
	var res map[string]interface{}
	json.Unmarshal(b, &res)
	str := res["body"].(string)
	index := strings.Index(str, "---")

	fmt.Println(index, str)
	//codeBlocks := extractCodeBlocks(res["body"].(string))
	//for _, block := range codeBlocks {
	//	var r []string
	//	json.Unmarshal([]byte(block), &r)
	//	fmt.Println(r)
	//}
}
