package service

import (
	"bytes"
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

	promptBody := prompt.Get()
	idx := bytes.IndexByte(promptBody, '\n')

	modelCtx, promptF := promptBody[:idx+1], promptBody[idx+1:]

	msg := fmt.Sprintf(
		string(promptF),
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

	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Println("[ERROR] Could not connect to gemini:", err)
		return "", errors.New("Something went really bad...")
	}

	countResp, err := client.Models.CountTokens(ctx, os.Getenv("MODEL"),
		genai.Text(sb.String()),
		nil,
	)
	if err != nil {
		log.Println("[ERROR] Could not count the amount of tokens:", err)
		return "", errors.New("Something went really bad...")
	}
	log.Println("[INFO] Tokens spent:", countResp.TotalTokens)

	config := &genai.GenerateContentConfig{
		SystemInstruction: genai.NewContentFromText(string(modelCtx), genai.RoleModel),
	}

	result, err := client.Models.GenerateContent(
		ctx,
		os.Getenv("MODEL"),
		genai.Text(sb.String()),
		config,
	)
	if err != nil {
		log.Println("[ERROR] Could not get a response from gemini:", err)
		return "", errors.New("Something went really bad...")
	}
	return result.Text(), nil
}

func (s *svc) TestPredict(ctx context.Context, text *textModel.TextReq) (string, error) {
	if err := validateRequest(text); err != nil {
		return "", err
	}

	msg := fmt.Sprintf(
		"Test, %v, %d, %v, %v",
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

	return "Just testing... don't want to waste my tokens...", nil
}

func validateRequest(text *textModel.TextReq) error {
	if text == nil || strings.TrimSpace(text.Body) == "" {
		return errors.New("Cannot make a prediction without the features.")
	}

	if text.TeamSize <= 0 {
		return errors.New("Cannot make a prediction without the team size.")
	}

	if strings.TrimSpace(text.Stack) == "" {
		return errors.New("Cannot make a prediction without the stack.")
	}

	if text.Level == "" {
		return errors.New("Cannot make a prediction without the team seniority.")
	}

	return nil
}
