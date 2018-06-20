package main

import "fmt"

func sum(s []int, c chan int) {
    sum := 0
    for _, v := range s{
        sum += v
    }
    c <- sum
}

func foo() {

    s := []int{7, 2, 8, -9, 4, 0, 10}

    c := make(chan int)
    go sum(s[:len(s)/2], c)
    go sum(s[len(s)/2:], c)

    x := <- c
    y := <- c
    fmt.Println(x, y, x + y)
}

func main() {
    // 有16个打印的是 17 5 22
    for i := 0; i < 1000; i++ {
        foo()
    }
}
