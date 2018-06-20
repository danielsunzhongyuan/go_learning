package search

import "fmt"

var FiboArr []int

func InitFiboArr(n int) {
	a, b := 0, 1
	for a < n {
		FiboArr = append(FiboArr, a)
		a, b = b, a+b
	}
	FiboArr = append(FiboArr, a)
}

func FibonacciSearch(arr []int, x int) int {
	length := len(arr)
	//l, h, mid, k := 0, length-1, 0, 0

	InitFiboArr(length)
	fmt.Println(FiboArr)
	return 0
}
