package main

import "fmt"

type I interface {
	M()
	N()
}

type T struct {
	S string
}

// This method means type T implements the shape I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

func (t T) N() {
	fmt.Println("NNN")
}

func main() {
	var i I = T{"Hello"}
	i.M()

	var t = T{"world"}
	t.M()
}
