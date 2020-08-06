package usecase

import (
	"context"
	"sync"

	"github.com/iwaltgen/grpc-go-web-todo/pkg/entity"
	"github.com/iwaltgen/grpc-go-web-todo/pkg/log"
	"github.com/iwaltgen/grpc-go-web-todo/pkg/repository"
)

// Todo domain logic
type Todo struct {
	logger *log.Logger
	repo   *repository.Todo
}

var (
	defaultTodo     *Todo
	defaultTodoOnce sync.Once
)

// DefaultTodo get default todo
func DefaultTodo() *Todo {
	defaultTodoOnce.Do(func() {
		defaultTodo = createTodo()
	})
	return defaultTodo
}

func newTodo(repo *repository.Todo) *Todo {
	return &Todo{
		logger: log.L("usecase.todo"),
		repo:   repo,
	}
}

// ListTodos get all todo
func (s *Todo) ListTodos(ctx context.Context) (ret []*entity.Todo, err error) {
	return s.repo.FindAll(ctx)
}
