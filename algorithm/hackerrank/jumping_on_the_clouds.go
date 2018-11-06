package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Complete the jumpingOnClouds function below.
func jumpingOnClouds(c []int32) int32 {
	var result int32 = 0

	int32tostring := make([]string, len(c))
	for i := 0; i < len(c); i++ {
		int32tostring[i] = strconv.Itoa(int(c[i]))
	}
	s := strings.Join(int32tostring, "")
	s2 := strings.Split(s, "1")
	for ind, substring := range s2 {
		if ind != len(s2)-1 {
			result += int32(len(substring)/2) + 1
		} else {
			result += int32(len(substring) / 2)
		}
	}
	return result
}

func main() {
	// 4 3 6
	fmt.Println(jumpingOnClouds([]int32{0, 0, 1, 0, 0, 1, 0}))
	fmt.Println(jumpingOnClouds([]int32{0, 0, 0, 1, 0, 0}))
	fmt.Println(jumpingOnClouds([]int32{0, 0, 1, 0, 0, 0, 0, 1, 0, 0}))
}
