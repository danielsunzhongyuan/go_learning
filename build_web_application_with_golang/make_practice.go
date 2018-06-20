package main

import (
    "fmt"
)

func main() {
    var s1 = make([]int, 0)
    var s2 []int
    fmt.Println(s1 == nil)  // false
    fmt.Println(s2 == nil)  // true
}
