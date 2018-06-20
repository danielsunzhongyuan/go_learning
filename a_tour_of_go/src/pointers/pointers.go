package main

import "fmt"

func foo1(x *int) {
	*x = *x + 1
}

func foo2(x *int) int {
	return *x + 1
}

func main() {
    i, j := 42, 2701

    p := &i
    fmt.Println(*p)
    *p = 21
    fmt.Println(i)

    p = &j
    *p = *p / 37
    fmt.Println(j)

	k := 32
    fmt.Println(k)
    q := &k
    foo1(q)
    fmt.Println(k)

	*q = foo2(q)
	fmt.Println(k)
}
