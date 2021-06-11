package transport

import (
	"fmt"
	"net/http"

	"github.mpi-internal.com/hu/homeoffice-calendar-api/internal/logger"
)

func Error(l logger.Logger, status int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l.With("method", r.Method).
			With("uri", r.RequestURI).
			With("response_status", fmt.Sprint(status)).
			Warning(http.StatusText(status))

		http.Error(w, http.StatusText(status), status)
	}
}
