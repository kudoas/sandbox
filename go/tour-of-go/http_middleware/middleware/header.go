package middleware

import (
	"log"
	"net/http"
)

func Header(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("middleware header")
		w.Header().Set("content-type", "application/json")
		w.Header().Set("x-custom-value", "my-value")
		next.ServeHTTP(w, r)
	})
}
