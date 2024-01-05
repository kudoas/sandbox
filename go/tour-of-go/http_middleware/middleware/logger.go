package middleware

import (
	"log"
	"net/http"
	"time"

	"go.uber.org/zap"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		logger, err := zap.NewProduction()
		if err != nil {
			log.Fatalf("can't initialize zap logger: %v", err)
		}
		// defer logger.Sync() Probably not needed
		logger.Info("fetch URL",
			zap.Duration("time", time.Since(start)),
		)
		next.ServeHTTP(w, r)
	})
}
