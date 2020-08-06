package repository

import (
	"context"
	"sync"

	"github.com/iwaltgen/grpc-go-web-todo/pkg/entity"
	"github.com/iwaltgen/grpc-go-web-todo/pkg/log"
)

// Todo CRUD entity
type Todo struct {
	logger    *log.Logger
	container sync.Map
}

var (
	defaultTodo     *Todo
	defaultTodoOnce sync.Once
)

// DefaultTodo get default todo
func DefaultTodo() *Todo {
	defaultTodoOnce.Do(func() {
		defaultTodo = newTodo()
	})
	return defaultTodo
}

func newTodo() *Todo {
	return &Todo{
		logger: log.L("repository.todo"),
	}
}

// FindAll get all todo
func (s *Todo) FindAll(context.Context) (ret []*entity.Todo, err error) {
	s.container.Range(func(k, v interface{}) bool {
		ret = append(ret, v.(*entity.Todo))
		return false
	})
	return ret, nil
}
