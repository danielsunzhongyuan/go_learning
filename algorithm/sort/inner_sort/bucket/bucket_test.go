package bucket

import (
	"fmt"
	"testing"
)

func TestRadixSortEmptySlice(t *testing.T) {
	a := []int{}
	BucketSort(a, len(a))
	fmt.Println(a)
}

func TestRadixSortOneElement(t *testing.T) {
	a := []int{1}
	BucketSort(a, len(a))
	fmt.Println(a)
}

func TestRadixSort(t *testing.T) {
	a := []int{29, 25, 3, 49, 9, 37, 21, 43}
	BucketSort(a, len(a))
	fmt.Println(a)
}
