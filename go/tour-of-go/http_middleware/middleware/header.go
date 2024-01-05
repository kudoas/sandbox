package middleware

import "net/http"

func Header(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("X-custom-header", "my-value")
		next.ServeHTTP(w, r)
	}
}
