package main

import (
	"fmt"
)

func main() {
	var a = 0.0
	var b = 0.0
	for {
		n, _ := fmt.Scan(&a, &b)
		if n == 0 {
			break
		} else {
			fmt.Printf("%.2f\n", a+b)
		}
	}
}
