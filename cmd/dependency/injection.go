package dependency

import (
	"hello-clean-architecture/internal/infrastructure/api/handler"
	"hello-clean-architecture/internal/infrastructure/api/router"
	"hello-clean-architecture/internal/interface/controller"
	"hello-clean-architecture/internal/interface/presenter_impl"
	"hello-clean-architecture/internal/usecase/hello"
	"hello-clean-architecture/internal/usecase/presenter"
	"net/http"
)

// injection は依存関係を解決し、各コンポーネントを構築するための構造体
type injection struct{}

// Injection はアプリケーションの主要な依存関係の構築を抽象化するインターフェース
type Injection interface {
	NewRouter() http.Handler
}

// NewInjection は新しいInjectionのインスタンスを作成
func NewInjection() Injection {
	return &injection{}
}

// NewRouter はアプリケーション全体のHTTPルーターを構築
func (i *injection) NewRouter() http.Handler {
	return router.New(i.NewHelloHandler())
}

// NewHello はHelloHandlerを構築
func (i *injection) NewHelloHandler() handler.Hello {
	return handler.New(i.NewHelloController())
}

// NewHelloController はHelloControllerを構築
func (i *injection) NewHelloController() controller.Hello {
	return controller.New(i.NewService())
}

// NewServiceはHelloServiceを構築
func (i *injection) NewService() hello.HelloService {
	return hello.New(i.NewHelloPresenter())
}

// NewHelloPresenter はHelloPresenterの実装を構築
func (i *injection) NewHelloPresenter() presenter.Hello {
	return presenter_impl.New()
}
