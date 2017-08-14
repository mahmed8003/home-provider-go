package server

import (
	"net/http"
	"time"

	routing "github.com/go-ozzo/ozzo-routing"
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
func NewRequestLogger(logger zerolog.Logger, enableLogs bool) routing.Handler {
	l := requestLogger{
		logger:     logger,
		enableLogs: enableLogs,
	}
	return l.middleware
}

func (l requestLogger) middleware(c *routing.Context) error {
	if !l.enableLogs {
		return c.Next()
	}

	start := time.Now()
	req := c.Request
	rw := &LogResponseWriter{c.Response, http.StatusOK, 0}
	c.Response = rw

	err := c.Next()

	end := time.Now()
	latency := end.Sub(start)

	var log *zerolog.Event
	if err != nil {
		log = l.logger.Error()
		log.Err(err)
	} else {
		log = l.logger.Info()
	}

	log.Time("at", end).
		Int64("latency", int64(latency/time.Microsecond)).
		Int("status", rw.Status).
		Int64("bytes", rw.BytesWritten).
		Str("ip", getClientIP(req)).
		Str("method", req.Method).
		Str("path", req.URL.Path).
		Msg("")

	return err
}

/*
LogResponseWriter :
*/
type LogResponseWriter struct {
	http.ResponseWriter
	Status       int
	BytesWritten int64
}

/*
Write :
*/
func (r *LogResponseWriter) Write(p []byte) (int, error) {
	written, err := r.ResponseWriter.Write(p)
	r.BytesWritten += int64(written)
	return written, err
}

/*
WriteHeader :
*/
func (r *LogResponseWriter) WriteHeader(status int) {
	r.Status = status
	r.ResponseWriter.WriteHeader(status)
}

/*
GetClientIP :
*/
func getClientIP(req *http.Request) string {
	ip := req.Header.Get("X-Real-IP")
	if ip == "" {
		ip = req.Header.Get("X-Forwarded-For")
		if ip == "" {
			ip = req.RemoteAddr
		}
	}
	return ip
}
