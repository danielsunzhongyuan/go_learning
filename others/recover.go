package main

import "fmt"

func doRecover() {
	fmt.Println("recovered => ", recover())
}

// this will fail recovering.
// recover() must be called in the "defer function"
func main2() {
	defer func() {
		doRecover()
	}()
	panic("not good")
}

// this works, but will print "recovered: <nil>" if there is no panic
func main1() {
	defer func() {
		fmt.Println("recovered:", recover())
	}()
	//panic("not good")
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered:", err)
		}
	}()
	//panic("not good")
}
