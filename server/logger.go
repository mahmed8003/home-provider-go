package server

import (
	"strconv"
	"time"

	"github.com/kataras/iris/context"
	"github.com/rs/zerolog"
)

type loggerConfig struct {
	logger zerolog.Logger
	// Status displays status code (bool).
	//
	// Defaults to true.
	status bool
	// IP displays request's remote address (bool).
	//
	// Defaults to true.
	ip bool
	// Method displays the http method (bool).
	//
	// Defaults to true.
	method bool
	// Path displays the request path (bool).
	//
	// Defaults to true.
	path bool

	enableLogs bool
}

type requestLoggerMiddleware struct {
	config loggerConfig
}

/*
NewLogger :
*/
func NewLogger(logger zerolog.Logger, status bool, ip bool, method bool, path bool, enableLogs bool) context.Handler {
	l := &requestLoggerMiddleware{config: loggerConfig{
		logger:     logger,
		status:     status,
		ip:         ip,
		method:     method,
		path:       path,
		enableLogs: enableLogs,
	}}
	return l.ServeHTTP
}

// Serve serves the middleware
func (l *requestLoggerMiddleware) ServeHTTP(ctx context.Context) {

	if !l.config.enableLogs {
		ctx.Next()
		return
	}

	var latency time.Duration
	var startTime, endTime time.Time
	startTime = time.Now()

	ctx.Next()

	endTime = time.Now()
	latency = endTime.Sub(startTime)

	log := l.config.logger.Info()
	log.Time("at", endTime)
	log.Int64("latency", int64(latency/time.Microsecond))

	if l.config.status {
		log.Str("status", strconv.Itoa(ctx.GetStatusCode()))
	}

	if l.config.ip {
		log.Str("ip", ctx.RemoteAddr())
	}

	if l.config.method {
		log.Str("method", ctx.Method())
	}

	if l.config.path {
		log.Str("path", ctx.Path())
	}

	log.Msg("")
}
