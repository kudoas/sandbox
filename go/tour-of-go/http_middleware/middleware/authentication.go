package middleware

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isPublicHost(r.Host) {
			writeForbiddenResponse(w)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func isPublicHost(host string) bool {
	return host == "public.example"
}

func writeForbiddenResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusForbidden)
	errorMessage := ErrorResponse{
		Message: "forbidden",
	}
	b, _ := json.Marshal(errorMessage)
	_, _ = w.Write(b)
}
