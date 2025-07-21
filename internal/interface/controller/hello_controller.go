package controllers

import "hello-clean/internal/usecase/service"

// HelloControllerインターフェースの実装
type helloController struct {
	helloService service.HelloService
}

// HTTPリクエストを処理し、usecase層のサービスを呼び出すインターフェース
type HelloController interface {
	GetHello() string
}

// registry層で呼ばれ、usecase層に注入される具体的なプレゼンター実装を提供
func NewHelloController(hs service.HelloService) HelloController {
	return &helloController{
		helloService: hs,
	}
}

func (hc *helloController) GetHello() string {
	// usecase層のサービスを呼び出す
	// コントローラーはビジネスロジックの具体的な実装には関心がなく、
	// 単にサービスに処理を依頼し、結果を受け取るだけ

	formattedMessage := hc.helloService.GetHelloMessage()
	return formattedMessage
}
