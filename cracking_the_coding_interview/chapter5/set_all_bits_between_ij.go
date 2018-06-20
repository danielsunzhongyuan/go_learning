package main

import "fmt"

/*
Given two 32-bit numbers, N and M, and two bit positions, i and j.
Write a method to set all bits between i and j in N equal to M.
Example:
Input: N = 10000000000, M = 10101, i=2, j=6
Output: N = 10001010100
*/

func updateBits(n, m int, i, j uint) int {
	max := 1<<31 - 1
	max = max << 1
	max += 1

	left := max - ((1 << j) - 1)
	right := (1 << i) - 1
	mask := left | right
	return (n & mask) | (m << i)
}

func main() {
	fmt.Println(updateBits(1<<10, 21, 2, 6)) // 1108
}
