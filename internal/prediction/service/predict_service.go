package service

import (
	"context"
	"errors"
	"log"
	"strings"

	"github.com/rafaeldepontes/go-predict/internal/prediction"
	textModel "github.com/rafaeldepontes/go-predict/internal/text/model"

	"google.golang.org/genai"
)

var (
	prompt string
	model  string = "gemini-3-flash-preview"
)

type svc struct{}

func NewService() prediction.Service {
	return &svc{}
}

func (s *svc) Predict(ctx context.Context, text *textModel.TextReq) (string, error) {
	if text == nil || text.Body == "" {
		return "", errors.New("Cannot send a empty message...")
	}

	var sb strings.Builder

	// Build prompt for Gemini
	sb.WriteString("Based on this entry: \"")
	sb.WriteString(text.Body)
	sb.WriteString("\"")
	sb.WriteString(prompt)

	// Create the client config
	// ...

	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Println("[ERROR Could not connect to gemini:", err)
		return "", errors.New("Something went really bad...")
	}

	result, err := client.Models.GenerateContent(
		ctx,
		model,
		genai.Text(sb.String()),
		nil,
	)
	if err != nil {
		log.Println("[ERROR] Could not get a response from gemini:", err)
		return "", errors.New("Something went really bad...")
	}
	return result.Text(), nil
}
