package main

import "fmt"

func main() {
	fmt.Println("Hello, world.")
	var v1 int
	var v2 string
	var v3 [10]int
	var v4 []int
	var v5 struct {
		f int
	}
	var v6 *int
	var v7 map[string]int
	var v8 func(a int) int

	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3)
	fmt.Println(v4)
	fmt.Println(v5)
	fmt.Println(v6)
	fmt.Println(v7)
	fmt.Println(v8)

	array := [5]int{11, 12, 13, 14, 15}
	fmt.Println(array)

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, 5, 10)
	fmt.Println(slice1)
	fmt.Println(slice2)

	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		numberOfDays
	)
	fmt.Println(Sunday)
	fmt.Println(Monday)
	fmt.Println(Saturday)
	fmt.Println(numberOfDays)

	x, y := 1, 2
	fmt.Println((x + y) / 2)
}
