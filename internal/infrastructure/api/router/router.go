package router

import (
	"hello-clean/internal/infrastructure/api/handler"
	"net/http"
)

// HTTPパスと、それに対応するHTTPハンドラーを紐付けるルーティングを定義

func NewRouter(h handler.HelloHandler) *http.ServeMux {
	mux := http.NewServeMux()
	// http.HandleFuncが要求するfunc(ResponseWriter, *Request)のシグネチャと一致するため、直接登録可能
	mux.HandleFunc("/", h.HandleHello)
	return mux
}
