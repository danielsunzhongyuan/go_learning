package radix

import (
	"fmt"
	"testing"
)

func TestRadixSortEmptySlice(t *testing.T) {
	a := []int{}
	LsdRadixSort(a, len(a))
	fmt.Println(a)
}

func TestRadixSortOneElement(t *testing.T) {
	a := []int{1}
	LsdRadixSort(a, len(a))
	fmt.Println(a)
}

func TestRadixSort(t *testing.T) {
	a := []int{20, 90, 64, 289, 998, 365, 852, 123, 789, 456}
	LsdRadixSort(a, len(a))
	fmt.Println(a)
}

func TestRadixSort2EmptySlice(t *testing.T) {
	a := []int{}
	LsdRadixSort2(a, len(a))
	fmt.Println(a)
}

func TestRadixSort2OneElement(t *testing.T) {
	a := []int{1}
	LsdRadixSort2(a, len(a))
	fmt.Println(a)
}

func TestRadixSort2(t *testing.T) {
	a := []int{20, 90, 64, 289, 998, 365, 852, 123, 789, 456}
	LsdRadixSort2(a, len(a))
	fmt.Println(a)
}
