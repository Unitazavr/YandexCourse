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
func TestCalc(t *testing.T) {
	type Test struct {
		input    string
		expected float64
	}
	args := []Test{
		{"2+2", 4},
		{"2*2+2", 6},
		{"2+(2+2)", 6},
		{"2+(2+(2+(2))", 0},
		{"1/2", 0.5},
		{"a2+2", 0},
		{"2:2", 1}, //0 потому что деление должно быть через "/"
		{"2+2+", 0},
		{"", 0},
	}
	for _, input := range args {
		got, _ := Calc(input.input)
		expected := input.expected
		if got != expected {
			t.Errorf("Calc(%s) == %g, want %g", input.input, got, expected)
		}
	}
}
