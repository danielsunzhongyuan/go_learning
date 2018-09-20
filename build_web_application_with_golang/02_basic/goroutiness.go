package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	runtime.GOMAXPROCS(1)
	for i := 0; i < 10; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}
