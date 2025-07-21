package registry

import (
	"hello-clean/internal/infrastructure/api/handler"
	"hello-clean/internal/infrastructure/api/router"
	controllers "hello-clean/internal/interface/controller"
	presenters "hello-clean/internal/interface/presenter"
	"hello-clean/internal/usecase/presenter"
	"hello-clean/internal/usecase/service"
	"net/http"
)

// interactor は依存関係を解決し、各コンポーネントを構築するための構造体
type interactor struct{}

// Interactor はアプリケーションの主要な依存関係の構築を抽象化するインターフェース
type Interactor interface {
	NewAppRouter() *http.ServeMux
}

// NewInteractor は新しいInteractorのインスタンスを作成
func NewInteractor() Interactor {
	return &interactor{}
}

// NewAppRouter はアプリケーション全体のHTTPルーターを構築
func (i *interactor) NewAppRouter() *http.ServeMux {
	return router.NewRouter(i.NewHelloHandler())
}

// NewHelloHandler はHelloHandlerを構築
func (i *interactor) NewHelloHandler() handler.HelloHandler {
	return handler.NewHelloHandler(i.NewHelloController())
}

// NewHelloController はHelloControllerを構築
func (i *interactor) NewHelloController() controllers.HelloController {
	return controllers.NewHelloController(i.NewHelloService())
}

// NewHelloService はHelloServiceを構築
func (i *interactor) NewHelloService() service.HelloService {
	return service.NewHelloService(i.NewHelloPresenter())
}

// NewHelloPresenter はHelloPresenterの実装を構築
func (i *interactor) NewHelloPresenter() presenter.HelloPresenter {
	return presenters.NewHelloPresenter()
}
