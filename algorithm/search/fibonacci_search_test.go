package search

import "testing"

func TestFibonacciSearch(t *testing.T) {
	a := []int{0, 1, 16, 24, 35, 47, 59, 62, 73, 88, 99}
	FibonacciSearch(a, 10)
}
