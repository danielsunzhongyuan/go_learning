package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Println("Start Goroutines")

	go func() {
		defer wg.Done()

		// 如果 count < 3 的话，会先打印大写字母，打印完再打印小写字母，因为运行的太快了。
		// 调整成 count<2000，就能看见输出不那么规整了。会出现一行大写字母没打印完的情况下，就切换到小写字母了。
		for count := 0; count < 2000; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
			fmt.Println()
		}
	}()

	go func() {
		defer wg.Done()

		for count := 0; count < 2000; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
			fmt.Println()
		}
	}()

	fmt.Println("Waiting to finish")
	wg.Wait()
	fmt.Println("Terminating Program")
}
