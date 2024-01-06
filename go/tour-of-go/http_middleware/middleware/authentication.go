package middleware

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isPublicHost(r.Host) {
			if err := writeForbiddenResponse(w); err != nil {
				// TODO: internal server error
				log.Fatalf("failed to write response: %v", err)
			}
			return
		}
		next.ServeHTTP(w, r)
	})
}

func isPublicHost(host string) bool {
	return host == "public.example"
}

func writeForbiddenResponse(w http.ResponseWriter) error {
	w.WriteHeader(http.StatusForbidden)
	errorMessage := ErrorResponse{
		Message: "forbidden",
	}
	b, err := json.Marshal(errorMessage)
	if err != nil {
		return err
	}
	if _, err := w.Write(b); err != nil {
		return err
	}
	return nil
}
