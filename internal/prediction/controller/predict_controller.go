package controller

import (
	"encoding/json"
	"net/http"

	"github.com/rafaeldepontes/go-predict/internal/prediction"
	service "github.com/rafaeldepontes/go-predict/internal/prediction/service"
	textModel "github.com/rafaeldepontes/go-predict/internal/text/model"
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
	var text *textModel.TextReq
	if err := json.NewDecoder(r.Body).Decode(text); err != nil {
		http.Error(w, "Something went really bad...", http.StatusInternalServerError)
		return
	}

	_, err := c.Service.Predict(r.Context(), text)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
