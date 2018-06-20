package main

import "runtime"
//import "sync"
import "fmt"


func MaxParallelism() int {
	maxProcs := runtime.GOMAXPROCS(0)
	fmt.Printf("maxProcs: %d\n", maxProcs)
	numCPU := runtime.NumCPU()
	fmt.Printf("numCPU: %d\n", numCPU)
	if maxProcs < numCPU {
		return maxProcs
	}
	return numCPU
}

func main() {
	x := MaxParallelism()
	fmt.Println("x:", x)
}
	
