package counting

import (
	"fmt"
	"testing"
)

func TestCountingSortEmptySlice(t *testing.T) {
	a := []int{}
	CountingSort(a)
	fmt.Println(a)
}

func TestCountingSortOneElement(t *testing.T) {
	a := []int{1}
	CountingSort(a)
	fmt.Println(a)
}

func TestCountingSort(t *testing.T) {
	a := []int{8, 5, 2, 6, 9, 3, 1, 4, 0, 7}
	CountingSort(a)
	fmt.Println(a)
}

func TestCountingSort2EmptySlice(t *testing.T) {
	a := []int{}
	CountingSort2(a)
	fmt.Println(a)
}

func TestCountingSort2OneElement(t *testing.T) {
	a := []int{1}
	CountingSort2(a)
	fmt.Println(a)
}

func TestCountingSort2(t *testing.T) {
	a := []int{8, 5, 2, 6, 9, 3, 1, 4, 0, 7}
	CountingSort2(a)
	fmt.Println(a)
}
