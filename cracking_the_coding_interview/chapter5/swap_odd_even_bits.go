package main

import "fmt"

/*
Write a program to swap odd and even bits in an integer with as few instructions as possible (e g , bit 0 and bit 1 are swapped, bit 2 and bit 3 are swapped, etc)

交换奇偶位的比特
*/

func SwapOddEvenBits(x int) int {
	return ((x&0xaaaaaaaa)>>1) | ((x&0x55555555)<<1)
}

func main() {
	x := 7
	fmt.Println(SwapOddEvenBits(x))
}
