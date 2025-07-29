package hello

import (
	"hello-clean-architecture/internal/domain/message"
	"hello-clean-architecture/internal/usecase/presenter"
)

// ビジネスロジックの本体
// HelloServiceインターフェースの実装
type helloService struct {
	helloPresenter presenter.Hello
}

// "Hello, World!"メッセージに関連するビジネスロジックを定義するインターフェース
type HelloService interface {
	GetHelloMessage() string
}

// helloServiceが依存するHelloPresenterの具体的な実装を外部から注入
func New(hp presenter.Hello) HelloService {
	return &helloService{
		helloPresenter: hp,
	}
}

// "Hello, World!"メッセージを生成し、プレゼンターを通じて整形
func (hs *helloService) GetHelloMessage() string {
	// プレゼンターインターフェースを通じてメッセージを整形
	// このサービスはプレゼンターの具体的な実装（例：JSON、プレーンテキストなど）を知らない

	hello := message.New("Hello, World!")
	return hs.helloPresenter.FormatHelloMessage(hello)
}
