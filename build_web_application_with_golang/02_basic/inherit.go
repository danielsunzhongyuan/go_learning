package main

import "fmt"

type A struct {
	Name string
	Age  int64
}

type B struct {
	Name string
	Age  int64
}

type C struct {
	A
	B
	Name string
}

func (a *A) SayHi() {
	fmt.Println("A say hi:", a.Name)
}

func (b *B) SayHi() {
	fmt.Println("B say hi:", b.Name)
}

func (c *C) SayHi() {
	fmt.Println("C say hi:", c.Name)
}

func main() {
	c := C{A{Name: "a name in C", Age: 10}, B{Name: "b name in C", Age: 12}, "c name"}
	// 如果C没有重写 SayHi方法，那么就不能用下面的方式调用SayHi，必须指定调用的是A的SayHi还是B的SayHi
	// 即：c.A.SayHi() 或者 c.B.SayHi()
	c.SayHi()

	// 同理，成员变量也是如此，如果没有复写 Age 这个变量，那么就需要指定Age的具体来源
	// 即：c.A.Age 或者 c.B.Age
	fmt.Println(c.B.Age)
}
