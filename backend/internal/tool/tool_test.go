package tool_test

import (
	"os"
	"testing"

	"github.com/rafaeldepontes/go-predict/internal/tool"
)

func TestChecksEnv(t *testing.T) {
	tests := []struct {
		name       string
		setupFiles []string // files to create before running
		input      string
		want       string
	}{
		{
			name:       "file exists",
			setupFiles: []string{"file1.env"},
			input:      "file1.env",
			want:       "file1.env",
		},
		{
			name:       "file missing, fallback exists",
			setupFiles: []string{".env.example"},
			input:      "missing.env",
			want:       ".env.example",
		},
		{
			name:       "both files missing",
			setupFiles: []string{},
			input:      "missing.env",
			want:       ".env.example",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup temporary files
			for _, f := range tt.setupFiles {
				fh, err := os.Create(f)
				if err != nil {
					t.Fatalf("Failed to create file %s: %v", f, err)
				}
				fh.Close()
				// cleanup after test
				defer os.Remove(f)
			}

			src := tt.input
			tool.ChecksEnv(&src)

			if src != tt.want {
				t.Errorf("ChecksEnv() = %v, want %v", src, tt.want)
			}
		})
	}
}
