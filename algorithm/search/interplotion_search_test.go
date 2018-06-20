package search

import (
	"testing"
)

func TestInterplotionSearch(t *testing.T) {
	a := []int{0, 1, 16, 24, 35, 47, 59, 62, 73, 88, 99}
	if InterplotionSearch(a, 62) != 7 {
		t.Fatal("search 62, should return 7")
	}
	if InterplotionSearch(a, 0) != 0 {
		t.Fatal("search 0, should return 0")
	}
	if InterplotionSearch(a, 99) != 10 {
		t.Fatal("search 99, should return 10")
	}
	if InterplotionSearch(a, -1) != -1 {
		t.Fatal("search -1, should return -1")
	}
	if InterplotionSearch(a, 63) != -1 {
		t.Fatal("search 63, should return -1")
	}
	if InterplotionSearch(a, 100) != -1 {
		t.Fatal("search 100, should return -1")
	}
}
