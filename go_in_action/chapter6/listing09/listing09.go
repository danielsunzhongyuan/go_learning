package main

import (
	"fmt"
	"runtime"
	"sync"
)

var counter int
var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go incCounter(1)
	go incCounter(2)
	wg.Wait()
	fmt.Println("Final Counter", counter) // 期望是4，实际是2

	// go build -race 可以build一个包，运行后可以检测竞争状态
}

func incCounter(id int) {
	defer wg.Done()
	for count := 0; count < 2; count++ {
		value := count
		runtime.Gosched()
		value++
		counter = value
	}
}
