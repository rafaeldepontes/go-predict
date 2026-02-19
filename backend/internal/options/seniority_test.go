package options_test

import (
	"testing"

	"github.com/rafaeldepontes/go-predict/internal/options"
)

func TestGetSeniorities(t *testing.T) {
	seniorities := options.GetSeniorities()

	if len(seniorities) != 4 {
		t.Fatalf("expected 4 seniorities, got %d", len(seniorities))
	}

	if seniorities[0].ID != 1 || seniorities[0].Data != "Junior" {
		t.Errorf("unexpected first seniority: %+v", seniorities[0])
	}

	last := seniorities[len(seniorities)-1]
	if last.ID != 4 || last.Data != "Expert" {
		t.Errorf("unexpected last seniority: %+v", last)
	}

	seen := make(map[int]bool)
	for _, s := range seniorities {
		if seen[s.ID] {
			t.Errorf("duplicate ID found: %d", s.ID)
		}
		seen[s.ID] = true
	}
}
