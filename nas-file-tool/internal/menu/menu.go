package menu

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

// singleton 单例对象类型
type singleton struct {
	menuMap map[string]func()
}

var (
	instance *singleton
	once     sync.Once
)

// GetInstance 返回单例实例
func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{menuMap: make(map[string]func())} // 初始化逻辑
		fmt.Println("Singleton instance created")
	})
	return instance
}

func (s *singleton) AddMenuItem(name string, f func()) {
	s.menuMap[name] = f
}
func (s *singleton) ShowMenu() {
	defer func() {
		fmt.Print("按回车键退出程序...")
		endKey := make([]byte, 1)
		_, _ = os.Stdin.Read(endKey) // 等待用户输入任意内容后按回车
		os.Exit(0)
	}()
	s.AddMenuItem("退出程序", func() { os.Exit(0) })
	keys := make([]string, 0, len(s.menuMap))
	for k := range s.menuMap {
		keys = append(keys, k)
	}
	keys = append(keys)
	sort.Strings(keys)
	fmt.Println("=== 主菜单 ===")
	for i, key := range keys {
		fmt.Printf("%d. %s\n", i+1, key)
	}

	fmt.Print("请输入数字选择: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	choice, err := strconv.Atoi(input)
	if err != nil || choice < 1 || choice > len(keys) {
		fmt.Println("无效输入，请重试")
		return
	}

	selectedKey := keys[choice-1]
	s.menuMap[selectedKey]()
}
