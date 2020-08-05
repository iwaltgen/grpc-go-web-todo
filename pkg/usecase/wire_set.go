package usecase

import (
	"github.com/google/wire"
)

// WireSet all entities for wire inject
var WireSet = wire.NewSet(DefaultTodoService)
