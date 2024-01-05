package middleware

import (
	"net/http"
)

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if host := r.Host; host == "public.example" {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte(`{"message": "forbidden"}`))
			return
		}
		next.ServeHTTP(w, r)
	})
}
