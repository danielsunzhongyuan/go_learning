package insertion

import (
	"fmt"
	"testing"
)

func TestInsertionSortEmptySlice(t *testing.T) {
	a := []int{}
	InsertionSort(a)
	fmt.Println(a)
}

func TestInsertionSortOneElement(t *testing.T) {
	a := []int{11}
	InsertionSort(a)
	fmt.Println(a)
}

func TestInsertionSort(t *testing.T) {
	fmt.Println("###")
	a := []int{8, 5, 2, 6, 9, 3, 1, 4, 0, 7}
	InsertionSort(a)
	fmt.Println(a)
}

func TestInsertionSortDichotomy(t *testing.T) {
	fmt.Println("###")
	a := []int{8, 5, 2, 6, 9, 3, 1, 4, 0, 7}
	InsertionSortDichotomy(a)
	fmt.Println(a)
}

func TestInsertionSortDichotomy2(t *testing.T) {
	fmt.Println("###")
	a := []int{8, 5, 2, 5, 5, 5, 1, 4, 0, 7}
	InsertionSortDichotomy(a)
	fmt.Println(a)
}
