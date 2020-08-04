package server

import (
	"io"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"

	"github.com/iwaltgen/grpc-go-web-todo/pkg/log"
)

// NewJaegerTracer initialize opentracing
func NewJaegerTracer() (opentracing.Tracer, io.Closer, error) {
	cfg := &jaegercfg.Configuration{
		ServiceName: "grpc-go-web-todo",
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
	}

	cfg, err := cfg.FromEnv()
	if err != nil {
		log.L("server.tracing").
			Warn("config from environment error", log.Error(err))
	}

	return cfg.NewTracer(jaegercfg.MaxTagValueLength(8 * 1024))
}
