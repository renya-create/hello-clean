package controller

import "hello-clean-architecture/internal/usecase/hello"

// HelloControllerインターフェースの実装
type helloController struct {
	helloService hello.HelloService
}

// HTTPリクエストを処理し、usecase層のサービスを呼び出すインターフェース
type Hello interface {
	GetHello() string
}

// dependency層で呼ばれ、usecase層に注入される具体的なプレゼンター実装を提供
func New(h hello.HelloService) Hello {
	return &helloController{
		helloService: h,
	}
}

func (h *helloController) GetHello() string {
	// usecase層のサービスを呼び出す
	// コントローラーはビジネスロジックの具体的な実装には関心がなく、
	// 単にサービスに処理を依頼し、結果を受け取るだけ

	Message := h.helloService.GetHelloMessage()
	return Message
}
