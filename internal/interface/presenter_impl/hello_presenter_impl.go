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

// helloPresenter構造体がpresenter.HelloPresenterインターフェースを「実装」するために必要なメソッド
func (hp *hello) FormatHelloMessage(msg *message.HelloMessage) string {
	// 必要であれば、ここでJSON形式に変換したり、HTMLタグを追加可能
	return fmt.Sprintf("%s\n", msg.GetMessage())
}
