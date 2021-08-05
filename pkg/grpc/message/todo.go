package message

import (
	"github.com/iwaltgen/grpc-go-web-todo/pkg/entity"
	"google.golang.org/protobuf/types/known/timestamppb"

	todov1 "github.com/iwaltgen/grpc-go-web-todo/api/todo/v1"
)

// TodoFromProto convert *todov1.Todo to *entity.Todo
func TodoFromProto(value *todov1.Todo) *entity.Todo {
	return &entity.Todo{
		ID:          value.Id,
		Description: value.Description,
		Completed:   value.Completed,
		ModifiedAt:  value.ModifiedAt.AsTime(),
		CreatedAt:   value.CreatedAt.AsTime(),
	}
}

// TodoProto convert *entity.Todo to *todov1.Todo
func TodoProto(value *entity.Todo) *todov1.Todo {
	return &todov1.Todo{
		Id:          value.ID,
		Description: value.Description,
		Completed:   value.Completed,
		ModifiedAt:  timestamppb.New(value.ModifiedAt),
		CreatedAt:   timestamppb.New(value.CreatedAt),
	}
}

// TodoProtoList convert []*entity.Todo to []*todov1.Todo
func TodoProtoList(list []*entity.Todo) (ret []*todov1.Todo) {
	for _, v := range list {
		ret = append(ret, TodoProto(v))
	}
	return ret
}
