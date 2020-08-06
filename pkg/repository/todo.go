package repository

import (
	"context"
	"sync"

	"github.com/iwaltgen/grpc-go-web-todo/pkg/entity"
	"github.com/iwaltgen/grpc-go-web-todo/pkg/log"
)

// Todo is CRUD todo entity
type Todo struct {
	logger    *log.Logger
	container sync.Map
}

var (
	defaultTodo     *Todo
	defaultTodoOnce sync.Once
)

// DefaultTodo is get default todo repository
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

// FindAll is get all todo entity
func (s *Todo) FindAll(ctx context.Context) (ret []*entity.Todo, err error) {
	s.container.Range(func(k, v interface{}) bool {
		ret = append(ret, v.(*entity.Todo))
		return true
	})
	return ret, nil
}

// Create is create todo entity
func (s *Todo) Create(ctx context.Context, v *entity.Todo) error {
	s.container.Store(v.ID, v)
	return nil
}

// Update is update todo entity
func (s *Todo) Update(ctx context.Context, v *entity.Todo) error {
	s.container.Store(v.ID, v)
	return nil
}

// DeleteByID is delete todo entity
func (s *Todo) DeleteByID(ctx context.Context, id string) error {
	s.container.Delete(id)
	return nil
}
