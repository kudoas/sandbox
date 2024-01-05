package middleware

import (
	"net/http"
)

func Header(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		w.Header().Set("x-custom-value", "my-value")
		next.ServeHTTP(w, r)
	})
}
