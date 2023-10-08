package main

import "testing"

func TestTask(t *testing.T) {
	n := 8
	arr := []int{3, 2, 1, 4, 9, 1, 4, 6}
	expected := 50
	actual, err := getSquare(n, arr)

	if err != nil {
		t.Errorf("Should not produce an error")
	}
	if expected != actual {
		t.Errorf("Result incorrect, got: %d, want %d", actual, expected)
	}
}
