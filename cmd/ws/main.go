package main

import (
	"fmt"
	"github.com/xxl6097/go-frp-panel/pkg/comm/ws"
	"github.com/xxl6097/go-frp-panel/pkg/utils"
)

func main() {
	ws.GetClientInstance().Init("ws://uuxia.cn:6500/frp", "admin", "het002402")

	// 示例：创建容量为3的队列
	queue := utils.NewFixedQueue[int](3)

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	fmt.Println("当前队列:", queue.Items()) // 输出: [1 2 3]

	queue.Enqueue(4)                       // 替换最早的元素1
	fmt.Println("入队4后:", queue.Items()) // 输出: [2 3 4]

	queue.Dequeue()
	fmt.Println("出队后:", queue.Items()) // 输出: [3 4]
	select {}
}
