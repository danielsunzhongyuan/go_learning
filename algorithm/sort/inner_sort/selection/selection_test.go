package selection

import (
	"fmt"
	"testing"
)

func TestSelectionSortEmptySlice(t *testing.T) {
	a := []int{}
	SelectionSort(a)
	fmt.Println(a)
}

func TestSelectionSortOneElement(t *testing.T) {
	a := []int{-1}
	SelectionSort(a)
	fmt.Println(a)
}

func TestSelectionSort(t *testing.T) {
	a := []int{8, 5, 2, 6, 9, 3, 1, 4, 0, 7}
	SelectionSort(a)
	fmt.Println(a)
}

// SelectionSort2
func TestSelectionSort2EmptySlice(t *testing.T) {
	a := []int{}
	SelectionSort(a)
	fmt.Println(a)
}

func TestSelectionSort2OneElement(t *testing.T) {
	a := []int{-1}
	SelectionSort(a)
	fmt.Println(a)
}

func TestSelectionSort2(t *testing.T) {
	a := []int{8, 5, 2, 6, 9, 3, 1, 4, 0, 7}
	SelectionSort(a)
	fmt.Println(a)
}
