package service_test

import (
	"context"
	"testing"

	"github.com/rafaeldepontes/go-predict/internal/prediction/service"
	"github.com/rafaeldepontes/go-predict/internal/text/model"
)

func TestService_TestPredict(t *testing.T) {
	svc := service.NewService()

	ctx := context.Background()

	tests := []struct {
		name    string
		req     *model.TextReq
		wantErr bool
	}{
		{
			name: "valid request",
			req: &model.TextReq{
				Body:     "Feature A and B",
				TeamSize: 3,
				Level:    "Senior",
				Stack:    "Go, React",
			},
			wantErr: false,
		},
		{
			name: "missing body",
			req: &model.TextReq{
				Body:     "",
				TeamSize: 3,
				Level:    "Senior",
				Stack:    "Go",
			},
			wantErr: true,
		},
		{
			name: "invalid team size",
			req: &model.TextReq{
				Body:     "Feature",
				TeamSize: 0,
				Level:    "Senior",
				Stack:    "Go",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := svc.TestPredict(ctx, tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("TestPredict() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && got != "Just testing... don't want to waste my tokens..." {
				t.Errorf("TestPredict() got = %v, want %v", got, "Just testing... don't want to waste my tokens...")
			}
		})
	}
}
