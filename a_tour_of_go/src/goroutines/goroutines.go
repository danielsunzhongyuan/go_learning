package main

import (
    "fmt"
    "time"
)

func say(s string) {
    for i := 0; i < 9; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Printf("%s sleep %d\n", s, i)
    }
}

func main() {
    go say("world")
    say("hello")

    // 没有下面这句，如果主程序退出很快，那么 world 有可能少打印一个
    time.Sleep(1000 * time.Millisecond)
}
