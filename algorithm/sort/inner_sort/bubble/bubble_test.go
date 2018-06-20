package bubble

import (
	"fmt"
	"testing"
)

func TestBubbleSortEmptySlice(t *testing.T) {
	a := []int{}
	BubbleSort(a)
	fmt.Println(a)
}

func TestBubbleSortOneElement(t *testing.T) {
	a := []int{1}
	BubbleSort(a)
	fmt.Println(a)
}

func TestBubbleSort(t *testing.T) {
	a := []int{8, 5, 2, 6, 9, 3, 1, 4, 0, 7}
	BubbleSort(a)
	fmt.Println(a)
}
