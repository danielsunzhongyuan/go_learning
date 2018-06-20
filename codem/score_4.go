package main

import (
	"fmt"
	"math"
	"sort"
)

/*
1. 只考虑0和C的情况，要修复一个bug，就是如果有多个成绩并列的时候，要以 people[k]的成绩作为参考。 通过了 77.78% 测试用例
2. 从 0遍历到C，只通过了 44.44% 测试用例
3. 还是按照只考虑 0和C 的情况，用改进后的代码，只通过了 37.04%
4. 修复一个bug， 55.56% (BUG：setMissingScoreAndMaxScore里判断max那里，应该先设置missingScore再判断 max）
5. 再修复一个bug，48.15%（setMissingScoreAndMaxScore那里，要对所有的 person 都重新设置 maxScore，而不是只设置 special person进行设置。照抄代码果然不靠谱）
6. 再修复bug，77.78%, 还是 setMissingScoreAndMaxScore这里，终于又回到起点了
7. 换一个思路，原先只考虑 0 和 C的情况，后来考虑遍历0到C，但是因为别的bug导致通过率很低，现在继续遍历，96.30%
8. 为啥需要遍历呢？因为如果 C 是这一轮的最大值，那么很多人在这一轮的分数就会变小，同时因为加权，从而改变整体的相对排名。所以进一步优化，只考虑0，然后遍历missingMax～C
9. 通过描述#8，知道了另一个bug：如果missingMax比C要大，那么遍历就失效了，相当于只考虑了0的状况，加上一个if else 后，还是 96.30%。
因为其实，曾经直接从0遍历到C的时候，就已经涵盖了这种情况，那时候就是96.30%，所以看来还是在别的地方有bug，没考虑边界吧。。。
*/

type Score struct {
	Value    int
	Special  bool
	MaxScore int
}

type Person struct {
	Result      int
	PersonIndex int
	TotalScore  float64
	Scores      []Score
	Special     bool
}

type SortBy func(p, q *Person) bool

type PersonWrapper struct {
	people []Person
	by     SortBy
}

func (pw PersonWrapper) Len() int {
	return len(pw.people)
}

func (pw PersonWrapper) Swap(i, j int) {
	pw.people[i], pw.people[j] = pw.people[j], pw.people[i]
}

func (pw PersonWrapper) Less(i, j int) bool {
	return pw.by(&pw.people[i], &pw.people[j])
}

func SortPerson(people []Person, by SortBy) {
	sort.Sort(PersonWrapper{people: people, by: by})
}

func main() {
	n, m, k, C := 0, 0, 0, 0
	fmt.Scan(&n, &m, &k, &C)
	weights := make([]float64, m, m)

	totalWeight := 0.0
	for i := 0; i < m; i++ {
		x := 0.0
		fmt.Scan(&x)
		totalWeight += x
		weights[i] = x
	}
	for i := 0; i < m; i++ {
		weights[i] /= totalWeight
	}

	people := make([]Person, 0, n)
	for i := 0; i < n; i++ {
		special := false
		scores := make([]Score, 0, m)
		for j := 0; j < m; j++ {
			x := 0
			fmt.Scan(&x)
			if x == -1 {
				special = true
				scores = append(scores, Score{Value: 0, Special: special}) // 先设置成 0，后面设置成C
			} else {
				scores = append(scores, Score{Value: x, Special: false})
			}
		}
		people = append(people, Person{PersonIndex: i, Scores: scores, Special: special})
	}
	// Get Input Done

	// Set MaxScore, Calculate TotalScore, SortByTotalScoreDesc
	setMaxScore(people, n, m)
	calculateTotalScore(people, n, m, weights)
	SortPerson(people, func(p, q *Person) bool {
		return p.TotalScore > q.TotalScore
	})
	setResult(people, n, k)

	// get "max of the round except the missing score"
	missingMax := getTheMaxExceptMissingScore(people, n, m)

	if missingMax < C {
		// set missing score from "missingMax" to "C"
		for i := missingMax; i <= C; i++ {
			setMissingScoreAndMaxScore(people, n, m, i)
			calculateTotalScore(people, n, m, weights)
			SortPerson(people, func(p, q *Person) bool {
				return p.TotalScore > q.TotalScore
			})
			setResult(people, n, k)
		}
	} else {
		setMissingScoreAndMaxScore(people, n, m, C)
		calculateTotalScore(people, n, m, weights)
		SortPerson(people, func(p, q *Person) bool {
			return p.TotalScore > q.TotalScore
		})
		setResult(people, n, k)
	}

	// SortByOriginalIndex asc, then print the result
	SortPerson(people, func(p, q *Person) bool {
		return p.PersonIndex < q.PersonIndex
	})
	for i := 0; i < n; i++ {
		fmt.Println(people[i].Result)
	}
}

func setMissingScoreAndMaxScore(people []Person, n, m, missingScore int) {
	for i := 0; i < n; i++ {
		if people[i].Special {
			for j := 0; j < m; j++ {
				if people[i].Scores[j].Special {
					people[i].Scores[j].Value = missingScore
				}
			}
		}
	}
	setMaxScore(people, n, m)
}

func getTheMaxExceptMissingScore(people []Person, n, m int) int {
	for i := 0; i < n; i++ {
		if people[i].Special {
			for j := 0; j < m; j++ {
				if people[i].Scores[j].Special {
					return people[i].Scores[j].MaxScore
				}
			}
		}
	}
	return 0
}

func calculateTotalScore(people []Person, n, m int, weight []float64) {
	for i := 0; i < n; i++ {
		people[i].TotalScore = 0.0
		for j := 0; j < m; j++ {
			if people[i].Scores[j].MaxScore > 0 {
				people[i].TotalScore += float64(people[i].Scores[j].Value) / float64(people[i].Scores[j].MaxScore) * weight[j]
			}
		}
	}
}

func setMaxScore(people []Person, n, m int) {
	for i := 0; i < m; i++ {
		maxScoreTmp := 0
		for j := 0; j < n; j++ {
			if maxScoreTmp < people[j].Scores[i].Value {
				maxScoreTmp = people[j].Scores[i].Value
			}
		}
		for k := 0; k < n; k++ {
			people[k].Scores[i].MaxScore = maxScoreTmp
		}
	}
}

func setResult(people []Person, n, k int) {
	for i := 0; i < n; i++ {
		switch people[i].Result {
		case 3:
			continue
		case 1:
			if i >= k {
				people[i].Result = 3
			}
		case 2:
			if i < k {
				people[i].Result = 3
			}
		default:
			if i < k {
				people[i].Result = 1
			} else {
				people[i].Result = 2
			}
		}
	}
	if k < n && math.Abs(people[k-1].TotalScore-people[k].TotalScore) < 0.00001 {
		for i := 0; i < n && people[i].Result < 3; i++ {
			if math.Abs(people[i].TotalScore-people[k].TotalScore) < 0.00001 {
				people[i].Result = 3
			}
		}
	}
}
