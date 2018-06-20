package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var rwmutex sync.RWMutex

	for i := 0; i < 3; i++ {
		go func(i int) {
			fmt.Printf("Try to lock for reading... [%d]\n", i)
			rwmutex.RLock()
			fmt.Printf("Locked for reading. [%d]\n", i)
			time.Sleep(time.Second * 2)
			fmt.Printf("Try to unlock for reading... [%d]\n", i)
			rwmutex.RUnlock()
			fmt.Printf("Unlocked for reading. [%d]\n", i)
		}(i)
	}
	time.Sleep(time.Microsecond * 100)
	fmt.Println("Try to lock for writing...")
	rwmutex.Lock()
	fmt.Println("Locked for writing.")
}
