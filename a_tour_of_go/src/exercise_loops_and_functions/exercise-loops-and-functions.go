package main

import (
    "fmt"
)

func Sqrt(x float64) float64 {
    z := float64(1)
    for (z/2-x/z/2) > 0.001 || (z/2-x/z/2) < -0.001 {
        z = z/2 + x/z/2
    }
    return z
}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(Sqrt(4))
}

