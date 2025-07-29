package main

import (
	"fmt"
	"hello-clean-architecture/internal/dependency"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	err := run()
	if err != nil {
		slog.Error("Application failed to start", "error", err)
		os.Exit(1)
	}
}

func run() error {
	dependencyInjection := dependency.NewInjection()
	router := dependencyInjection.NewRouter()

	port := ":8080"
	slog.Info(fmt.Sprintf("Server starting on port %s", port))

	if err := http.ListenAndServe(port, router); err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}
