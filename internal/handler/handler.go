package handler

import (
	"github.com/go-chi/chi/v5"
	chiM "github.com/go-chi/chi/v5/middleware"
	appModel "github.com/rafaeldepontes/go-predict/internal/application/model"
)

func ConfigHandler(r *chi.Mux, app *appModel.Application) {
	r.Use(chiM.StripSlashes)
	r.Group(func(r chi.Router) {
		r.Post("/api/v1/predict", app.PredictionController.Predict)
	})
}
