package main

import "fmt"

func main() {
    s := []int{2,3,5,7,11,13,17,19}
    printSlice(s)

    s = s[:0]
    printSlice(s)

    s = s[:4]
    printSlice(s)

    // here the cap is reduced to 6.
    s = s[2:]
    printSlice(s)

    s = s[1:]
    printSlice(s)

    s = s[:5]
    printSlice(s)

}

func printSlice(s []int) {
    fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
