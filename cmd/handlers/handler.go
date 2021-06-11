package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.mpi-internal.com/hu/homeoffice-calendar-api/internal/logger"
	"github.mpi-internal.com/hu/homeoffice-calendar-api/internal/transport"
)

const (
	moduleNameHTTPHandler = "http-handler"
)

// NewHandler returns the initialised Handler.
func NewHandler(l logger.Logger) http.Handler {
	handlerLogger := l.With("module", moduleNameHTTPHandler)
	r := chi.NewRouter()

	r.Use(logger.ChiRequestLoggerMiddleware(handlerLogger))
	r.NotFound(transport.Error(handlerLogger, http.StatusNotFound))
	r.MethodNotAllowed(transport.Error(handlerLogger, http.StatusMethodNotAllowed))

	r.Route("/_", func(r chi.Router) {
		r.Get("/health", transport.NewHealthHandler())
	})

	return r
}
