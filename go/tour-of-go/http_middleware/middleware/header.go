package middleware

import (
	"log"
	"net/http"
)

func Header(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("middleware header")
		w.Header().Add("X-custom-header", "my-value")
		next.ServeHTTP(w, r)
	})
}
