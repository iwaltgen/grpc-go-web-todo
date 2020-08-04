package log

import (
	"go.uber.org/zap"
)

// Logger zap.logger
type Logger = zap.Logger

// L zap.L.Named
func L(name string) *Logger {
	return zap.L().Named(name)
}
