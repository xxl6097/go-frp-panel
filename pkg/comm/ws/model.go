package ws

// Message 发送JSON消息
type Message[T any] struct {
	Action string `json:"action"`
	DevMac string `json:"devMac"`
	DevIp  string `json:"devIp"`
	Data   T      `json:"data"`
}
