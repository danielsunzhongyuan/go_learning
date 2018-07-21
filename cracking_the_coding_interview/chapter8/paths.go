package main

import "fmt"

func factorial(n int) int {
	res := 1
	for n > 1 {
		res *= n
		n -= 1
	}
	return res
}

func paths(x, y int) int {
	return factorial(x+y-2) / factorial(x-1) / factorial(y-1)
}

func main() {
	fmt.Printf("3 * 3 grid has %d paths\n", paths(3, 3))
	fmt.Printf("2 * 4 grid has %d paths\n", paths(2, 4))
	fmt.Printf("10 * 10 grid has %d paths\n", paths(10, 10))
}
