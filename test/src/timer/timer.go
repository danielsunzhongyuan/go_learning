package main

import (
	"fmt"
	"time"
)

func timer(duration time.Duration) chan bool {
	ch := make(chan bool)
	go func() {
		time.Sleep(duration)
		ch <- true
	}()
	return ch
}

func main() {
	timeout := timer(time.Second * 2)
	for {
		select {
		case <-timeout:
			fmt.Println("2 seconds...")
			//timeout = timer(time.Second * 2) // 这样就一直tick tok 下去了
			return
		}
	}
}
