package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	appModel "github.com/rafaeldepontes/go-predict/internal/application/model"
	"github.com/rafaeldepontes/go-predict/internal/handler"
	predCont "github.com/rafaeldepontes/go-predict/internal/prediction/controller"
)

var app *appModel.Application

func init() {
	env := ".env"
	// TODO: add the tool package.
	if err := godotenv.Load(env); err != nil {
		log.Fatalln("[ERROR] Could not load env:", err)
	}

	app = newApplication()
}

func main() {
	r := chi.NewRouter()
	handler.ConfigHandler(r, app)

	log.Println("Application running...")
	log.Fatalln(http.ListenAndServe(":8080", r))
}

func newApplication() *appModel.Application {
	return &appModel.Application{
		PredictionController: predCont.NewController(),
	}
}
