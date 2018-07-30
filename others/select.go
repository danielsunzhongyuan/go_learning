package main

import (
    "fmt"
    "time"
    "runtime"
)

func main() {
    c1 := make(chan string)
    c2 := make(chan string)

    go func() {
        time.Sleep(time.Second * 1)
        c1 <- "one"
    }()

    go func() {
        time.Sleep(time.Second * 2)
        c2 <- "two"
    }()

    for i := 0; i < 2; i++ {
        select {
            case msg1 := <- c1:
                fmt.Println("received", msg1)
            case msg2 := <- c2:
                fmt.Println("received", msg2)
        }
    }

    runtime.GOMAXPROCS(1)
    int_chan := make(chan int, 1)
    string_chan := make(chan string, 1)
    int_chan <- 1
    string_chan <- "hello"
    select {
    case value := <-int_chan:
        fmt.Println(value)
    case value := <-string_chan:
        panic(value)
    }
    // it will randomly panic or just print 1 instead.
}
