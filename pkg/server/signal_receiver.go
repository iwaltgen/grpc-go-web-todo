package server

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/iwaltgen/grpc-go-web-todo/pkg/log"
)

type signalTermReceiver struct {
	logger *log.Logger
	Cancel context.CancelFunc
}

// create term signal receiver
func newSignalTermReceiver(cancel context.CancelFunc) *signalTermReceiver {
	return &signalTermReceiver{
		logger: log.L("server.signal"),
		Cancel: cancel,
	}
}

// Serve start serving
func (r *signalTermReceiver) Serve(ctx context.Context) {
	signals := []os.Signal{
		syscall.SIGINT,
		syscall.SIGKILL,
		syscall.SIGTERM,
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, signals...)

	r.logger.Info("serve started",
		log.Stringer("signal", syscall.SIGINT),
		log.Stringer("signal", syscall.SIGKILL),
		log.Stringer("signal", syscall.SIGTERM),
	)

	select {
	case sig := <-c:
		r.logger.Info("os signal received", log.Stringer("signal", sig))
		r.Cancel()

	case <-ctx.Done():
	}
}
