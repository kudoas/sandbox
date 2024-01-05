package middleware

import (
	"encoding/json"
	"net/http"
)

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
	errorMessage := `{"message": "forbidden"}`
	b, _ := json.Marshal(errorMessage)
	w.Write(b)
}
