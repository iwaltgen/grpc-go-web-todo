// +build wireinject

package usecase

import (
	"github.com/google/wire"

	"github.com/iwaltgen/grpc-go-web-todo/pkg/repository"
)

func createTodo() *Todo {
	wire.Build(wire.NewSet(newTodo, repository.WireSet))
	return nil
}
