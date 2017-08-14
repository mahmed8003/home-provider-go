package server

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/rs/zerolog"
)

/*
RequestLogger :
*/
type RequestLogger func(http.Handler) http.Handler

/*
logger :
*/
type requestLogger struct {
	logger     zerolog.Logger
	enableLogs bool
}

/*
NewRequestLogger :
*/
func NewRequestLogger(logger zerolog.Logger, enableLogs bool) RequestLogger {
	l := requestLogger{
		logger:     logger,
		enableLogs: enableLogs,
	}
	return l.middleware
}

func (l requestLogger) middleware(next http.Handler) http.Handler {
	if !l.enableLogs {
		return next
	}

	fn := func(w http.ResponseWriter, r *http.Request) {

		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

		start := time.Now()

		defer func() {
			end := time.Now()
			latency := end.Sub(start)

			log := l.logger.Info()
			log.Time("at", end).
				Int64("latency", int64(latency/time.Microsecond)).
				Int("status", ww.Status()).
				Int("bytes", ww.BytesWritten()).
				Str("ip", r.RemoteAddr).
				Str("method", r.Method).
				Str("path", r.URL.Path).
				Msg("")
		}()

		next.ServeHTTP(ww, r)
	}

	return http.HandlerFunc(fn)
}
