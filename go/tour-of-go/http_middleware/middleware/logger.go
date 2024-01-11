package middleware

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

// prevent logger initialization every request
func NewLogger() (*zap.Logger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	return logger, nil
}

func Logger(logger *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			logger.Info("request completed",
				zap.String("method", r.Method),
				zap.String("host", r.Host),
				zap.String("url", r.URL.String()),
				zap.Duration("duration", time.Since(start)),
			)
		})
	}
}
