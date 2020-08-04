package server

import (
	"fmt"
	"runtime"

	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/iwaltgen/grpc-go-web-todo/pkg/log"
)

// gRPC panic handler
func grpcRecovery(logger *log.Logger) grpc_recovery.RecoveryHandlerFunc {
	return func(r interface{}) error {
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

		return status.Error(codes.Internal, err.Error())
	}
}
