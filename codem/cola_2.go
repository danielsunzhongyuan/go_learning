package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Happiness struct {
	M           int
	T           int
	Total       float64
	OriginIndex int
}

type AllHappiness []Happiness

func (a AllHappiness) Len() int {
	return len(a)
}

func (a AllHappiness) Less(i, j int) bool {
	if a[i].Total == a[j].Total {
		return i > j
	}
	return a[i].Total > a[j].Total
}

func (a AllHappiness) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	n, m, k := 0, 0, 0
	fmt.Scan(&n, &m, &k)

	allHappiness := make(AllHappiness, 0, k)

	for i := 0; i < k; i++ {
		mei, tuan := 0, 0
		fmt.Scan(&mei, &tuan)
		allHappiness = append(allHappiness, Happiness{
			M:           mei,
			T:           tuan,
			Total:       float64(m)/float64(n)*float64(mei) + float64(n-m)/float64(n)*float64(tuan),
			OriginIndex: i,
		})
	}
	sort.Sort(allHappiness)

	res := make([]string, k, k)
	for i := 0; i < k; i++ {
		res[i] = "0"
	}

	res[allHappiness[0].OriginIndex] = strconv.Itoa(n)
	fmt.Println(strings.Join(res, " "))
}
