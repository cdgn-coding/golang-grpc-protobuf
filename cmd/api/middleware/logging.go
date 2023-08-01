package middleware

import (
	"context"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type interceptorLogger struct {
	logger *zap.Logger
}

func newInterceptorLogger(logger *zap.Logger) *interceptorLogger {
	return &interceptorLogger{logger: logger}
}

func (i interceptorLogger) Log(ctx context.Context, level logging.Level, msg string, fields ...any) {
	f := make([]zap.Field, len(fields)/2)
	for i := 0; i < len(fields); i += 2 {
		f[i/2] = zap.Any(fields[i].(string), fields[i+1])
	}

	switch level {
	case logging.LevelDebug:
		i.logger.Debug(msg, f...)
	case logging.LevelInfo:
		i.logger.Info(msg, f...)
	case logging.LevelWarn:
		i.logger.Warn(msg, f...)
	case logging.LevelError:
		i.logger.Error(msg, f...)
	}
}

func NewLoggerInterceptor(logger *zap.Logger) grpc.UnaryServerInterceptor {
	interceptorLogger := newInterceptorLogger(logger)
	return logging.UnaryServerInterceptor(interceptorLogger)
}
