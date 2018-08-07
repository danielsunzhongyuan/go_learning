package main

import "fmt"

func GetMaxSum(arr []int) int {
	maxSum := 0
	sum := 0

	for _, num := range arr {
		sum += num
		if maxSum < sum {
			maxSum = sum
		} else if sum < 0 {
			sum = 0
		}
	}
	return maxSum
}

func main() {
	fmt.Println(GetMaxSum([]int{2, -8, 3, -2, 4, -10}))
}
