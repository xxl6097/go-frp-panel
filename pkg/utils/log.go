package utils

import (
	"github.com/xxl6097/glog/glog"
	"net/http"
	"sync"
)

// LogQueue 定义日志队列
type LogQueue struct {
	mu sync.Mutex
	//messages []string
	messages *FixedQueue[string]
	clients  map[chan string]struct{}
}

// NewLogQueue 初始化日志队列
func NewLogQueue() *LogQueue {
	return &LogQueue{
		//messages: make([]string, 0),
		messages: NewFixedQueue[string](100),
		clients:  make(map[chan string]struct{}),
	}
}

// AddMessage 生产者添加日志消息
func (q *LogQueue) AddMessage(message string) {
	q.mu.Lock()
	defer q.mu.Unlock()
	//q.messages = append(q.messages, message)
	q.messages.Enqueue(message)
	// 通知所有客户端有新消息
	for client := range q.clients {
		client <- message
	}
}

// RegisterClient 消费者注册客户端
func (q *LogQueue) RegisterClient(client chan string) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.clients[client] = struct{}{}
}

// UnregisterClient 消费者注销客户端
func (q *LogQueue) UnregisterClient(client chan string) {
	q.mu.Lock()
	defer q.mu.Unlock()
	close(client)
	delete(q.clients, client)
	glog.Debug("UnregisterClient", client)
}

// SseHandler 处理函数
func SseHandler(queue *LogQueue) http.HandlerFunc {
	glog.Debug("SseHandler....1")
	return func(w http.ResponseWriter, r *http.Request) {
		glog.Debug("SseHandler....2")
		// 设置响应头
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
			return
		}

		// 为客户端创建一个消息通道
		client := make(chan string)
		queue.RegisterClient(client)
		defer queue.UnregisterClient(client)

		// 发送历史消息
		queue.mu.Lock()
		//for _, message := range queue.messages {
		//	w.Write([]byte("data: " + message + "\n\n"))
		//	flusher.Flush()
		//}
		for _, message := range queue.messages.Items() {
			w.Write([]byte("data: " + message + "\n\n"))
			flusher.Flush()
		}
		queue.mu.Unlock()

		// 监听新消息
		for {
			select {
			case message, ok := <-client:
				if !ok {
					//return
				}
				w.Write([]byte("data: " + message + "\n\n"))
				flusher.Flush()
				//case <-r.Context().Done():
				//	glog.Debug("客户端Done...")
				//	w.Write([]byte("data: 客户端断开\n\n"))
				//	flusher.Flush()
				//	return
			}
		}
	}
}
