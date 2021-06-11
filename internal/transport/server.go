package transport

import (
	"net/http"

	"github.com/ory/graceful"

	"github.mpi-internal.com/hu/homeoffice-calendar-api/internal/logger"
)

func Start(l logger.Logger, cfg Config, h http.Handler) {
	l = l.With("address", cfg.Address)
	server := graceful.WithDefaults(&http.Server{
		Addr:    cfg.Address,
		Handler: h,
	})

	l.Info("HTTP server is starting")

	if err := graceful.Graceful(server.ListenAndServe, server.Shutdown); err != nil {
		l.Fatal("Failed to gracefully shutdown")
	}

	l.Info("Server was shutdown gracefully")
}
