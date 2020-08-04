package server

import (
	"time"

	"github.com/labstack/echo/v4"

	"github.com/iwaltgen/grpc-go-web-todo/pkg/log"
)

const httpAccessLog = "finished call"

// http log middleware
func httpLogger(logger *log.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			start := time.Now()
			req := c.Request()
			res := c.Response()

			if err = next(c); err != nil {
				c.Error(err)
			}

			status := res.Status
			fields := []log.Field{
				log.String("peer.address", c.RealIP()),
				log.Int("http.status", status),
				log.String("http.uri", req.RequestURI),
				log.Int64("file.bytes", res.Size),
				log.Duration("http.duration", time.Since(start)),
			}

			switch {
			case 500 <= status:
				logger.Error(httpAccessLog, fields...)
			case 400 <= status:
				logger.Warn(httpAccessLog, fields...)
			default:
				logger.Debug(httpAccessLog, fields...)
			}
			return
		}
	}
}
