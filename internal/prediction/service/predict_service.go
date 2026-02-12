package service

import "github.com/rafaeldepontes/go-predict/internal/prediction"

type svc struct{}

func NewService() prediction.Service {
	return &svc{}
}

func (s *svc) Predict(body string) (string, error) {
	return "", nil
}
