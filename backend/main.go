package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	appModel "github.com/rafaeldepontes/go-predict/internal/application/model"
	"github.com/rafaeldepontes/go-predict/internal/handler"
	predCont "github.com/rafaeldepontes/go-predict/internal/prediction/controller"
	"github.com/rafaeldepontes/go-predict/internal/rate/limit"
	"github.com/rafaeldepontes/go-predict/internal/tool"
)

var app *appModel.Application

func init() {
	env := ".env"
	tool.ChecksEnv(&env)
	if err := godotenv.Load(env); err != nil {
		log.Fatalln("[ERROR] Could not load env:", err)
	}

	app = newApplication()
}

func main() {
	r := chi.NewRouter()
	handler.ConfigHandler(r, app)
	port := os.Getenv("BACKEND_PORT")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	go func() {
		log.Println("[INFO] Application running...")
		log.Fatalln(http.ListenAndServe(":"+port, r))
	}()

	<-sigChan
	log.Println("[INFO] Trying to shut down gracefully")
	//...
	log.Println("[INFO] Shut down successfully")
}

func newApplication() *appModel.Application {
	return &appModel.Application{
		Middleware:           limit.NewMiddleware(),
		PredictionController: predCont.NewController(),
	}
}
