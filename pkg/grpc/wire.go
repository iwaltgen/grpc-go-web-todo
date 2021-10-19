//go:build wireinject
// +build wireinject

package grpc

import (
	"github.com/google/wire"

	"github.com/iwaltgen/grpc-go-web-todo/pkg/usecase"

	todov1 "github.com/iwaltgen/grpc-go-web-todo/api/todo/v1"
)

func createTodoService() todov1.TodoServiceServer {
	wire.Build(wire.NewSet(newTodoService), usecase.WireSet)
	return nil
}
