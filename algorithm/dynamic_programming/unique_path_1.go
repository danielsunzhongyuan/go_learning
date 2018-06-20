package main

import "fmt"

func uniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}

	arr := make([][]int, 0, m)
	for i := 0; i < m; i++ {
		arr = append(arr, make([]int, n, n))
	}

	for i := 0; i < m; i++ {
		arr[i][n-1] = 1
	}
	for i := 0; i < n; i++ {
		arr[m-1][i] = 1
	}
	for i := m - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			arr[i][j] = arr[i+1][j] + arr[i][j+1]
		}
	}
	return arr[0][0]
}

func main() {
	fmt.Println(uniquePaths(3, 2))
	fmt.Println(uniquePaths(7, 3))
}
