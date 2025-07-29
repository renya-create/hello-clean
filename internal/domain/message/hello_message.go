package message

type HelloMessage struct {
	message string
}

func New(msg string) *HelloMessage {
	return &HelloMessage{message: msg}
}

func (h *HelloMessage) GetMessage() string {
	return h.message
}
