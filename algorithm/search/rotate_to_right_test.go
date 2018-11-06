package search

import (
	"testing"
)

func TestRotateToRight(t *testing.T) {
	a := []int{0, 1, 16, 24, 35, 47, 59, 62, 73, 88, 99}
	if RotateToRight(a, len(a)) != 0 {
		t.Fatal("rotate 0 time")
	}
	a = []int{24, 35, 47, 59, 62, 73, 88, 99, 0, 1, 16}
	if RotateToRight(a, len(a)) != 8 {
		t.Fatal("rotate 8 times")
	}
	a = []int{99, 0, 1, 16, 24, 35, 47, 59, 62, 73, 88}
	if RotateToRight(a, len(a)) != 1 {
		t.Fatal("rotate one time")
	}
}
