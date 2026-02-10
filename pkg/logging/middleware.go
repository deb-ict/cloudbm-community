package logging

import (
	"log/slog"
	"net/http"
)

type LoggingMiddleware struct {
}

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func NewMiddleware() *LoggingMiddleware {
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
		logWriter := &loggingResponseWriter{ResponseWriter: w}
		next.ServeHTTP(logWriter, r.WithContext(ctx))

		if logWriter.statusCode >= 200 && logWriter.statusCode < 400 {
			logger.InfoContext(ctx, "HTTP request completed",
				slog.Int("http.response.status_code", logWriter.statusCode),
			)
		} else {
			logger.ErrorContext(ctx, "HTTP request failed",
				slog.Int("http.response.status_code", logWriter.statusCode),
			)
		}
	})
}
