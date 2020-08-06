package repository

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/iwaltgen/grpc-go-web-todo/pkg/entity"
)

func TestTodoReadOperation(t *testing.T) {
	ctx := context.Background()
	repo := newTodo()

	ret, err := repo.FindAll(ctx)
	assert.NoError(t, err)
	assert.Empty(t, ret)
}

func TestTodoCreateOperation(t *testing.T) {
	ctx := context.Background()
	repo := newTodo()

	now := time.Now()
	todo := &entity.Todo{
		ID:          "test1",
		Description: "test 1 description",
		ModifiedAt:  now,
		CreatedAt:   now,
	}
	err := repo.Create(ctx, todo)
	assert.NoError(t, err)

	ret, err := repo.FindAll(ctx)
	assert.NoError(t, err)
	assert.Len(t, ret, 1)
}
