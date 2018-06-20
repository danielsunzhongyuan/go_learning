package main

func main() {
    ch := make(chan int)
    exit := make(chan int)

    go func() {
        for i := 1; i <= 20; i++ {
            println("g1:", <-ch)
            i++
            ch <- i
        }
    }()

    go func() {
        defer func() {
            close(ch)
            close(exit)
        }()
        for i := 0; i < 20; i++ {
            i++
            ch <- i
            println("g2:", <-ch)
        }
    }()

    <-exit
}
