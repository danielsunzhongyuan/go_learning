package main

import "fmt"

func xrange() chan int {
	var ch = make(chan int)

	go func() {
		for i := 2; ; i++ { // 从2开始自增的整数生成器
			ch <- i // 直到信道索要数据，才把i添加进信道
		}
	}()
	return ch
}

func filter(in chan int, number int) chan int {
	// 输入一个整数队列，筛出是number倍数的, 不是number的倍数的放入输出队列
	// in:  输入队列
	out := make(chan int)

	go func() {
		for {
			i := <-in
			if i%number != 0 {
				out <- i
			}
		}
	}()
	return out
}

func main() {
	const max = 100
	nums := xrange()
	number := <-nums

	for number <= max {
		fmt.Println(number)
		nums = filter(nums, number)
		number = <-nums
	}
}
