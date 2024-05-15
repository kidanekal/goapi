package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kidanekal/goapi/api"
	"golang.org/x/net/context"
)

// TestHealthHandler tests the health handler
func TestHealthHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	ctx := context.Background()

	api.HealthHandler(ctx, recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, recorder.Code)
	}
}
