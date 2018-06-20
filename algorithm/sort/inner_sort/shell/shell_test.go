package shell

import (
	"fmt"
	"testing"
)

func TestShellSortEmptySlice(t *testing.T) {
	fmt.Println("###")
	a := []int{}
	ShellSort(a)
	fmt.Println(a)
}

func TestShellSortOneElement(t *testing.T) {
	fmt.Println("###")
	a := []int{11}
	ShellSort(a)
	fmt.Println(a)
}

func TestShellSort(t *testing.T) {
	fmt.Println("###")
	a := []int{8, 5, 2, 6, 9, 3, 1, 4, 0, 7}
	ShellSort(a)
	fmt.Println(a)
}
