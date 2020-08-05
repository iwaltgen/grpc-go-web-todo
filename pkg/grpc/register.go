package grpc

import (
	ggrpc "google.golang.org/grpc"
)

// Register register all service server
func Register(srv *ggrpc.Server) {
	RegisterTodoServiceServer(srv)
}
