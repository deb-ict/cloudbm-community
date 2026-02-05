package logging

import (
	"log/slog"
	"net/http"
)

type LoggingMiddleware struct {
}

func NewLoggingMiddleware() *LoggingMiddleware {
	return &LoggingMiddleware{}
}

func (m *LoggingMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		logger := GetLoggerFromContext(ctx)
		logger = logger.With(
			slog.String("http.request.method", r.Method),
			slog.String("http.request.host", r.Host),
			slog.String("http.request.url", r.URL.Path),
		)

		ctx = WithLoggerInContext(ctx, logger)
		next.ServeHTTP(w, r.WithContext(ctx))

		if r.Response.StatusCode >= 200 && r.Response.StatusCode < 400 {
			logger.Info("HTTP request completed",
				slog.Int("http.response.status_code", r.Response.StatusCode),
			)
		} else {
			logger.Error("HTTP request completed with error",
				slog.Int("http.response.status_code", r.Response.StatusCode),
			)
		}
	})
}
