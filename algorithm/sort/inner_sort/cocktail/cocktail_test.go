package cocktail

import (
	"fmt"
	"testing"
)

func TestCockTailSortEmptySlice(t *testing.T) {
	a := []int{}
	CockTailSort(a)
	fmt.Println(a)
}

func TestCockTailSortOneElement(t *testing.T) {
	a := []int{8}
	CockTailSort(a)
	fmt.Println(a)
}

func TestCockTailSort(t *testing.T) {
	a := []int{8, 5, 2, 6, 9, 3, 1, 4, 0, 7}
	CockTailSort(a)
	fmt.Println(a)
}

// CockTailSort2
func TestCockTailSort2EmptySlice(t *testing.T) {
	a := []int{}
	CockTailSort(a)
	fmt.Println(a)
}

func TestCockTailSort2OneElement(t *testing.T) {
	a := []int{8}
	CockTailSort(a)
	fmt.Println(a)
}

func TestCockTailSort2(t *testing.T) {
	a := []int{8, 5, 2, 6, 9, 3, 1, 4, 0, 7}
	CockTailSort(a)
	fmt.Println(a)
}
