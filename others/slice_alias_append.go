package main

/*
https://bbs.csdn.net/topics/392372776
原来的目的是不使用 struct 的形式，而是别名的方式去定义一个stack，然后给stack上加上push、pop的操作。
但是不行。
*/

import (
	"fmt"
	"strconv"
)

const stackLength = 4

//type Stack []int
type Stack struct {
	A []string
}

func (self *Stack) Size() (lenth int) {
	lenth = len(self.A)
	return
}

func (self *Stack) String() (stackInfo string) {
	for key, value := range self.A {
		stackInfo += "[" + strconv.Itoa(key) + ":" + value + "]"
	}
	return
}

func (self *Stack) Push(s string) {
	self.A = append(self.A, s)
}

func (self *Stack) Pop() {
	lastKey := self.Size()
	if lastKey > 0 {
		self.A[lastKey-1] = ""
	}

}

func main() {
	var stack = Stack{A: make([]string, 0, stackLength)}
	fmt.Println(stack.Size())
	stack.Push("a")
	fmt.Println(stack)
	stack.Push("b")
	fmt.Println(stack)
	stack.Pop()
	fmt.Println(stack)

}
