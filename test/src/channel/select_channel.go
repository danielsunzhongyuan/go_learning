package main

import (
	"fmt"
	"math/rand"
	"time"
)

func doStuffInSelect(x int) int { // 一个比较耗时的事情，比如计算
	fmt.Println("calculating x ...", x)
	time.Sleep(time.Duration(rand.Intn(3)) * time.Second) //模拟计算

	return 100 - x // 假如100-x是一个很费时的计算
}

func branchInSelect(x int) chan int { // 每个分支开出一个goroutine做计算并把计算结果流入各自信道
	ch := make(chan int)
	go func() {
		ch <- doStuffInSelect(x)
	}()
	return ch
}

// 多路复合成一路，但是顺序被固定了，不能是某个搞定了就可以输出了。也就是肯定是安装 1 2 3 的结果输出
// 解决办法就是 select
func fanInUsingSelect(chs ...chan int) chan int {
	ch := make(chan int)

	for i := 0; i < len(chs); i++ {
		go func(i int) {
			defer close(chs[i])
			select {
			case v1 := <-chs[i]:
				ch <- v1
			}
		}(i)

	}

	return ch
}

func main() {
	result := fanInUsingSelect(branchInSelect(1), branchInSelect(2), branchInSelect(3))
	for i := 0; i < 3; i++ {
		fmt.Println(<-result)
	}

}
