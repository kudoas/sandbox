package middleware_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kudoas/enjoy-middleware/middleware"
)

func TestAuthentication(t *testing.T) {
	testCase := []struct {
		name               string
		host               string
		expectedStatusCode int
	}{
		{
			name:               "Authorized / valid host header",
			host:               "admin.internal",
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "UnAuthorized / invalid host header",
			host:               "public.example",
			expectedStatusCode: http.StatusForbidden,
		},
	}

	testHandler := setupTestHandler(t)

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			req.Host = tc.host
			rec := httptest.NewRecorder()

			testHandler.ServeHTTP(rec, req)

			switch tc.host {
			case "public.example":
				if rec.Code == http.StatusOK {
					t.Errorf("Expected status code %d, got %d", http.StatusOK, rec.Code)
				}
				var resBody middleware.ErrorResponse
				json.NewDecoder(bytes.NewReader(rec.Body.Bytes())).Decode(&resBody)
				if resBody.Message != "forbidden" {
					t.Errorf("Expected response body %s, got %s", "forbidden", resBody.Message)
				}
			case "admin.internal":
				if rec.Code == http.StatusForbidden {
					t.Errorf("Expected status code %d, got %d", http.StatusForbidden, rec.Code)
				}
			}
		})
	}
}

func setupTestHandler(t *testing.T) http.Handler {
	t.Helper()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	return middleware.Authentication(handler)
}
