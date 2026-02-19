package limit_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rafaeldepontes/go-predict/internal/rate/limit"
)

// TestGlobalRateLimit checks that the rate limiter enforces limits correctly
func TestGlobalRateLimit(t *testing.T) {
	mw := limit.NewMiddleware()

	// Simple handler that returns 200 OK
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	limitedHandler := mw.GlobalRateLimit(handler)

	// First two requests should succeed (burst of 2 allowed)
	for i := 0; i < 2; i++ {
		req := httptest.NewRequest("GET", "/", nil)
		rec := httptest.NewRecorder()

		limitedHandler.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("expected status 200, got %d on request %d", rec.Code, i+1)
		}
	}

	// Third request should be rejected (rate limit exceeded)
	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	limitedHandler.ServeHTTP(rec, req)

	if rec.Code != http.StatusTooManyRequests {
		t.Errorf("expected status 429, got %d on request 3", rec.Code)
	}
}
