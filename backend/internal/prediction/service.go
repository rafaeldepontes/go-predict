package prediction

import (
	"context"

	textModel "github.com/rafaeldepontes/go-predict/internal/text/model"
)

type Service interface {
	Predict(ctx context.Context, text *textModel.TextReq) (string, error)
}
