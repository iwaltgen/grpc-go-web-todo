package usecase

import (
	"context"
	"sort"
	"sync"

	"github.com/gogo/status"
	"google.golang.org/grpc/codes"

	"github.com/iwaltgen/grpc-go-web-todo/pkg/entity"
	"github.com/iwaltgen/grpc-go-web-todo/pkg/event"
	"github.com/iwaltgen/grpc-go-web-todo/pkg/log"
	"github.com/iwaltgen/grpc-go-web-todo/pkg/repository"
)

// Todo is domain logic for todo entity
type Todo struct {
	logger  *log.Logger
	repo    *repository.Todo
	emitter *event.Emitter
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
		logger:  log.L("usecase.todo"),
		repo:    repo,
		emitter: &event.Emitter{},
	}
}

// ListTodos is get all todo entity
func (s *Todo) ListTodos(ctx context.Context) (ret []*entity.Todo, err error) {
	ret, err = s.repo.FindAll(ctx)
	sort.Sort(sort.Reverse(todoSortByCreatedAt(ret)))
	return ret, err
}

// CreateTodo is create todo entity
func (s *Todo) CreateTodo(ctx context.Context, v *entity.Todo) error {
	err := s.repo.Create(ctx, v)
	if err == nil {
		s.emitter.Publish(event.EventCreate, v)
	}
	return err
}

// UpdateTodo is update todo entity
func (s *Todo) UpdateTodo(ctx context.Context, v *entity.Todo) error {
	err := s.repo.Update(ctx, v)
	if err == nil {
		s.emitter.Publish(event.EventUpdate, v)
	}
	return err
}

// DeleteTodo is delete todo entity
func (s *Todo) DeleteTodo(ctx context.Context, id string) error {
	if ret, ok := s.repo.FindByID(ctx, id); ok {
		err := s.repo.DeleteByID(ctx, id)
		if err == nil {
			s.emitter.Publish(event.EventDelete, ret)
		}
		return err
	}
	return status.Errorf(codes.NotFound, "not found error: %s", id)
}

// Subscribe is subscribe entity change event
func (s *Todo) Subscribe(handler event.Subscription, evt event.Event) func() {
	return s.emitter.Subscribe(handler, event.WithEvent(evt))
}

// todo entity sort by modified time
type todoSortByCreatedAt []*entity.Todo

func (s todoSortByCreatedAt) Len() int {
	return len(s)
}

func (s todoSortByCreatedAt) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s todoSortByCreatedAt) Less(i, j int) bool {
	return s[i].CreatedAt.Before(s[j].CreatedAt)
}
