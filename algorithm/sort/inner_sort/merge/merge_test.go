package merge

import (
	"fmt"
	"testing"
)

func TestMergeSortRecursionEmptySlice(t *testing.T) {
	a := []int{}
	MergeSortRecursion(a, 0, len(a)-1)
	fmt.Println(a)
}

//func TestMergeSortRecursionOneElement(t *testing.T) {
//	a := []int{1}
//	MergeSortRecursion(a, 0, len(a)-1)
//	fmt.Println(a)
//}
//
//func TestMergeSortRecursion(t *testing.T) {
//	a := []int{8, 5, 2, 6, 9, 3, 1, 4, 0, 7}
//	MergeSortRecursion(a, 0, len(a)-1)
//	fmt.Println(a)
//}
//
//func TestMergeSortIterationEmptySlice(t *testing.T) {
//	a := []int{}
//	MergeSortIteration(a, len(a))
//	fmt.Println(a)
//}
//
//func TestMergeSortIterationOneElement(t *testing.T) {
//	a := []int{1}
//	MergeSortIteration(a, len(a))
//	fmt.Println(a)
//}
//
//func TestMergeSortIteration(t *testing.T) {
//	a := []int{8, 5, 2, 6, 9, 3, 1, 4, 0, 7}
//	MergeSortIteration(a, len(a))
//	fmt.Println(a)
//}
