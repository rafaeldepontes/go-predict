package controller

import (
	"net/http"

	"github.com/rafaeldepontes/go-predict/internal/prediction"
	service "github.com/rafaeldepontes/go-predict/internal/prediction/service"
)

type predCont struct {
	Service prediction.Service
}

func NewController() prediction.Controller {
	return &predCont{
		Service: service.NewService(),
	}
}

func (c *predCont) Predict(w http.ResponseWriter, r *http.Request) {
	_, err := c.Service.Predict("")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
