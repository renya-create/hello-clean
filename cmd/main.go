package main

import (
	"flag"
	"fmt"
	"hello-clean-architecture/internal/dependency"
	"log/slog"
	"net/http"
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
	slog.Info(fmt.Sprintf("Server starting on port %s", port))

	if err := http.ListenAndServe(port, router); err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}
