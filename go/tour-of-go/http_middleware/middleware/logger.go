package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("middleware logger")
		start := time.Now()
		// defer func() {
		// 	if errClose := r.Body.Close(); errClose != nil {
		// 		log.Println("failed to close body, should never happen")
		// 	}
		// }()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
	})
}
