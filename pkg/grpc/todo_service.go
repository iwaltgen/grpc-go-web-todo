package grpc

import (
	"context"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"github.com/gogo/status"
	"github.com/iwaltgen/grpc-go-web-todo/pkg/log"
	"github.com/iwaltgen/grpc-go-web-todo/pkg/usecase"

	todov1 "github.com/iwaltgen/grpc-go-web-todo/api/todo/v1"
)

type todoService struct {
	logger      *log.Logger
	todoUsecase *usecase.TodoService
}

var (
	defaultTodoService todov1.TodoServiceServer
	defaultTodoOnce    sync.Once
)

// DefaultTodoService get default todo service
func DefaultTodoService() todov1.TodoServiceServer {
	defaultTodoOnce.Do(func() {
		defaultTodoService = createTodoService()
	})
	return defaultTodoService
}

// RegisterTodoServiceServer register TodoServiceServer
func RegisterTodoServiceServer(srv *grpc.Server) {
	todov1.RegisterTodoServiceServer(srv, DefaultTodoService())
}

func newTodoService(todoUsecase *usecase.TodoService) todov1.TodoServiceServer {
	return &todoService{
		logger:      log.L("grpc.todo"),
		todoUsecase: todoUsecase,
	}
}

func (s *todoService) ListTodos(ctx context.Context, req *todov1.ListTodosRequest) (*todov1.ListTodosResponse, error) {
	_, err := s.todoUsecase.ListTodos(ctx)
	return nil, err
}

func (s *todoService) CreateTodo(ctx context.Context, req *todov1.CreateTodoRequest) (*todov1.Unit, error) {
	return nil, status.Error(codes.Unimplemented, codes.Unimplemented.String())
}

func (s *todoService) UpdateTodo(ctx context.Context, req *todov1.UpdateTodoRequest) (*todov1.Unit, error) {
	return nil, status.Error(codes.Unimplemented, codes.Unimplemented.String())
}

func (s *todoService) DeleteTodo(ctx context.Context, req *todov1.DeleteTodoRequest) (*todov1.Unit, error) {
	return nil, status.Error(codes.Unimplemented, codes.Unimplemented.String())
}

func (s *todoService) SubscribeEvent(*todov1.SubscribeEventRequest, todov1.TodoService_SubscribeEventServer) error {
	return status.Error(codes.Unimplemented, codes.Unimplemented.String())
}
