package prompt_test

import (
	"os"
	"testing"

	"github.com/rafaeldepontes/go-predict/internal/prompt"
)

func TestGet(t *testing.T) {
	tests := []struct {
		name        string
		fileContent string
		wantErr     bool
	}{
		{
			name:        "file exists",
			fileContent: "This is a test prompt",
			wantErr:     false,
		},
		{
			name:        "file missing",
			fileContent: "",
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup temporary file if content provided
			var tmpFileName string
			if tt.fileContent != "" {
				f, err := os.CreateTemp("", "prompt_test_*.txt")
				if err != nil {
					t.Fatalf("Failed to create temp file: %v", err)
				}
				defer os.Remove(f.Name())
				_, _ = f.WriteString(tt.fileContent)
				f.Close()
				tmpFileName = f.Name()
			} else {
				tmpFileName = "/non/existent/file/path"
			}

			// Set environment variable
			os.Setenv("FILE_PATH", tmpFileName)
			defer os.Unsetenv("FILE_PATH")

			prompt.ResetForTest()

			got := prompt.Get()
			if tt.wantErr {
				if *got != "" {
					t.Errorf("Expected empty prompt for missing file, got: %v", *got)
				}
			} else {
				if *got != tt.fileContent {
					t.Errorf("Get() = %v, want %v", *got, tt.fileContent)
				}
			}
		})
	}
}
