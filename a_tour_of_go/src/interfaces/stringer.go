package main

import "fmt"

type Person struct {
    Name string
    Age int
}

func (p Person) String() string {
    return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
    a := Person{"Arthur Dent", 43}
    z := Person{"Zaphod Beeblebrox", 9001}

    aa := a.String()
    fmt.Println(aa)
    fmt.Println(a, z)
}
