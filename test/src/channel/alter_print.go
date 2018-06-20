package main

func main() {
    ch1 := make(chan int)
    ch2 := make(chan int)

    exit := make(chan int)

    go func() {
        defer close(ch2)
        for i := 1; i <= 20; i++ {
            println("g1:", <-ch1)
            i++
            ch2 <- i
        }
    }()

    go func() {
        defer close(ch1)
        for i := 0; i < 20; i++ {
            i++
            ch1 <- i
            println("g2:", <-ch2)
        }
        close(exit)
    }()

    <-exit
}
