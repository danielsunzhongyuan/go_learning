package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	// 如果只有一行或者一列，那么只要有障碍物就过不去，返回0，否则返回1
	if m == 1 || n == 1 {
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if obstacleGrid[i][j] == 1 {
					return 0
				}
			}
		}
		return 1
	}

	arr := make([][]int, 0, m)
	for i := 0; i < m; i++ {
		arr = append(arr, make([]int, n, n))
	}

	// 如果终点是一个障碍的话，就根本过不去
	if obstacleGrid[m-1][n-1] == 1 {
		return 0
	}
	// 如果终点不是障碍的话，就把arr的终点置为1
	arr[m-1][n-1] = 1

	// 把前 m-1 行的最后一列设置一下，如果有障碍物，就设置为0，否则等于它的下面一个位置的值。
	// 其实如果某一行的最后一个位置为障碍物，那么它上面的每一行的最后一列都是 0
	for i := m - 2; i >= 0; i-- {
		if obstacleGrid[i][n-1] == 1 {
			arr[i][n-1] = 0
		} else {
			arr[i][n-1] = arr[i+1][n-1]
		}
	}

	// 把最后一行的前 n-1 列设置一下，如果有障碍物，就设置为0，否则等于它的右边一个位置的值。
	// 其实如果最后一行的某一个位置为障碍物，那么它左边的每一列都是 0
	for i := n - 2; i >= 0; i-- {
		if obstacleGrid[m-1][i] == 1 {
			arr[m-1][i] = 0
		} else {
			arr[m-1][i] = arr[m-1][i+1]
		}
	}

	for i := m - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			if obstacleGrid[i][j] == 1 {
				arr[i][j] = 0
			} else {
				arr[i][j] = arr[i+1][j] + arr[i][j+1]
			}
		}
	}
	return arr[0][0]
}

func main() {
	fmt.Println(uniquePathsWithObstacles([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}))
}
