package presenter

import "hello-clean-architecture/internal/domain/message"

// "Hello, World!"メッセージのプレゼンテーションを抽象化するインターフェース
type Hello interface {
	// FormatHelloMessage はHelloMessageを外部に提示する形式に整形
	FormatHelloMessage(msg *message.HelloMessage) string
}
