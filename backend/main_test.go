package main_test

import (
	"log"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	if code > 0 {
		log.Fatalf("Expected 0, got %d", code)
	}
}
