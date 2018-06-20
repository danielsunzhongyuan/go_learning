package quick

import (
	"fmt"
	"testing"
)

func TestQuickSortEmptySlice(t *testing.T) {
	a := []int{}
	QuickSort(a, 0, len(a)-1)
	fmt.Println(a)
}

func TestQuickSortOneElement(t *testing.T) {
	a := []int{1}
	QuickSort(a, 0, len(a)-1)
	fmt.Println(a)
}

func TestQuickSort(t *testing.T) {
	a := []int{8, 5, 2, 6, 9, 3, 1, 4, 0, 7}
	QuickSort(a, 0, len(a)-1)
	fmt.Println(a)
}

func TestQuickSort2EmptySlice(t *testing.T) {
	a := []int{}
	QuickSort2(a, 0, len(a)-1)
	fmt.Println(a)
}

func TestQuickSort2OneElement(t *testing.T) {
	a := []int{1}
	QuickSort2(a, 0, len(a)-1)
	fmt.Println(a)
}

func TestQuickSort2(t *testing.T) {
	a := []int{8, 5, 2, 6, 9, 3, 1, 4, 0, 7}
	QuickSort2(a, 0, len(a)-1)
	fmt.Println(a)
}
