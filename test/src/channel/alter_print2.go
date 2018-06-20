package main

import (
    "fmt"
    "runtime"
    "time"
)

func main() {
    a := make(chan int, 1)
    b := make(chan int, 0)
    runtime.GOMAXPROCS(1)
    for i := 1; i <= 10; i++ {
        go func(i int) {
            <-a
            fmt.Println(2*i - 1)
            b <- i
        }(i)
    }

    for i := 1; i <= 10; i++ {
        go func(i int) {
            <-b
            fmt.Println(2 * i)
            a <- i
        }(i)
    }
    a <- 1
    time.Sleep(4 * time.Second)
}
