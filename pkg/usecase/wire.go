// +build wireinject

package usecase

import (
	"github.com/google/wire"
)

func createTodoService() *TodoService {
	wire.Build(wire.NewSet(newTodoService))
	return nil
}
