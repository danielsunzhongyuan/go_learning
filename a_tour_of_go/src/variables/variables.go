package main

import "fmt"

var c, python bool
var java = "no!"
//k := 3   not allowed outside the function body
var i, j int = 1, 2


func main() {
    swift := "swift"   // allowed
    fmt.Println(c, python, java)
    fmt.Println(i, j)
    fmt.Println(swift)
}
