package main

import (
	"fmt"
	"math"
)

func isPrime(n int64) bool {
	tmp := int64(math.Sqrt(float64(n)))
	for i := int64(2); i < tmp+1; i++ {
		if n%i == 0 {
			fmt.Println(i, n/i)
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(power(10, 9))
	fmt.Println(isPrime(power(10, 9) + 7))
}

func power(x, y int) int64 {
	res := int64(1)
	for y > 0 {
		res *= int64(x)
		y -= 1
	}
	return res
}
