package presenter_impl

import (
	"fmt"
	"hello-clean-architecture/internal/domain/message"
	"hello-clean-architecture/internal/usecase/presenter"
)

type hello struct{}

func New() presenter.Hello {
	return &hello{}
}

func (hp *hello) FormatHelloMessage(msg *message.Hello) string {
	// 必要であれば、ここでJSON形式に変換したり、HTMLタグを追加可能
	return fmt.Sprintf("%s\n", msg.GetMessage())
}
