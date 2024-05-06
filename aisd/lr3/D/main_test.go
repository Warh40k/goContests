package main

import "testing"

func TestTour(t *testing.T) {
	ans := []string{"YES", "NO", "YES"}
	rats := []int{9, 2, 4, 8, 1}
	rews := []int{10, 2, 3, 1, 4}
	codes := []int{1, 2, 3, 4, 5}
	length := len(rats)

	qAskExpect, qTourExpect := new(Queue[int]), new(Queue[int])
	qAskExpect.push(4)
	qAskExpect.push(3)
	qAskExpect.push(2)
	qTourExpect.push(1)
	qTourExpect.push(4)
	qTourExpect.push(2)
	qTourExpect.push(5)

	answers := new(Queue[string])
	for _, val := range ans {
		answers.push(val)
	}

	ratings, rewards := make([]*Value, length), make([]*Value, length)
	for i := 0; i < length; i++ {
		city := &City{code: codes[i]}
		ratings[i] = &Value{city: city, value: rats[i]}
		rewards[i] = &Value{city: city, value: rews[i]}
	}

	qAsk, qTour := findTourPath(length, answers, rewards, ratings)

	for qAsk.size != 0 {
		got, expect := qAsk.pop(), qAskExpect.pop()
		if got != expect {
			t.Errorf("Incorrect, expected: %d, got: %d", expect, got)
		}
	}

	for qTour.size != 0 {
		got, expect := qTour.pop(), qTourExpect.pop()
		if got != expect {
			t.Errorf("Incorrect, expected: %d, got: %d", expect, got)
		}
	}
}
