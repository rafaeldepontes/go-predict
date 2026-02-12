package handler

import (
	"os"

	"github.com/go-chi/cors"

	"github.com/go-chi/chi/v5"
	chiM "github.com/go-chi/chi/v5/middleware"
	appModel "github.com/rafaeldepontes/go-predict/internal/application/model"
)

func ConfigHandler(r *chi.Mux, app *appModel.Application) {
	r.Use(chiM.StripSlashes)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{os.Getenv("FRONTEND_URL")},
		AllowedMethods:   []string{"POST", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Authorization"},
		AllowCredentials: false,
		MaxAge:           300, //5 min
	}))

	r.Group(func(r chi.Router) {
		r.Post("/api/v1/predict", app.PredictionController.Predict)
	})
}
