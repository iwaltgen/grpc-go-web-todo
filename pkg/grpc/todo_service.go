package grpc

import (
	"context"
	"sync"

	"github.com/gogo/protobuf/types"
	"github.com/gogo/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"github.com/iwaltgen/grpc-go-web-todo/pkg/grpc/message"
	"github.com/iwaltgen/grpc-go-web-todo/pkg/log"
	"github.com/iwaltgen/grpc-go-web-todo/pkg/usecase"

	todov1 "github.com/iwaltgen/grpc-go-web-todo/api/todo/v1"
)

type todoService struct {
	logger      *log.Logger
	todoUsecase *usecase.Todo
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

func newTodoService(todoUsecase *usecase.Todo) todov1.TodoServiceServer {
	return &todoService{
		logger:      log.L("grpc.todo"),
		todoUsecase: todoUsecase,
	}
}

func (s *todoService) ListTodos(ctx context.Context, req *todov1.ListTodosRequest) (*todov1.ListTodosResponse, error) {
	ret, err := s.todoUsecase.ListTodos(ctx)
	if err != nil {
		return nil, err
	}

	return &todov1.ListTodosResponse{
		Todos: message.TodoProtoList(ret),
	}, nil
}

func (s *todoService) CreateTodo(ctx context.Context, req *todov1.CreateTodoRequest) (*types.Empty, error) {
	err := s.todoUsecase.CreateTodo(ctx, message.TodoFromProto(req.Todo))
	if err != nil {
		return nil, err
	}

	return &types.Empty{}, nil
}

func (s *todoService) UpdateTodo(ctx context.Context, req *todov1.UpdateTodoRequest) (*types.Empty, error) {
	err := s.todoUsecase.UpdateTodo(ctx, message.TodoFromProto(req.Todo))
	if err != nil {
		return nil, err
	}

	return &types.Empty{}, nil
}

func (s *todoService) DeleteTodo(ctx context.Context, req *todov1.DeleteTodoRequest) (*types.Empty, error) {
	err := s.todoUsecase.DeleteTodo(ctx, req.TodoId)
	if err != nil {
		return nil, err
	}

	return &types.Empty{}, nil
}

func (s *todoService) SubscribeEvent(*todov1.SubscribeEventRequest, todov1.TodoService_SubscribeEventServer) error {
	return status.Error(codes.Unimplemented, codes.Unimplemented.String())
}
