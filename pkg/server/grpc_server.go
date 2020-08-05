package server

import (
	"context"
	"fmt"
	"net"
	"runtime"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	pgrpc "github.com/iwaltgen/grpc-go-web-todo/pkg/grpc"
	"github.com/iwaltgen/grpc-go-web-todo/pkg/log"
)

// GRPC gRPC Server
type GRPC struct {
	x509Generator
	*grpc.Server
	logger *log.Logger
}

// NewGRPC create gRPC server
func NewGRPC() (ret *GRPC) {
	logger := log.L("server.grpc")

	ret = &GRPC{
		logger: logger,
	}

	certificate, err := ret.newCertificate(DevHosts)
	if err != nil {
		logger.Panic("new x509 certificate error", log.Error(err))
	}

	creds := credentials.NewServerTLSFromCert(&certificate)
	allowAll := func(ctx context.Context) (context.Context, error) { return ctx, nil }

	zapOptDuration := grpc_zap.WithDurationField(grpc_zap.DurationToDurationField)
	srv := grpc.NewServer(
		grpc.Creds(creds),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_opentracing.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(logger, zapOptDuration),
			grpc_validator.UnaryServerInterceptor(),
			grpc_auth.UnaryServerInterceptor(allowAll),
			grpc_recovery.UnaryServerInterceptor(
				grpc_recovery.WithRecoveryHandler(ret.recovery),
			),
		)),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_opentracing.StreamServerInterceptor(),
			grpc_zap.StreamServerInterceptor(logger, zapOptDuration),
			grpc_validator.StreamServerInterceptor(),
			grpc_auth.StreamServerInterceptor(allowAll),
			grpc_recovery.StreamServerInterceptor(
				grpc_recovery.WithRecoveryHandler(ret.recovery),
			),
		)),
	)
	ret.Server = srv

	logger.Info("enable middleware", log.String("type", "tags"))
	logger.Info("enable middleware",
		log.String("type", "opentracing"),
		log.String("tracer", "jaeger tracer"),
		log.String("url", "http://localhost:16686"),
	)
	logger.Info("enable middleware", log.String("type", "logger"))
	logger.Info("enable middleware", log.String("type", "validator"))
	logger.Info("enable middleware", log.String("type", "auth"))
	logger.Info("enable middleware", log.String("type", "recovery"))

	pgrpc.Register(srv)
	reflection.Register(srv)
	return ret
}

// Serve start serving
func (g *GRPC) Serve(ctx context.Context) {
	addr := grpcPort
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		g.logger.Panic("failed to listen", log.Error(err))
	}

	go func() {
		if err := g.Server.Serve(ln); err != nil {
			g.logger.Error("serve error", log.Error(err))
		} else {
			g.logger.Info("serve done", log.String("addr", addr))
		}
	}()
	g.logger.Info("serve started", log.String("addr", addr))

	<-ctx.Done()

	stopped := make(chan struct{})
	go func() {
		g.GracefulStop()
		close(stopped)
	}()

	t := time.NewTimer(shutdownWaitTimeout)
	select {
	case <-t.C:
		g.Stop()

	case <-stopped:
		t.Stop()
	}
}

func (g *GRPC) recovery(r interface{}) error {
	err, ok := r.(error)
	if !ok {
		err = fmt.Errorf("%v", r)
	}

	stack := make([]byte, panicPrintStackSize)
	length := runtime.Stack(stack, panicPrintStackAll)
	g.logger.Error("[PANIC RECOVER]",
		log.ByteString("stack", stack[:length]),
		log.Error(err),
	)

	return status.Error(codes.Internal, err.Error())
}
