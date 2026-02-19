package controller_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rafaeldepontes/go-predict/internal/prediction/controller"
	textModel "github.com/rafaeldepontes/go-predict/internal/text/model"
)

// mockService implements prediction.Service
type mockService struct {
	predictFunc     func(ctx context.Context, text *textModel.TextReq) (string, error)
	testPredictFunc func(ctx context.Context, text *textModel.TextReq) (string, error)
}

func (m *mockService) Predict(ctx context.Context, text *textModel.TextReq) (string, error) {
	if m.predictFunc != nil {
		return m.predictFunc(ctx, text)
	}
	return "", nil
}

func (m *mockService) TestPredict(ctx context.Context, text *textModel.TextReq) (string, error) {
	if m.testPredictFunc != nil {
		return m.testPredictFunc(ctx, text)
	}
	return "", nil
}

func TestController_Predict(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    any
		mockResponse   string
		mockError      error
		wantStatusCode int
		wantData       string
	}{
		{
			name: "successful prediction",
			requestBody: textModel.TextReq{
				Body:     "Feature A",
				TeamSize: 3,
				Level:    "Senior",
				Stack:    "Go",
			},
			mockResponse:   "Predicted value",
			mockError:      nil,
			wantStatusCode: http.StatusCreated,
			wantData:       "Predicted value",
		},
		{
			name:           "invalid JSON body",
			requestBody:    "invalid json",
			mockResponse:   "",
			mockError:      nil,
			wantStatusCode: http.StatusInternalServerError,
			wantData:       "",
		},
		{
			name: "service returns error",
			requestBody: textModel.TextReq{
				Body:     "Feature B",
				TeamSize: 2,
				Level:    "Junior",
				Stack:    "Go",
			},
			mockResponse:   "",
			mockError:      errors.New("service error"),
			wantStatusCode: http.StatusBadRequest,
			wantData:       "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Encode request body as JSON
			var reqBody bytes.Buffer
			if b, ok := tt.requestBody.(textModel.TextReq); ok {
				json.NewEncoder(&reqBody).Encode(b)
			} else if s, ok := tt.requestBody.(string); ok {
				reqBody.WriteString(s)
			}

			req := httptest.NewRequest(http.MethodPost, "/predict", &reqBody)
			w := httptest.NewRecorder()

			ctrl := &controller.PredCont{
				Service: &mockService{
					predictFunc: func(ctx context.Context, text *textModel.TextReq) (string, error) {
						return tt.mockResponse, tt.mockError
					},
				},
			}

			ctrl.Predict(w, req)

			res := w.Result()
			defer res.Body.Close()

			if res.StatusCode != tt.wantStatusCode {
				t.Errorf("StatusCode = %d, want %d", res.StatusCode, tt.wantStatusCode)
			}

			if tt.wantData != "" {
				var respBody map[string]string
				json.NewDecoder(res.Body).Decode(&respBody)
				if respBody["data"] != tt.wantData {
					t.Errorf("Response data = %v, want %v", respBody["data"], tt.wantData)
				}
			}
		})
	}
}
