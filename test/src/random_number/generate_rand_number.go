package main

import "fmt"

func rand01() chan int {
	ch := make(chan int)

	go func() {
		for {
			select {
			case ch <- 0:
			case ch <- 1:
			}
		}
	}()
	return ch
}

func main() {
	generator := rand01()

	for i := 0; i < 10; i++ {
		fmt.Println(<-generator)
	}
}
