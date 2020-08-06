package usecase

import (
	"context"
	"sort"
	"sync"

	"github.com/iwaltgen/grpc-go-web-todo/pkg/entity"
	"github.com/iwaltgen/grpc-go-web-todo/pkg/log"
	"github.com/iwaltgen/grpc-go-web-todo/pkg/repository"
)

// Todo is domain logic for todo entity
type Todo struct {
	logger *log.Logger
	repo   *repository.Todo
}

var (
	defaultTodo     *Todo
	defaultTodoOnce sync.Once
)

// DefaultTodo is get default todo usecase
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

// ListTodos is get all todo entity
func (s *Todo) ListTodos(ctx context.Context) (ret []*entity.Todo, err error) {
	ret, err = s.repo.FindAll(ctx)
	sort.Sort(todoSortByModifiedAt(ret))
	return ret, err
}

// CreateTodo is create todo entity
func (s *Todo) CreateTodo(ctx context.Context, v *entity.Todo) error {
	return s.repo.Create(ctx, v)
}

// UpdateTodo is update todo entity
func (s *Todo) UpdateTodo(ctx context.Context, v *entity.Todo) error {
	return s.repo.Update(ctx, v)
}

// DeleteTodo is delete todo entity
func (s *Todo) DeleteTodo(ctx context.Context, id string) error {
	return s.repo.DeleteByID(ctx, id)
}

// todo entity sort by modified time
type todoSortByModifiedAt []*entity.Todo

func (s todoSortByModifiedAt) Len() int {
	return len(s)
}

func (s todoSortByModifiedAt) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s todoSortByModifiedAt) Less(i, j int) bool {
	return s[i].ModifiedAt.Before(s[j].ModifiedAt)
}
