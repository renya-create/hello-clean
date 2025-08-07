package main

import (
	"flag"
	"hello-clean-architecture/cmd/dependency"
	"hello-clean-architecture/internal/infrastructure/server"
	"log/slog"
	"os"
)

const (
	DefaultPort = ":8080"
)

func main() {
	err := run()
	if err != nil {
		slog.Error("Application failed to start", "error", err)
		os.Exit(1)
	}
}

func getPort() string {
	port := flag.String("port", DefaultPort, "Server port (e.g., :8080)")
	flag.Parse()
	return *port
}

func run() error {
	dependencyInjection := dependency.NewInjection()
	router := dependencyInjection.NewRouter()

	port := getPort()

	return server.StartServer(port, router)
}
