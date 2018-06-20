package main

import "fmt"

// Queens1 只能打印 8 皇后问题，全遍历、慢、扩展性差
func Queens() {
	arr := make([]int, 9, 9)
	numberOfSolutions := 0
	tries := 0
	for i1 := 1; i1 <= 8; i1++ {
		arr[1] = i1
		for i2 := 1; i2 <= 8; i2++ {
			arr[2] = i2
			for i3 := 1; i3 <= 8; i3++ {
				arr[3] = i3
				for i4 := 1; i4 <= 8; i4++ {
					arr[4] = i4
					for i5 := 1; i5 <= 8; i5++ {
						arr[5] = i5
						for i6 := 1; i6 <= 8; i6++ {
							arr[6] = i6
							for i7 := 1; i7 <= 8; i7++ {
								arr[7] = i7
								for i8 := 1; i8 <= 8; i8++ {
									arr[8] = i8
									if checkForQueens(arr, 8) {
										numberOfSolutions++
										fmt.Println(arr[1:])
									}
									tries++
								}
							}
						}
					}
				}
			}
		}
	}
	fmt.Println("Total solutions:", numberOfSolutions, "with", tries, "tries")
}

func checkForQueens(arr []int, n int) bool {
	for i := 2; i <= n; i++ {
		for j := 1; j <= i-1; j++ {
			if arr[i] == arr[j] || abs(arr[i]-arr[j]) == i-j {
				return false
			}
		}
	}
	return true
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

// totalNQueens 时间上太慢， pow(9, 9) 就受不了了
func totalNQueens(n int) int {
	arr := make([]int, n, n)
	tmp, count := 0, 0
	for i := 0; i < pow(n, n); i++ {
		tmp = i
		for j := 0; j < n; j++ {
			arr[j] = tmp % n
			tmp /= n
		}
		if checkForTotalNQueens(arr, n) {
			count++
		}
	}
	return count
}

func pow(x, n int) int {
	ret := 1
	for i := 0; i < n; i++ {
		ret *= x
	}
	return ret
}

func checkForTotalNQueens(arr []int, n int) bool {
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if arr[i] == arr[j] || abs(arr[i]-arr[j]) == i-j {
				return false
			}
		}
	}
	return true
}

// dfs
func dfs(arr []int, numberOfSolutions []int, n int, r int) {
	if r > n {
		fmt.Println(arr[1:])
		numberOfSolutions[0]++
		return
	}
	for i := 1; i <= n; i++ {
		if checkForDfs(arr, r, i) {
			arr[r] = i
			dfs(arr, numberOfSolutions, n, r+1)
			arr[r] = 0
		}
	}
}

func checkForDfs(arr []int, r, c int) bool {
	for i := 1; i < r; i++ {
		if arr[i] == c || arr[i]-c == r-i || arr[i]-c == i-r {
			return false
		}
	}
	return true
}

func main() {
	//Queens()
	//fmt.Println(totalNQueens(5))

	numberOfSolutions := make([]int, 1, 1)
	n := 8
	arr := make([]int, n+1, n+1)
	dfs(arr, numberOfSolutions, n, 1)
	fmt.Println(numberOfSolutions)
}
