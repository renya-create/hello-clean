package model

type HelloMessage struct {
	Message string
}

// コンストラクタ
func NewHelloMessage(msg string) *HelloMessage {
	return &HelloMessage{Message: msg}
}

// getter
func (hm *HelloMessage) GetMessage() string {
	return hm.Message
}
