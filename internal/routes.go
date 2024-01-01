package internal

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog/v2"
	"github.com/riandyrn/otelchi"
)

func NewRouter(config Config) *chi.Mux {
	r := chi.NewRouter()

	tracer := NewTracer(context.Background(), config.ServiceName)

	// Middlewares
	r.Use(httplog.RequestLogger(NewLogger(config.ServiceName)))
	r.Use(otelchi.Middleware(
		config.ServiceName,
		otelchi.WithChiRoutes(r),
		otelchi.WithTracerProvider(tracer),
	))

	r.Get("/ping", ping(config.PongerAddr))
	r.Get("/pong", pong())

	return r

}
