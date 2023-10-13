package main

import "testing"

func TestTask(t *testing.T) {
	arr := [][][]int{
		{
			{3, 2, 1, 5, 2},
			{4, 8, 3, 1, 2},
		},
	}
	k := []int{
		25,
		24,
		19,
		18,
		17,
		16,
		15,
		14,
		13,
		12,
		11,
		10,
	}
	expected := []int{
		11,
	}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(k); j++ {
			actual := getKRatio(k[j], len(arr[i][0]), arr[i])
			if expected[i] != actual {
				t.Errorf("Result incorrect, got: %d, want %d", actual, expected[i])
			}
		}
	}

}
