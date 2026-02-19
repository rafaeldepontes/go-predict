package model

import (
	"github.com/rafaeldepontes/go-predict/internal/prediction"
	"github.com/rafaeldepontes/go-predict/internal/rate/limit"
)

type Application struct {
	PredictionController prediction.Controller
	Middleware           *limit.Middleware
}
