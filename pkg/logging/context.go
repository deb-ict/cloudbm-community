package logging

import (
	"context"
	"log/slog"
)

type ContextKey string

const (
	LoggerContextKey ContextKey = "cbm.logger"
)

func WithLoggerInContext(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, LoggerContextKey, logger)
}

func GetLoggerFromContext(ctx context.Context) *slog.Logger {
	logger := tryGetLoggerFromContext(ctx)
	if logger == nil {
		logger = slog.Default()
	}
	return logger
}

func tryGetLoggerFromContext(ctx context.Context) *slog.Logger {
	value := ctx.Value(LoggerContextKey)
	if value == nil {
		return nil
	}
	logger, ok := value.(*slog.Logger)
	if !ok {
		return nil
	}
	return logger
}
