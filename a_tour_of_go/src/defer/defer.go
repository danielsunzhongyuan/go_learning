package main

/*
A defer statement defers the execution of a function until the surrounding function returns.

The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
*/
import "fmt"

func main() {
    defer fmt.Println("world3")
    defer fmt.Println("world2")
    defer fmt.Println("world1")
    fmt.Println("hello")
}
