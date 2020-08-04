//go:generate statik -ns=server.http -src=../../public -f

package server

import (
	"context"
	"net/http"
	"strings"

	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
	"github.com/rakyll/statik/fs"

	"github.com/iwaltgen/grpc-go-web-todo/pkg/log"
	// _ "github.com/iwaltgen/grpc-go-web-todo/pkg/server/statik"
)

// HTTP HTTP Server
type HTTP struct {
	x509Generator
	*echo.Echo
	fs     http.FileSystem
	logger *log.Logger
}

// NewHTTP create HTTP server
func NewHTTP() *HTTP {
	logger := log.L("server.http")

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.Use(jaegertracing.Trace(opentracing.GlobalTracer()))
	e.Use(httpLogger(logger))
	e.Use(httpRecovery(logger))

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

	return &HTTP{
		Echo:   e,
		fs:     statikFS,
		logger: logger,
	}
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

	h.logger.Info("serve started", log.String("addr", addr), log.Bool("insecure", insecure))
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
