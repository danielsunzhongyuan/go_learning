package main

import "fmt"

func main() {
	ch := make(chan int)
	for i := 1; i <= 20; i++ {
		go func() {
			fmt.Println(<-ch)
		}()
		ch <- i
	}
}

