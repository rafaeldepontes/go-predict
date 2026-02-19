package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-chi/chi/v5"
	appModel "github.com/rafaeldepontes/go-predict/internal/application/model"
	"github.com/rafaeldepontes/go-predict/internal/handler"
)

type mockMiddleware struct{}

func (m *mockMiddleware) GlobalRateLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-RateLimit-Test", "true")
		next.ServeHTTP(w, r)
	})
}

type mockPredictionController struct{}

func (m *mockPredictionController) Predict(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"result":"ok"}`))
}

func setupRouter() *chi.Mux {
	os.Setenv("FRONTEND_URL", "http://localhost:3000")

	app := &appModel.Application{
		Middleware:           &mockMiddleware{},
		PredictionController: &mockPredictionController{},
	}

	r := chi.NewRouter()
	handler.ConfigHandler(r, app)

	return r
}

func TestPredictRoute(t *testing.T) {
	r := setupRouter()

	req := httptest.NewRequest(http.MethodPost, "/api/v1/predict", nil)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", rec.Code)
	}

	if rec.Header().Get("X-RateLimit-Test") != "true" {
		t.Error("rate limit middleware not executed")
	}
}

func TestSenioritiesRoute(t *testing.T) {
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/seniorities", nil)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", rec.Code)
	}

	if rec.Header().Get("Content-Type") != "application/json" {
		t.Error("expected application/json content type")
	}

	var body interface{}
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Error("invalid JSON response")
	}
}

func TestStacksRoute(t *testing.T) {
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/stacks", nil)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", rec.Code)
	}
}

func TestCORS(t *testing.T) {
	r := setupRouter()

	req := httptest.NewRequest(http.MethodOptions, "/api/v1/predict", nil)
	req.Header.Set("Origin", "http://localhost:3000")

	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if rec.Header().Get("Access-Control-Allow-Origin") == "" {
		t.Error("CORS headers not set")
	}
}
