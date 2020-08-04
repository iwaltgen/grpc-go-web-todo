//go:generate statik -ns=server.http -src=../../public -include=*.jpg,*.png,*.html,*.css,*.js

package server

import (
	"context"
	"fmt"
	"net/http"
	"runtime"
	"strings"

	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/opentracing/opentracing-go"
	"github.com/rakyll/statik/fs"

	"github.com/iwaltgen/grpc-go-web-todo/pkg/log"
	_ "github.com/iwaltgen/grpc-go-web-todo/pkg/server/statik" // frontend embedded resource
)

// HTTP HTTP Server
type HTTP struct {
	x509Generator
	*echo.Echo
	logger *log.Logger
}

// NewHTTP create HTTP server
func NewHTTP() (ret *HTTP) {
	logger := log.L("server.http")

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	ret = &HTTP{
		Echo:   e,
		logger: logger,
	}

	e.Use(jaegertracing.Trace(opentracing.GlobalTracer()))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339}\t${method}\t${uri}\t${status} ${latency_human}\t" +
			"${remote_ip} ${bytes_in} ${bytes_out} ${error}\n",
	}))
	e.Use(ret.recovery)

	logger.Info("enable middleware",
		log.String("type", "opentracing"),
		log.String("tracer", "jaeger tracer"),
		log.String("url", "http://localhost:16686"),
	)
	logger.Info("enable middleware", log.String("type", "logger"))
	logger.Info("enable middleware", log.String("type", "recovery"))

	statikFS, err := fs.NewWithNamespace("server.http")
	if err != nil {
		logger.Panic("new statik file system error", log.Error(err))
	}

	e.GET("/*", echo.WrapHandler(http.FileServer(statikFS)))
	return ret
}

// Serve start serving
func (h *HTTP) Serve(ctx context.Context) {
	addr := ":443"
	insecure := !strings.Contains(addr, "443")

	go func() {
		if err := h.startServe(addr, insecure); err != nil && err != http.ErrServerClosed {
			h.logger.Error("serve error", log.Error(err))
			return
		}

		h.logger.Info("serve done", log.String("addr", addr))
	}()

	h.logger.Info("serve started", log.String("addr", addr))
	<-ctx.Done()

	sctx, cancel := context.WithTimeout(context.Background(), shutdownWaitTimeout)
	defer cancel()
	if err := h.Shutdown(sctx); err != nil {
		h.logger.Error("shutdown error", log.Error(err))
	}
}

func (h *HTTP) startServe(addr string, insecure bool) error {
	if insecure {
		return h.Start(addr)
	}

	cert, key, err := h.newX509KeyPairBytes(DevHosts)
	if err != nil {
		return err
	}

	return h.StartTLS(addr, cert, key)
}

func (h *HTTP) recovery(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer func() {
			if r := recover(); r != nil {
				err, ok := r.(error)
				if !ok {
					err = fmt.Errorf("%v", r)
				}

				stack := make([]byte, panicPrintStackSize)
				length := runtime.Stack(stack, panicPrintStackAll)
				h.logger.Error("[PANIC RECOVER]",
					log.ByteString("stack", stack[:length]),
					log.Error(err),
				)
				c.Error(err)
			}
		}()
		return next(c)
	}
}
