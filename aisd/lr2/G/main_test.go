package main

import "testing"

func TestTask(t *testing.T) {
	arr := [][]int{
		{3, 2, 9, 1, 1, 4, 4, 6},
		{1, 2, 3, 4, 5},
		{5, 4, 3, 2, 1},
		{13, 8, 12, 9},
	}
	expected := []int{
		50,
		9,
		9,
		96,
	}

	for i := 0; i < len(arr); i++ {
		actual, err := getSquare(len(arr[i]), arr[i])
		if err != nil {
			t.Errorf("Should not produce an error")
		}
		if expected[i] != actual {
			t.Errorf("Result incorrect, got: %d, want %d", actual, expected[i])
		}
	}

}
