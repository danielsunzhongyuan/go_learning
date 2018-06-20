package main

import (
	"fmt"
	"github.com/pkg/errors"
	"math"
)

/*
给若干张1，2，5元的纸币，求可以凑出11元的最少张数
*/

func dpMin(smallChanges []int, target int) (int, error) {
	fan := make([]int, target+1, target+1)
	fan[0] = 1
	record := make([]string, target+1, target+1)

	lic := []string{}
	for _, i := range smallChanges {
		for j := i; j <= target; j++ {
			if fan[j-i] != target {
				fan[j] = int(math.Min(float64(fan[j-i]), float64(fan[j])))
				if fan[j-i]+1 <= fan[j] {
					record[j] = record[j-i] + string(i) + "*"
				}
			}
		}
		if record[len(record)-1] != "" {
			lic = append(lic, record[len(record)-1])
		}
	}
	if len(lic) == 0 {
		return 0, errors.New("无解")
	} else {
		fmt.Println(lic)
		return fan[target], nil
	}
}

func main() {
	a, err := dpMin([]int{2, 5}, 11)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)
}
