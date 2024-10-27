package main

import "testing"

func TestValidateInput(t *testing.T) {
	type Test struct {
		input    string
		expected bool
	}
	args := []Test{
		{"2+2", true},
		{"2*2+2", true},
		{"2+(2+2)", true},
		{"2+(2+(2+(2))", false},
		{"1/2", true},
		{"a2+2", false},
		{"2:2", false},
		{"2+2+", false},
		{"", false},
	}
	for _, input := range args {
		got := ValidateInput(input.input)
		expected := input.expected
		if got != expected {
			t.Errorf("ValidateInput(%s) == %t, want %t", input.input, got, expected)
		}
	}
}
