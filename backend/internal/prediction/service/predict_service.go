package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/rafaeldepontes/go-predict/internal/prediction"
	"github.com/rafaeldepontes/go-predict/internal/prompt"
	textModel "github.com/rafaeldepontes/go-predict/internal/text/model"

	"google.golang.org/genai"
)

type svc struct{}

func NewService() prediction.Service {
	return &svc{}
}

func (s *svc) Predict(ctx context.Context, text *textModel.TextReq) (string, error) {
	if err := validateRequest(text); err != nil {
		return "", err
	}

	msg := fmt.Sprintf(
		*prompt.Get(),
		text.Body,
		text.TeamSize,
		text.Level,
		text.Stack,
	)

	var sb strings.Builder
	if _, err := sb.WriteString(msg); err != nil {
		log.Println("[ERROR] Could not create the prompt:", err)
		return "", errors.New("Something went really bad...")
	}

	// Create the client config
	// ...

	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Println("[ERROR] Could not connect to gemini:", err)
		return "", errors.New("Something went really bad...")
	}

	result, err := client.Models.GenerateContent(
		ctx,
		os.Getenv("MODEL"),
		genai.Text(sb.String()),
		nil,
	)
	if err != nil {
		log.Println("[ERROR] Could not get a response from gemini:", err)
		return "", errors.New("Something went really bad...")
	}
	return result.Text(), nil
}

func validateRequest(text *textModel.TextReq) error {
	if text == nil || strings.TrimSpace(text.Body) == "" {
		return errors.New("Cannot send an empty message, please type something...")
	}

	if text.TeamSize <= 0 {
		return errors.New("Cannot send a message with no team, define a team size...")
	}

	if strings.TrimSpace(text.Stack) == "" {
		return errors.New("Cannot send a message with no stack, please choose a stack...")
	}

	return nil
}
