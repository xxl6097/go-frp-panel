package ws

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/xxl6097/glog/glog"
	"github.com/xxl6097/go-frp-panel/pkg"
	"github.com/xxl6097/go-frp-panel/pkg/utils"
	"math"
	"net/http"
	"os"
	"runtime"
	"sync"
	"time"
)

// Websocketclient WebSocket客户端结构体
type Websocketclient struct {
	conn           *websocket.Conn // WebSocket连接
	header         *http.Header
	url            string                                // 服务器地址
	reconnectDelay time.Duration                         // 重连延迟
	maxReconnects  int                                   // 最大重连次数
	isConnected    bool                                  // 连接状态
	openHandler    func(*websocket.Conn, *http.Response) // 消息处理函数
	messageHandler func([]byte)                          // 消息处理函数
	errorHandler   func(error)                           // 错误处理函数
	closeHandler   func(int, string)                     // 关闭处理函数
}

// NewWebSocketClient 创建WebSocket客户端实例
func NewWebSocketClient(url string, header *http.Header) *Websocketclient {
	return &Websocketclient{
		url:            url,
		reconnectDelay: 5 * time.Second,
		maxReconnects:  10,
		header:         header,
	}
}

// SetReconnectConfig 设置重连配置
func (c *Websocketclient) SetReconnectConfig(delay time.Duration, maxReconnects int) {
	c.reconnectDelay = delay
	c.maxReconnects = maxReconnects
}

func (c *Websocketclient) SetOpenHandler(handler func(*websocket.Conn, *http.Response)) {
	c.openHandler = handler
}

// SetMessageHandler 设置消息处理回调
func (c *Websocketclient) SetMessageHandler(handler func([]byte)) {
	c.messageHandler = handler
}

// SetErrorHandler 设置错误处理回调
func (c *Websocketclient) SetErrorHandler(handler func(error)) {
	c.errorHandler = handler
}

// SetCloseHandler 设置关闭处理回调
func (c *Websocketclient) SetCloseHandler(handler func(int, string)) {
	c.closeHandler = handler
}

// Connect 连接到WebSocket服务器
func (c *Websocketclient) Connect() error {
	var reconnects int

	for {
		conn, resp, err := websocket.DefaultDialer.Dial(c.url, *c.header)
		c.conn = conn
		if err == nil {
			if c.openHandler != nil {
				c.openHandler(conn, resp)
			}
			c.isConnected = true
			glog.Printf("WebSocket连接成功: %s", c.url)
			go c.readMessages()
			return nil
		}

		reconnects++
		if c.maxReconnects > 0 && reconnects > c.maxReconnects {
			err = fmt.Errorf("连接失败，已达到最大重连次数: %w", err)
			return err
		}

		glog.Printf("%s 连接失败，尝试重连 %v %v", c.url, err, resp.Status)
		time.Sleep(c.reconnectDelay)
	}
}

// Close 关闭WebSocket连接
func (c *Websocketclient) Close() error {
	if !c.isConnected || c.conn == nil {
		return nil
	}

	c.isConnected = false
	return c.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
}

// SendText 发送文本消息
func (c *Websocketclient) SendText(message string) error {
	return c.sendMessage(websocket.TextMessage, []byte(message))
}

// SendJSON 发送JSON消息
func (c *Websocketclient) SendJSON(data interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return c.sendMessage(websocket.TextMessage, jsonData)
}

// sendMessage 发送消息的底层实现
func (c *Websocketclient) sendMessage(messageType int, data []byte) error {
	if !c.isConnected || c.conn == nil {
		return fmt.Errorf("WebSocket未连接")
	}

	return c.conn.WriteMessage(messageType, data)
}

// readMessages 读取服务器消息的协程
func (c *Websocketclient) readMessages() {
	defer func() {
		c.isConnected = false
		if c.conn != nil {
			c.conn.Close()
			c.conn = nil
		}
	}()

	for {
		if !c.isConnected {
			return
		}

		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if c.errorHandler != nil {
				c.errorHandler(err)
			}
			glog.Printf("WebSocket断开: %v", err)
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				glog.Printf("WebSocket意外关闭: %v", err)
				//go c.reconnect()
			}
			go c.reconnect()
			return
		}

		if c.messageHandler != nil {
			c.messageHandler(message)
		}
	}
}

// reconnect 尝试重新连接
func (c *Websocketclient) reconnect() {
	if c.isConnected {
		return
	}

	glog.Println("开始重新连接...")
	// 避免阻塞其他操作
	go func() {
		err := c.Connect()
		if err != nil && c.errorHandler != nil {
			c.errorHandler(fmt.Errorf("重连失败: %w", err))
		}
	}()
}

type Client struct {
	cls     *Websocketclient
	clients map[string]*Websocketclient
}

var (
	instance *Client
	once     sync.Once
)

// GetClientInstance 返回单例实例
func GetClientInstance() *Client {
	once.Do(func() {
		instance = &Client{
			clients: make(map[string]*Websocketclient),
		}
		glog.Println("WebSocket Singleton client instance created")
	})
	return instance
}

func (c *Client) NewClient(id, serverAddress, authorization string) {
	var err error
	defer func() {
		if err != nil {
			glog.Errorf("websocket连接失败: %v", err)
		}
	}()
	if id == "" {
		err = fmt.Errorf("id is nil")
		return
	}
	if serverAddress == "" {
		err = fmt.Errorf("serverAddress is nil")
		return
	}
	if authorization == "" {
		err = fmt.Errorf("authorization is nil")
		return
	}
	baseUrl := fmt.Sprintf("ws://%s/frp", serverAddress)
	header := c.header(id, authorization)
	cls := NewWebSocketClient(baseUrl, header)
	// 设置消息处理函数
	cls.SetMessageHandler(func(message []byte) {
		glog.Printf("收到消息: %s", string(message))
	})
	// 设置错误处理函数
	cls.SetOpenHandler(func(conn *websocket.Conn, response *http.Response) {
		glog.Errorf("连接成功: %v,%v,Status:%v", conn.LocalAddr(), conn.RemoteAddr(), response.Status)
	})

	// 设置错误处理函数
	cls.SetErrorHandler(func(err error) {
		glog.Errorf("发生错误: %v", err)
	})

	// 设置关闭处理函数
	cls.SetCloseHandler(func(code int, text string) {
		glog.Errorf("连接关闭: %d %s", code, text)
	})

	// 设置重连配置
	cls.SetReconnectConfig(5*time.Second, math.MaxInt)

	c.clients[serverAddress] = cls
	go func() {
		// 连接到服务器
		glog.Warnf("websocket连接 baseurl:%s,header:%+v", baseUrl, header)
		if err := cls.Connect(); err != nil {
			glog.Errorf("连接失败: %v", err)
		}
	}()
}

func (c *Client) header(id, authorization string) *http.Header {
	header := &http.Header{}
	header.Set("Authorization", "Basic "+authorization)
	devInfo, err := utils.GetDeviceInfo()
	if err == nil {
		wsid := uuid.New().String() // 生成版本4的随机UUID
		hostname, e := os.Hostname()
		if e == nil {
			header.Set("DevName", hostname)
			glog.Debug("DevName", hostname)
		}
		header.Set("OsType", fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH))
		header.Set("LocalMacAddress", devInfo.MacAddress)
		header.Set("AppVersion", pkg.AppVersion)
		header.Set("LocalIpv4", devInfo.Ipv4)
		header.Set("InterfaceName", devInfo.Name)
		header.Set("DisplayName", devInfo.DisplayName)
		header.Set("FrpID", id)
		header.Set("WebSocketID", wsid)
	} else {
		glog.Error("获取设备信息失败", err)
	}
	return header
}

// SendText 发送文本消息
//func (c *client) SendText(message string) error {
//	return c.cls.sendMessage(websocket.TextMessage, []byte(message))
//}

func (c *Client) SendJSON1(data interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return c.cls.sendMessage(websocket.TextMessage, jsonData)
}
func (c *Client) SendJSON(data interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	for k, v := range c.clients {
		if v != nil {
			err = v.sendMessage(websocket.TextMessage, jsonData)
			glog.Debugf("sendMessage %s %v", k, err)
		}
	}
	return err
}
func (c *Client) SetOpenHandler(handler func(*websocket.Conn, *http.Response)) {
	//if c.cls != nil && handler != nil {
	//	c.cls.SetOpenHandler(handler)
	//}
	if handler == nil {
		return
	}
	for _, v := range c.clients {
		if v != nil {
			v.SetOpenHandler(handler)
		}
	}
}
func (c *Client) SetMessageHandler(handler func([]byte)) {
	if handler == nil {
		return
	}
	//if c.cls != nil && handler != nil {
	//	c.cls.SetMessageHandler(handler)
	//}

	for _, v := range c.clients {
		if v != nil {
			v.SetMessageHandler(handler)
		}
	}
}

//func (c *Client) GetClient() *Websocketclient {
//	return c.cls
//}

func Connect(baseUrl, user, pass string) {
	header := http.Header{}
	header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(glog.Sprintf("%s:%s", user, pass))))
	conn, resp, err := websocket.DefaultDialer.DialContext(context.Background(), baseUrl, header)
	if err != nil {
		glog.Fatal("连接失败:", err, resp)
	}
	glog.Debug(conn, resp, err)
	defer conn.Close()
}
