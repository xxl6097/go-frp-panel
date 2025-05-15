package sse

import (
	"encoding/json"
	"fmt"
	"github.com/xxl6097/glog/glog"
	"github.com/xxl6097/go-frp-panel/pkg/comm/iface"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"
)

// SSEServer 管理所有客户端连接和事件广播
type SSEServer struct {
	clients      map[string]*iface.SSEClient
	register     chan *iface.SSEClient
	unregister   chan *iface.SSEClient
	broadcast    chan iface.SSEEvent
	mu           sync.RWMutex
	clientIDFunc func(r *http.Request) string
	callback     iface.OnSSECallBack
}

// NewServer 创建一个新的SSE服务器
func NewServer() iface.ISSE {
	return &SSEServer{
		clients:      make(map[string]*iface.SSEClient),
		register:     make(chan *iface.SSEClient),
		unregister:   make(chan *iface.SSEClient),
		broadcast:    make(chan iface.SSEEvent),
		clientIDFunc: defaultClientIDFunc,
	}
}

func (s *SSEServer) SetSSECallBack(back iface.OnSSECallBack) {
	s.callback = back
}

// defaultClientIDFunc 默认的客户端ID生成函数
func defaultClientIDFunc(r *http.Request) string {
	level := r.URL.Query().Get("type")
	glog.Infof("sse新连接 %s %v", r.URL.Path, level)
	if level != "" {
		return fmt.Sprintf("%s-%s-%d", level, r.RemoteAddr, time.Now().UnixNano())
	}
	return fmt.Sprintf("%s-%d", r.RemoteAddr, time.Now().UnixNano())
}

// SetClientIDFunc 设置自定义的客户端ID生成函数
func (s *SSEServer) SetClientIDFunc(f func(r *http.Request) string) {
	s.clientIDFunc = f
}

// ServeHTTP 实现HTTP处理接口
func (s *SSEServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 设置SSE响应头
	headers := w.Header()
	headers.Set("Content-Type", "text/event-stream")
	headers.Set("Cache-Control", "no-cache")
	headers.Set("Connection", "keep-alive")
	headers.Set("Access-Control-Allow-Origin", "*")

	// 处理客户端断开连接
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	frpId := r.URL.Query().Get("frpId")
	secKey := r.URL.Query().Get("secKey")

	// 创建客户端
	client := &iface.SSEClient{
		SseId:     s.clientIDFunc(r),
		FrpID:     frpId,
		SecKey:    secKey,
		Send:      make(chan iface.SSEEvent, 256),
		CloseConn: make(chan struct{}),
	}

	// 注册客户端
	s.register <- client

	// 确保客户端断开时注销
	defer func() {
		s.unregister <- client
	}()

	// 发送心跳包保持连接
	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	// 发送初始连接确认
	s.sendEvent(w, iface.SSEEvent{Event: "connected", Payload: map[string]string{"client_id": client.SseId}})
	flusher.Flush()

	// 监听客户端发送通道和关闭信号
	for {
		select {
		case event, ok := <-client.Send:
			if !ok {
				return
			}
			s.sendEvent(w, event)
			flusher.Flush()
		case <-ticker.C:
			// 发送心跳包
			fmt.Fprintf(w, ": heartbeat\n\n")
			flusher.Flush()
		case <-client.CloseConn:
			return
		case <-r.Context().Done():
			return
		}
	}
}

// 发送事件到客户端
func (s *SSEServer) sendEvent(w http.ResponseWriter, event iface.SSEEvent) {
	data, err := json.Marshal(event)
	if err != nil {
		fmt.Fprintf(w, "data: %s\n\n", err.Error())
		return
	}
	fmt.Fprintf(w, "data: %s\n\n", data)
}

// Send 发送事件到客户端
func (s *SSEServer) Send(client *iface.SSEClient, event iface.SSEEvent) error {
	//if w == nil {
	//	return fmt.Errorf("ResponseWriter is nil")
	//}
	//// 处理客户端断开连接
	//flusher, ok := w.(http.Flusher)
	//if !ok {
	//	return fmt.Errorf("ResponseWriter to Flusher not supported")
	//}
	//if flusher == nil {
	//	return fmt.Errorf("flusher is nil")
	//}
	//defer flusher.Flush()
	//data, err := json.Marshal(event)
	//if err != nil {
	//	fmt.Fprintf(w, "data: %s\n\n", err.Error())
	//	return err
	//}
	//fmt.Fprintf(w, "data: %s\n\n", data)
	//return nil

	if client == nil {
		return fmt.Errorf("client is nil")
	}
	select {
	case client.Send <- event:
		return nil
	default:
		close(client.Send)
		delete(s.clients, client.SseId)
		return fmt.Errorf("send fail")
	}
}

// Start 启动服务器循环
func (s *SSEServer) Start() {
	go s.run()
}

// run 服务器主循环
func (s *SSEServer) run() {
	for {
		select {
		case client := <-s.register:
			s.mu.Lock()
			s.clients[client.SseId] = client
			glog.Debugf("register:%s,%p", client.SseId, s.callback)
			if s.callback != nil {
				s.callback.OnSseNewConnection(client)
			}
			s.mu.Unlock()
		case client := <-s.unregister:
			s.mu.Lock()
			if _, ok := s.clients[client.SseId]; ok {
				glog.Error("unregister client id:", client.SseId)
				if s.callback != nil {
					s.callback.OnSseDisconnect(client)
				}
				close(client.Send)
				delete(s.clients, client.SseId)
			}
			s.mu.Unlock()
		case event := <-s.broadcast:
			s.mu.RLock()
			for id, client := range s.clients {
				select {
				case client.Send <- event:
				default:
					close(client.Send)
					delete(s.clients, id)
				}
			}
			s.mu.RUnlock()
		}
	}
}

// Broadcast 向所有客户端广播事件
func (s *SSEServer) Broadcast(event iface.SSEEvent) {
	s.broadcast <- event
}

// BroadcastTo 向特定客户端广播事件
func (s *SSEServer) BroadcastTo(sseId string, event iface.SSEEvent) bool {
	//s.mu.RLock()
	//defer s.mu.RUnlock()

	client, ok := s.clients[sseId]
	if !ok {
		return false
	}

	select {
	case client.Send <- event:
		return true
	default:
		close(client.Send)
		delete(s.clients, sseId)
		return false
	}
}

// BroadcastByType 向特定客户端广播事件
func (s *SSEServer) BroadcastByType(typeName string, event iface.SSEEvent) bool {
	//s.mu.RLock()
	//defer s.mu.RUnlock()

	for k, client := range s.clients {
		if strings.HasPrefix(k, typeName) {
			select {
			case client.Send <- event:
				return true
			default:
				close(client.Send)
				delete(s.clients, k)
				return false
			}
		}
	}
	return true
}

// CloseClient 关闭特定客户端连接
func (s *SSEServer) CloseClient(clientID string) bool {
	//s.mu.Lock()
	//defer s.mu.Unlock()

	client, ok := s.clients[clientID]
	if !ok {
		return false
	}

	close(client.CloseConn)
	delete(s.clients, clientID)
	return true
}

// GetClientCount 获取当前连接的客户端数量
func (s *SSEServer) GetClientCount() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.clients)
}

// GetClientIDs 获取所有客户端ID
func (s *SSEServer) GetClientIDs() []string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	ids := make([]string, 0, len(s.clients))
	for id := range s.clients {
		ids = append(ids, id)
	}
	return ids
}

func test() {
	// 创建SSE服务器
	server := NewServer()
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
		glog.Println("SSEServer started on :8080")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			glog.Fatalf("SSEServer error: %v", err)
		}
	}()

	// 优雅关闭
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	glog.Println("Shutting down server...")
	// 关闭所有客户端连接
	for _, clientID := range server.GetClientIDs() {
		server.CloseClient(clientID)
	}
	glog.Println("SSEServer gracefully shutdown")
}
