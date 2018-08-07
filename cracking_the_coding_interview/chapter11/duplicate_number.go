package main

import "fmt"

func main() {
	a := []int{1, 10, 20323, 10, 20230, 20323}
	fmt.Println(checkDuplicates(a))
}

type BitSet struct {
	bitset []int
}

func (bs *BitSet) get(x int) bool {
	wordNumber := x >> 5
	bitNumber := uint(x & 0x1f)
	return (bs.bitset[wordNumber] & (1 << bitNumber)) != 0
}

func (bs *BitSet) set(x int) {
	wordNumber := x >> 5
	bitNumber := uint(x & 0x1f)
	bs.bitset[wordNumber] |= 1 << bitNumber
}

func checkDuplicates(arr []int) []int {
	results := make([]int, 0)
	bs := BitSet{bitset: make([]int, 1000, 1000)}
	for _, num := range arr {
		if bs.get(num - 1) {
			results = append(results, num)
		} else {
			bs.set(num - 1)
		}
	}
	return results
}
