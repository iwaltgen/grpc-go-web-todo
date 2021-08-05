// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package grpc

import (
	"github.com/iwaltgen/grpc-go-web-todo/api/todo/v1"
	"github.com/iwaltgen/grpc-go-web-todo/pkg/usecase"
)

// Injectors from wire.go:

func createTodoService() todov1.TodoServiceServer {
	todo := usecase.DefaultTodo()
	todoServiceServer := newTodoService(todo)
	return todoServiceServer
}
