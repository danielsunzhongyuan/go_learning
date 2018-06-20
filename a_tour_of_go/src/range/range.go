package main

import "fmt"
import "golang.org/x/tour/pic"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func Pic(dx, dy int) [][]uint8 {
    ret := make([][]uint8, dy)
    for i := 0; i < dy; i++ {
        ret[i] = make([]uint8, dx)
    }
    for j := 0; j < dy; j++ {
        for i := 0; i < dx; i++ {
            ret[i][j] = uint8(i*j)
        }
    }
    return ret
}

func main() {
    for i, v := range pow {
        fmt.Printf("2**%d = %d\n", i, v)
    }

    pow2 := make([]int, 10)

    for i := range pow2 {
        pow2[i] = 1 << uint(i)
    }

    for _, value := range pow2 {
        fmt.Printf("%d\n", value)
    }

    pic.Show(Pic)
}
