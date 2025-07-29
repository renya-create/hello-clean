package handler

import (
	"fmt"
	"hello-clean-architecture/internal/interface/controller"
	"net/http"
)

// HTTPリクエストを受け取り、interface層のコントローラーに処理を委譲する、
// Webフレームワーク（ここではGoの標準ライブラリnet/http）に特化したHTTPハンドラー

type hello struct {
	controller controller.Hello
}

type Hello interface {
	HandleHello(w http.ResponseWriter, r *http.Request)
}

func New(hc controller.Hello) Hello {
	return &hello{
		controller: hc,
	}
}

func (h *hello) HandleHello(w http.ResponseWriter, r *http.Request) {
	message := h.controller.GetHello()
	fmt.Fprint(w, message)
}
