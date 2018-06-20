package main

import (
	"fmt"
	"reflect"
)

func Decorator(decoPtr, fn interface{}) (err error) {
	var decoratedFunc, targetFunc reflect.Value

	decoratedFunc = reflect.ValueOf(decoPtr).Elem()
	targetFunc = reflect.ValueOf(fn)

	v := reflect.MakeFunc(targetFunc.Type(),
		func(in []reflect.Value) (out []reflect.Value) {
			fmt.Println("before")
			out = targetFunc.Call(in)
			fmt.Println("after")
			return
		})

	decoratedFunc.Set(v)
	return
}

func foo(a, b, c int) int {
	fmt.Printf("%d, %d, %d \n", a, b, c)
	return a + b + c
}

func bar(a, b string) string {
	fmt.Printf("%s, %s \n", a, b)
	return a + b
}

type MyFoo func(int, int, int) int
type MyBar func(string, string) string

func main() {
	var myfoo MyFoo
	Decorator(&myfoo, foo)
	myfoo(1, 2, 3)

	var mybar MyBar
	Decorator(&mybar, bar)
	mybar("a", "b")

	// 上面的方式需要提前声明一个函数类型，下面的方式不需要
	myfoo2 := foo
	Decorator(&myfoo2, foo)
	myfoo2(4, 5, 6)
}
