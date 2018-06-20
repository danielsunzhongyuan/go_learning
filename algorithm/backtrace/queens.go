package main

import (
	"fmt"
	"strings"
)

func solveNQueens(n int) [][]string {
	results := totalNQueens(n)
	ret := make([][]string, 0)
	for _, result := range results {
		tmp := []string{}
		for _, col := range result[1:] {
			tmpRow := generateString(col, n)
			tmp = append(tmp, tmpRow)
		}
		ret = append(ret, tmp)
	}
	return ret
}

func generateString(col, n int) string {
	ret := make([]string, 0)
	for i := 0; i < col-1; i++ {
		ret = append(ret, ".")
	}
	ret = append(ret, "Q")
	for i := 0; i < n-col; i++ {
		ret = append(ret, ".")
	}
	return strings.Join(ret, "")
}

func totalNQueens(n int) [][]int {
	results := make([][]int, 0)
	arr := make([]int, n+1, n+1)
	dfs(&results, arr, n, 1)
	return results
}

func dfs(results *[][]int, arr []int, n, r int) {
	if r > n {
		tmp := make([]int, len(arr), len(arr))
		copy(tmp, arr)
		*results = append(*results, tmp)
		return
	}
	for i := 1; i <= n; i++ {
		if check(arr, r, i) {
			arr[r] = i
			dfs(results, arr, n, r+1)
			arr[r] = 0
		}
	}
}

func check(arr []int, r, c int) bool {
	for i := 1; i < r; i++ {
		if arr[i] == c || arr[i]-c == r-i || arr[i]-c == i-r {
			return false
		}
	}
	return true
}

func main() {
	results := solveNQueens(8)
	for _, result := range results {
		fmt.Println(result)
	}

}
