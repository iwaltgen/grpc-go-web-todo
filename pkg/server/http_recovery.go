package server

import (
	"fmt"
	"runtime"

	"github.com/labstack/echo/v4"

	"github.com/iwaltgen/grpc-go-web-todo/pkg/log"
)

// http panic handler
func httpRecovery(logger *log.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer func() {
				if r := recover(); r != nil {
					err, ok := r.(error)
					if !ok {
						err = fmt.Errorf("%v", r)
					}

					stack := make([]byte, panicPrintStackSize)
					length := runtime.Stack(stack, panicPrintStackAll)
					logger.Error("[PANIC RECOVER]",
						log.ByteString("stack", stack[:length]),
						log.Error(err),
					)
					c.Error(err)
				}
			}()
			return next(c)
		}
	}
}
