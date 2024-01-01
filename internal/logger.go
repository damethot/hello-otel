package internal

import (
	"log/slog"
	"time"

	"github.com/go-chi/httplog/v2"
)

func NewLogger(serviceName string) *httplog.Logger {
	logger := httplog.NewLogger(serviceName, httplog.Options{
		// JSON:             true,
		LogLevel:         slog.LevelDebug,
		Concise:          true,
		RequestHeaders:   true,
		MessageFieldName: "message",
		// TimeFieldFormat: time.RFC850,
		Tags: map[string]string{
			"version": "v1.0-81aa4244d9fc8076a",
			"env":     "dev",
		},
		QuietDownRoutes: []string{
			"/",
			"/ping",
		},
		QuietDownPeriod: 10 * time.Second,
		// SourceFieldName: "source",
	})

	return logger

}
