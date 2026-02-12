package model

import "github.com/rafaeldepontes/go-predict/internal/prediction"

type Application struct {
	PredictionController prediction.Controller
}
