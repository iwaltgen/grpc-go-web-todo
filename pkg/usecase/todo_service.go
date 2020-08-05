package usecase

import (
	"context"
	"sync"

	"github.com/gogo/status"
	"google.golang.org/grpc/codes"

	"github.com/iwaltgen/grpc-go-web-todo/pkg/entity"
	"github.com/iwaltgen/grpc-go-web-todo/pkg/log"
)

// TodoService todo domain service
type TodoService struct {
	logger *log.Logger
}

var (
	defaultTodoService *TodoService
	defaultTodoOnce    sync.Once
)

// DefaultTodoService get default todo service
func DefaultTodoService() *TodoService {
	defaultTodoOnce.Do(func() {
		defaultTodoService = createTodoService()
	})
	return defaultTodoService
}

func newTodoService() *TodoService {
	return &TodoService{
		logger: log.L("usecase.todo"),
	}
}

// ListTodos get all todo
func (s *TodoService) ListTodos(context.Context) ([]*entity.Todo, error) {
	// TODO(iwaltgen): repository query
	return nil, status.Error(codes.Unimplemented, codes.Unimplemented.String())
}
