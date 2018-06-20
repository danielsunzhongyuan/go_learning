package main

import "fmt"

type Vertex struct {
    X int
    Y int
}

func main() {
    v := Vertex{1, 2}
    fmt.Println(v)
    v.X = 4
    fmt.Println(v)

    p := &v
    p.Y = 1000
    fmt.Println(v)

    // struct-literals
    var v1 = Vertex{1, 2}
    var v2 = Vertex{X: 1}
    var v3 = Vertex{}
    p = &Vertex{7, 8}
    fmt.Println(v1, v2, v3, p)
}
