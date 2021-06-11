package logger

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
)

// ChiRequestLoggerMiddleware returns a logger handler using a custom LogFormatter.
func ChiRequestLoggerMiddleware(l Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			reqID := middleware.GetReqID(r.Context())
			scheme := "http"

			if r.TLS != nil {
				scheme = "https"
			}

			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			t1 := time.Now()

			defer func() {
				l.With("request_uri", fmt.Sprintf("%s://%s%s", scheme, r.Host, r.RequestURI)).
					With("reqID", reqID).
					With("proto", r.Proto).
					With("method", r.Method).
					With("remote_addr", r.RemoteAddr).
					With("status", fmt.Sprint(ww.Status())).
					With("bytes", fmt.Sprint(ww.BytesWritten())).
					With("response_time", fmt.Sprint(time.Since(t1))).
					Debug("Request done")
			}()

			next.ServeHTTP(ww, r)
		}

		return http.HandlerFunc(fn)
	}
}
