package main

import (
	"fmt"
	"math/rand"
	"time"
)

func doStuff(x int) int { // 一个比较耗时的事情，比如计算
	fmt.Println("calculating x ...", x)
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second) //模拟计算

	return 100 - x // 假如100-x是一个很费时的计算
}

func branch(x int) chan int { // 每个分支开出一个goroutine做计算并把计算结果流入各自信道
	ch := make(chan int)
	go func() {
		ch <- doStuff(x)
	}()
	return ch
}

// 多路复合成一路，但是顺序被固定了，不能是某个搞定了就可以输出了。也就是肯定是安装 1 2 3 的结果输出
// 解决办法就是 select
func fanIn(chs ...chan int) chan int {
	ch := make(chan int)

	for _, c := range chs {
		// 注意此处明确传值
		go func(c chan int) { ch <- <-c }(c) // 复合
	}

	return ch
}

func main() {
	result := fanIn(branch(1), branch(2), branch(3))
	for i := 0; i < 3; i++ {
		fmt.Println(<-result)
	}

}
