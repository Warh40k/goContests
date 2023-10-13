package main

import "testing"

func TestMendeleev(t *testing.T) {
	input := []string{
		"H4ZO2",
		"N(OY3)2",
		"Z3(O(N2O3)2)3",
	}
	expected := []string{
		"H4O2Z",
		"NO2Y6",
		"N12O21Z3",
	}

	for i := 0; i < len(input); i++ {
		actual := parseFormula(input[i])

		if expected[i] != actual {
			t.Errorf("Result incorrect, got: %s, want %s", actual, expected[i])
		}
	}
}
