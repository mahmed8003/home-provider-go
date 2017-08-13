package server

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

/*
Logger :
*/
func Logger(logger zerolog.Logger, enableLogs bool) gin.HandlerFunc {

	return func(c *gin.Context) {

		if !enableLogs {
			c.Next()
			return
		}

		// Start timer
		start := time.Now()
		path := c.Request.URL.Path

		// Process request
		c.Next()

		end := time.Now()
		latency := end.Sub(start)

		log := logger.Info()
		log.Time("at", end).
			Int64("latency", int64(latency/time.Microsecond)).
			Int("status", c.Writer.Status()).
			Str("ip", c.ClientIP()).
			Str("method", c.Request.Method).
			Str("path", path).
			Msg(c.Errors.ByType(gin.ErrorTypePrivate).String())
	}
}
