package options_test

import (
	"testing"

	"github.com/rafaeldepontes/go-predict/internal/options"
)

func TestGetStacks(t *testing.T) {
	stacks := options.GetStacks()

	if len(stacks) != 50 {
		t.Fatalf("expected 50 stacks, got %d", len(stacks))
	}

	if stacks[0].ID != 1 || stacks[0].Data != "Go - Built in (net/http)" {
		t.Errorf("unexpected first stack: %+v", stacks[0])
	}

	last := stacks[len(stacks)-1]
	if last.ID != 50 || last.Data != "Serverless" {
		t.Errorf("unexpected last stack: %+v", last)
	}

	for i, stack := range stacks {
		expectedID := i + 1
		if stack.ID != expectedID {
			t.Errorf("expected ID %d, got %d", expectedID, stack.ID)
		}
	}

	for _, stack := range stacks {
		if stack.Data == "" {
			t.Errorf("stack ID %d has empty Data field", stack.ID)
		}
	}
}
