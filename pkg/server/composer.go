package server

import (
	"context"
	"sync"

	"github.com/iwaltgen/grpc-go-web-todo/pkg/log"
	"github.com/opentracing/opentracing-go"
)

// Composer all server management
type Composer struct {
	logger *log.Logger
}

// NewComposer create CopositeServer
func NewComposer() *Composer {
	return &Composer{log.L("server.composer")}
}

// Serve start serving
func (s *Composer) Serve(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	tracer, closer, err := NewJaegerTracer()
	if err != nil {
		s.logger.Warn("new opentracing tracer error", log.Error(err))
	}
	defer func() {
		_ = closer.Close()
	}()
	opentracing.SetGlobalTracer(tracer)

	type server interface {
		Serve(context.Context)
	}

	gsrv := NewGRPC()
	servers := []server{
		gsrv,
		NewHTTP(gsrv.Server),
		newSignalTermReceiver(cancel),
	}

	var wg sync.WaitGroup
	wg.Add(len(servers))

	for _, s := range servers {
		go func(s server) {
			s.Serve(ctx)
			wg.Done()
		}(s)
	}

	s.logger.Info("serve started", log.Int("count", len(servers)))
	wg.Wait()
}
