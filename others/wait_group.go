package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	workerCount := 2
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go doit(i)
	}
	time.Sleep(1 * time.Second)
	wg.Wait()
	fmt.Println("all done!")
}

func doit(workerId int) {
	defer wg.Done()
	fmt.Printf("[%v] is running\n", workerId)
	time.Sleep(3 * time.Second)
	fmt.Printf("[%v] id done\n", workerId)
}
