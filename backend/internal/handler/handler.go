package handler

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/go-chi/cors"

	"github.com/go-chi/chi/v5"
	chiM "github.com/go-chi/chi/v5/middleware"
	appModel "github.com/rafaeldepontes/go-predict/internal/application/model"
	"github.com/rafaeldepontes/go-predict/internal/options"
)

func ConfigHandler(r *chi.Mux, app *appModel.Application) {
	r.Use(chiM.StripSlashes)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{os.Getenv("FRONTEND_URL")},
		AllowedMethods:   []string{"POST", "GET", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Authorization"},
		AllowCredentials: false,
		MaxAge:           300, //5 min
	}))

	r.Group(func(r chi.Router) {
		r.Use(app.Middleware.GlobalRateLimit)

		r.Post("/api/v1/predict", app.PredictionController.Predict)

		r.Get("/api/v1/seniorities", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(options.GetSeniorities())
		})

		r.Get("/api/v1/stacks", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(options.GetStacks())
		})
	})
}
