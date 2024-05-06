package main

import (
	"math/rand"
	"testing"
)

func TestEvaluator(t *testing.T) {
	n, m := int(3*10e5), int(3*10e5)
	var ti, f, salary int
	var shaurmen = make([]*Worker, n)
	vacantWokers := new(SalaryMinHeap)
	var incomingOrders = new(Queue)
	for i := 0; i < n; i++ {
		salary = rand.Intn(1000) + 1
		shaurmen[i] = &Worker{salary: salary}
	}
	start := rand.Intn(10e4)
	for i := 0; i < m; i++ {
		ti = start + rand.Intn(10e3)
		f = ti + rand.Intn(10e3) + 1
		incomingOrders.push(&Order{start: ti, end: f})
	}
	vacantWokers.build(shaurmen, n)
	result := evaluateSalary(m, incomingOrders, vacantWokers)

	t.Log(result)
}
