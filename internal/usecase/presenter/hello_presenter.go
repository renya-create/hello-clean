package presenter

import "hello-clean/internal/domain/model"

// "Hello, World!"メッセージのプレゼンテーションを抽象化するインターフェース
type HelloPresenter interface {
	// FormatHelloMessage はHelloMessageを外部に提示する形式に整形
	FormatHelloMessage(msg *model.HelloMessage) string
}
