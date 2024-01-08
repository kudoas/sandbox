// internal server error

package response

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func InternalServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	errorMessage := ErrorResponse{
		Message: "internal server error",
	}
	b, _ := json.Marshal(errorMessage)
	_, _ = w.Write(b)
}
