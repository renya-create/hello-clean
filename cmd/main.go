package main

import (
	"fmt"
	"hello-clean/internal/registry"
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
	interactor := registry.NewInteractor()
	appRouter := interactor.NewAppRouter()

	port := ":8080"
	slog.Info(fmt.Sprintf("Server starting on port %s", port))

	if err := http.ListenAndServe(port, appRouter); err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}
