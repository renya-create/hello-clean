package handler

import (
	"fmt"
	controllers "hello-clean/internal/interface/controller"
	"net/http"
)

// HTTPリクエストを受け取り、interface層のコントローラーに処理を委譲する、
// Webフレームワーク（ここではGoの標準ライブラリnet/http）に特化したHTTPハンドラー

type helloHandler struct {
	helloController controllers.HelloController
}

type HelloHandler interface {
	HandleHello(w http.ResponseWriter, r *http.Request)
}

func NewHelloHandler(hc controllers.HelloController) HelloHandler {
	return &helloHandler{
		helloController: hc,
	}
}

func (hh *helloHandler) HandleHello(w http.ResponseWriter, r *http.Request) {
	message := hh.helloController.GetHello()
	// w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, message)
}
