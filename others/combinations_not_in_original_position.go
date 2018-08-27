package main

/*
有n（n≥1）个连续正整数按照1，2，3，······，n的顺序排列，现移动每个数使其均不在最初的位置，问满足此要求的排列一共有多少个。
*/

import "fmt"

func f(n int) int {
    if 1 >= n {
        return 0
    }
    if n % 2 == 0 {
        return 1 + n*f(n-1)
    }
    return -1 + n*f(n-1)
}

func main() {
    for i:=0; i<10; i++ {
        fmt.Println(i, f(i))
    }
}
