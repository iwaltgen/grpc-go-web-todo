package log

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func init() {
	config := zap.NewDevelopmentConfig()
	config.DisableCaller = true
	config.DisableStacktrace = true
	config.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	logger, err := config.Build()
	if err != nil {
		panic(fmt.Sprintf("log init error: %v", err))
	}
	zap.ReplaceGlobals(logger)
}
