package log

import (
	"fmt"

	"go.uber.org/zap"
)

func init() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(fmt.Sprintf("log init error: %v", err))
	}
	zap.ReplaceGlobals(logger)
}
