package main

import "fmt"

func main() {
	var i uint

	// 2 mistakes here
	//for i = 100; i <= 0; i-- {
	//	fmt.Printf("%d\n", i)
	//}

	// correction
	for i = 100; i > 0; i-- {
		fmt.Printf("%d\n", i)
	}
}
