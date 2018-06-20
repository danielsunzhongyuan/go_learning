package main

import "fmt"

// 方法1，从最后一个数字开始，分别除以2、3、5，以找到应该乘以2、3、5的数，然后比较这些数各自乘以2、3、5的最小值，append
// 慢
func nthUglyNumber(n int) int {
	uglies := []int{1, 2, 3, 4, 5, 6}

	if n <= 6 {
		return uglies[n-1]
	}
	for i := 6; i < n; i++ {
		x := next(uglies, i)
		uglies = append(uglies, x)
	}
	return uglies[n-1]
}

func next(uglies []int, ith int) int {
	return min(
		findNext(uglies, 2, ith),
		findNext(uglies, 3, ith),
		findNext(uglies, 5, ith),
	)
}

func findNext(uglies []int, x int, ith int) int {
	last := uglies[ith-1]
	tmp := last / x
	return uglies[search(uglies, tmp)] * x
}

func search(arr []int, x int) int {
	i, j := 0, len(arr)-1
	for i < j {
		mid := (i + j) / 2
		if arr[mid] > x {
			j = mid - 1
		} else if arr[mid] < x {
			i = mid + 1
		} else {
			return mid + 1
		}
	}
	if arr[i] <= x {
		return i + 1
	}
	return i
}

// 方法 2， 直接记录2，3，5倍数的index
// 方法2比方法1要快很多，但是当数字很大时，会更早的溢出，
// 比如计算第 12000000 个ugly number时，方法1大约花了4秒，得到 2814749767106560，而方法二得到的是 -500371793013067776
func nthUglyNumber2(n int) int {
	uglies := make([]int, n, n)
	uglies[0] = 1
	index2, index3, index5 := 0, 0, 0

	for i := 1; i < n; i++ {
		nextFactor2, nextFactor3, nextFactor5 := uglies[index2]*2, uglies[index3]*3, uglies[index5]*5
		if nextFactor2 <= nextFactor3 && nextFactor2 <= nextFactor5 {
			uglies[i] = nextFactor2
			index2++
		}
		if nextFactor3 <= nextFactor2 && nextFactor3 <= nextFactor5 {
			uglies[i] = nextFactor3
			index3++
		}
		if nextFactor5 <= nextFactor2 && nextFactor5 <= nextFactor3 {
			uglies[i] = nextFactor5
			index5++
		}
	}
	return uglies[n-1]
}

func main() {
	fmt.Println(nthUglyNumber(12000000))
	//fmt.Println(nthUglyNumber2(12000000))
}

func min(x, y, z int) int {
	if x < y && x < z {
		return x
	}
	if y < z {
		return y
	}
	return z
}
