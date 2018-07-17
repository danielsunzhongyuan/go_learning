package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// 用 atomic 来加锁

var wg sync.WaitGroup
var counter int64

func main() {
	wg.Add(2)
	go incCounter(1)
	go incCounter(2)
	wg.Wait()
	fmt.Println("Final counter:", counter)
}

func incCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		atomic.AddInt64(&counter, 1)
		runtime.Gosched()
	}
}
