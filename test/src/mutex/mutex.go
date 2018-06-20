package main

/*
go mutex是互斥锁，只有Lock和Unlock两个方法，在这两个方法之间的代码不能被多个goroutins同时调用到。
*/

import (
	"fmt"
	"sync"
	"time"
)

var mutex *sync.Mutex

func main() {
	mutex = new(sync.Mutex)
	go lockPrint(1)
	lockPrint(2)
	time.Sleep(3 * time.Second)
	fmt.Printf("%s\n", "exit!")
}

func lockPrint(i int) {
	println(i, "lock start")
	mutex.Lock()
	println(i, "in lock")
	time.Sleep(2 * time.Second)
	mutex.Unlock()
	println(i, "unlock")
}
