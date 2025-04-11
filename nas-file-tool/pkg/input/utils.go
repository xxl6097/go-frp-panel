package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func tips(title string) {
	str := strings.ReplaceAll(title, "请输入", "")
	str = strings.ReplaceAll(str, "please input", "")
	str = strings.ReplaceAll(str, "：", "")
	str = strings.ReplaceAll(str, ":", "")
	str = fmt.Sprintf("【%s】不允许输入空", str)
	fmt.Println(str)
}
func InputStringEmpty(title, defaultString string) string {
	reader := bufio.NewReader(os.Stdin)
	//glog.Print(title)
	fmt.Print(title)
	input, err := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if err != nil {
		return InputString(title)
	}
	if input == "" {
		return defaultString
	}
	//return strings.TrimSpace(input)
	return input
}

func InputString(title string) string {
	reader := bufio.NewReader(os.Stdin)
	//glog.Print(title)
	fmt.Print(title)
	input, err := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if err != nil {
		return InputString(title)
	}
	//return strings.TrimSpace(input)
	if len(input) == 0 {
		tips(title)
		return InputString(title)
	}
	return input
}
func InputInt(title string) int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(title)
	input, err := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if err != nil {
		return InputInt(title)
	}
	if len(input) == 0 {
		tips(title)
		return InputInt(title)
	}
	num, err := strconv.Atoi(input)
	if err != nil {
		return InputInt(title)
	}
	return num
}
