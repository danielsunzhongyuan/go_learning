package heap

import (
	"fmt"
	"testing"
)

func TestHeapSortEmptySlice(t *testing.T) {
	a := []int{}
	HeapSort(a, len(a))
	fmt.Println(a)
}

func TestHeapSortOneElement(t *testing.T) {
	a := []int{1}
	HeapSort(a, len(a))
	fmt.Println(a)
}

func TestHeapSort(t *testing.T) {
	a := []int{8, 5, 2, 6, 9, 3, 1, 4, 0, 7}
	HeapSort(a, len(a))
	fmt.Println(a)
}

func TestHeapSort2EmptySlice(t *testing.T) {
	a := []int{}
	HeapSort2(a, len(a))
	fmt.Println(a)
}

func TestHeapSort2OneElement(t *testing.T) {
	a := []int{1}
	HeapSort2(a, len(a))
	fmt.Println(a)
}

func TestHeapSort2(t *testing.T) {
	a := []int{8, 5, 2, 6, 9, 3, 1, 4, 0, 7}
	HeapSort2(a, len(a))
	fmt.Println(a)
}
