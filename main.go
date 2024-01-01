package main

import (
	"log/slog"
	"net/http"
	"os"
	"otel/internal"
)

func main() {
	config, err := internal.ParseEnv()
	if err != nil {
		slog.With(err).Error("failed to parse env")
		os.Exit(1)
	}

	r := internal.NewRouter(config)

	_ = http.ListenAndServe(":3000", r)
}
