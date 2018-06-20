package main

import "fmt"

type TT struct {
	A int
	B string
}

func main() {
	fmt.Printf("sd %v\n", "ksd")
	a := fmt.Errorf("%vzzz", "kd ")
	fmt.Println(a)
	fmt.Printf("%+5.2f\n", -123456.789)
	fmt.Printf("%q\n", "sdkf")

	t := &TT{A: 10, B: "tt"}
	fmt.Printf("%v\n", t)
	fmt.Printf("%+v\n", t)
	fmt.Printf("%#v\n", t)
	fmt.Printf("%T\n", t)
	fmt.Printf("%p\n", t)
	fmt.Printf("%q\n", t)
}
