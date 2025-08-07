package presenter

import "hello-clean-architecture/internal/domain/message"

// "Hello, World!"メッセージのプレゼンテーションを抽象化するインターフェース
type Hello interface {
	// FormatHelloMessage は message.Hello を外部に提示する形式に整形
	FormatHelloMessage(msg *message.Hello) string
}
