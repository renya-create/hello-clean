package server

import (
	"fmt"
	"log/slog"
	"net/http"
)

func StartServer(port string, router *http.ServeMux) error {
	slog.Info(fmt.Sprintf("Server starting on port %s", port))

	if err := http.ListenAndServe(port, router); err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}
