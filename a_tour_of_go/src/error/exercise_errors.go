package main

import (
    "fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
    return "cannot Sqrt negative number: %v", e
}

func Sqrt(x float64) (float64, error) {
    if x < 0 {
        return 0, ErrNegativeSqrt(x)
    } else {
        ret := 1.0
        for i := 0; i < 10; i++ {
            ret = ret - (ret * ret - x) / 2 / ret
        }
        return ret, nil
    }
}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(Sqrt(-2))
}
