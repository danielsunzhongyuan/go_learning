package main

import "fmt"

func sum(nums ...int) {
    fmt.Println(nums, " ")
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}

func main() {
    sum(9,2,3,4)
    nums := []int {4,2,3,4}
    sum(nums...)
}
