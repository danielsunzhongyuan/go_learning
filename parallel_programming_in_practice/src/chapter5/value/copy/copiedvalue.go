package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var countVal atomic.Value
	countVal.Store([]int{1, 3, 5, 7})
	anotherStore(&countVal)
	fmt.Printf("The count value: %+v\n", countVal.Load())
}

// 必须传进来的是指针，否则打印出来的还是 1，3，5，7。
// 而且 go vet 会报错。
func anotherStore(countVal *atomic.Value) {
	countVal.Store([]int{2, 4, 6, 8})
}
