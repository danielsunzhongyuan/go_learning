package main

import "fmt"

func main() {
	n := 16
	win := make([][]float64, n, n)
	winner := make([][]float64, n, n)

	for i := 0; i < n; i++ {
		win[i] = make([]float64, n, n)
		winner[i] = make([]float64, 4, 4) // 晋级8强的概率，晋级4强的概率，晋级决赛的概率，最终获胜的概率。总共4场比赛。
		for j := 0; j < n; j++ {
			fmt.Scan(&win[i][j])
		}
	}

	// 8强
	for i := 0; i < n; i++ {
		winner[i][0] = win[i][(i/2)*2+1^(i%2)]
	}

	for round := 1; round < 4; round++ {
		for i := 0; i < n; i++ {
			competitorIndices := getCompetitorIndices(i, round+1)
			for _, j := range competitorIndices {
				winner[i][round] += winner[i][round-1] * (winner[j][round-1] * win[i][j])
			}
		}
	}
	res := ""
	for i := 0; i < n; i++ {
		res += fmt.Sprintf("%.10f ", winner[i][3])
	}
	res = res[:len(res)-1] // remove last whitespace
	fmt.Println(res)
}

func getCompetitorIndices(i, round int) []int {
	teams := power(2, round)
	halfTeams := power(2, round-1)
	competitorIndices := []int{}
	for j := 0; j < halfTeams; j++ {
		competitorIndices = append(competitorIndices, (i/teams)*teams+(j+halfTeams+(i/halfTeams)*halfTeams)%teams)
	}
	return competitorIndices
}

func power(x int, n int) int {
	ans := 1

	for n != 0 {
		if n%2 == 1 {
			ans *= x
		}
		x *= x
		n /= 2
	}
	return ans
}
