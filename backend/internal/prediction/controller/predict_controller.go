package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/rafaeldepontes/go-predict/internal/prediction"
	service "github.com/rafaeldepontes/go-predict/internal/prediction/service"
	textModel "github.com/rafaeldepontes/go-predict/internal/text/model"
)

type PredCont struct {
	Service prediction.Service
}

func NewController() prediction.Controller {
	return &PredCont{
		Service: service.NewService(),
	}
}

func (c *PredCont) Predict(w http.ResponseWriter, r *http.Request) {
	var text textModel.TextReq
	if err := json.NewDecoder(r.Body).Decode(&text); err != nil {
		log.Println("[ERROR] Could not decode the request body:", err)
		http.Error(w, "Something went really bad...", http.StatusInternalServerError)
		return
	}

	result, err := c.Service.Predict(r.Context(), &text)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	resp := map[string]string{"data": result}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
