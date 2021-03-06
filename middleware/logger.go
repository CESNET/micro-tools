package middleware

import (
	"net/http"
	"time"

	"github.com/CESNET/micro-tools/log"
	"github.com/go-chi/chi/middleware"
)

// Logger is a middleware to log http requests.
func Logger(logger log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			wrap := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			next.ServeHTTP(wrap, r)

			logger.Debug().
				Str("request", r.Header.Get("X-Request-ID")).
				Str("proto", r.Proto).
				Str("method", r.Method).
				Str("remote", r.RemoteAddr).
				Int("status", wrap.Status()).
				Str("path", r.URL.Path).
				Dur("duration", time.Since(start)).
				Int("bytes", wrap.BytesWritten()).
				Msg("")
		})
	}
}
