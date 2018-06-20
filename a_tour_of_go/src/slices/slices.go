package main

import "fmt"

func main() {
    primes := [6]int {2, 3, 5, 7, 11, 13}
    var s []int = primes[1:4]
    fmt.Println(s)

    // slices-pointers
    names := [4]string {"John", "Paul", "George", "Ringo",}
    a := names[0:2]
    b := names[1:3]
    fmt.Println(a, b)

    b[0] = "XXX"
    fmt.Println(a, b)
    fmt.Println(names)

    // slice-literals
    q := []int{2,3,5,7,11,13}
    fmt.Println(q)
    r := []bool{true, false, true, true, false, true}
    fmt.Println(r)
    ss := []struct{
        i int
        b bool
    }{
        {2, true},
        {3, false},
        {5, true},
        {7, true},
        {11, false},
        {13, true},
    }
    fmt.Println(ss) 

    // slice-bounds
    fmt.Println(q[1:4])
    fmt.Println(q[:2])
    fmt.Println(q[3:])
    fmt.Println(q[:])
}
