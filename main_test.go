package main

import (
	"testing"
)

func TestGenerateGreeting(t *testing.T) {
	testCases := []struct {
		name          string
		expectedGreet string
	}{
		{"", "Hello, World!"},
		{"Alice", "Hello, Alice!"},
		{"@#$%^&*()", "Hello, @#$%^&*()!"},
		{"  John Doe   ", "Hello, John Doe!"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			greet := GenerateGreeting(tc.name)
			if greet != tc.expectedGreet {
				t.Errorf("Expected greeting: %s, got: %s", tc.expectedGreet, greet)
			}
		})
	}
}
