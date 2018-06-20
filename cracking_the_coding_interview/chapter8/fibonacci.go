package main

import (
	"fmt"
)

func generateFibonacci(n int, c chan int) {
	x, y := 0, 1
	for x <= n {
		c <- x
		x, y = y, x+y
	}
	close(c)
	return
}

func main() {
	c := make(chan int)
	go func() {
		generateFibonacci(13, c)
	}()

	for x, ok := <-c; ok; {
		fmt.Println(x)
		x, ok = <-c
	}
}
