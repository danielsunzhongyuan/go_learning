package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	x := "azAZ0123456789"
	xBytes := []byte(x)
	fmt.Println(xBytes)
	xBytes[0] = 'T'
	fmt.Println(string(xBytes))

	xRune := []rune(x)
	fmt.Println(xRune)

	fmt.Printf("%T %T %T\n", x[0], xBytes[0], xRune[0])

	fmt.Println(utf8.ValidString("ABC"))
	fmt.Println(utf8.ValidString("A\xfeC"))

	data := "♥"
	fmt.Printf("len of %v is %d\n", data, len(data))
	fmt.Printf("utf8 len of %v is %d\n", data, utf8.RuneCountInString(data))

	{
		data := "A\xfe\x02\xff\x04"
		for _, v := range data {
			fmt.Printf("%#x ", v)
		}
		//prints: 0x41 0xfffd 0x2 0xfffd 0x4 (not ok)

		fmt.Println()
		for _, v := range []byte(data) {
			fmt.Printf("%#x \n", v)
		}
		//prints: 0x41 0xfe 0x2 0xff 0x4 (good)
	}
	{
		a := " "
		switch a {
		case " ":
			fallthrough // this is necessary
		case "\t":
			println("yyy")
		case "\n", "\b": // or put options together
			fmt.Println("zzz")
		}
	}
	println("zzz")
}
