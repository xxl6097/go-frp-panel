package main

import (
	"fmt"
	"github.com/xxl6097/go-frp-panel/pkg/comm/iface"
	"github.com/xxl6097/go-frp-panel/pkg/comm/sse"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 创建SSE服务器
	server := sse.NewServer()
	server.Start()

	// 注册SSE处理器
	http.Handle("/events", server)

	// 模拟数据推送处理器
	http.HandleFunc("/push", func(w http.ResponseWriter, r *http.Request) {
		message := r.URL.Query().Get("message")
		if message == "" {
			message = "Hello, World!"
		}

		// 广播消息给所有客户端
		server.Broadcast(iface.SSEEvent{
			Event:   "message",
			Payload: map[string]string{"message": message, "timestamp": time.Now().String()},
		})

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Message broadcast to %d clients\n", server.GetClientCount())
	})

	// 启动HTTP服务器
	go func() {
		log.Println("SSEServer started on :8080")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatalf("SSEServer error: %v", err)
		}
	}()

	// 优雅关闭
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	log.Println("Shutting down server...")
	// 关闭所有客户端连接
	for _, clientID := range server.GetClientIDs() {
		server.CloseClient(clientID)
	}
	log.Println("SSEServer gracefully shutdown")
}
