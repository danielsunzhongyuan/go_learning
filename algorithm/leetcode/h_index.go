package main

import (
	"fmt"
	"sort"
)

type Citations []int

func (arr Citations) Less(i, j int) bool {
	return arr[i] < arr[j]
}

func (arr Citations) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func (arr Citations) Len() int {
	return len(arr)
}

func hIndex(citations []int) int {
	length := len(citations)
	if length == 0 {
		return 0
	}
	c := make(Citations, length, length)
	copy(c, citations)
	sort.Sort(c)
	for i, c := range c {
		if c >= length-i {
			return length - i
		}
	}
	return 0
}

func main() {
	fmt.Println(hIndex([]int{3, 0, 6, 1, 5}))
}
