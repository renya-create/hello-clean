package message

type Hello struct {
	message string
}

func New(msg string) *Hello {
	return &Hello{message: msg}
}

func (h *Hello) GetMessage() string {
	return h.message
}
