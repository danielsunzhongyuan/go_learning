package main

import "fmt"

func getSum(a int, b int) int {
	for a != 0 && b != 0 {
		a, b = a^b, (a&b)<<1
		fmt.Println("Middle result:", a, b)
	}
	if a != 0 {
		return a
	}
	return b
}

func main() {
	fmt.Println(getSum(-14, 16))
	fmt.Println(getSum(-1, 1))
}
