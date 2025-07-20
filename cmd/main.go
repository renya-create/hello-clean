package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	startServer()
}

func startServer() {
	http.HandleFunc("/", handleHello)

	fmt.Println("サーバーを起動中: http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("サーバ起動失敗: ", err)
	}
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "Hello, World!\n")
}
