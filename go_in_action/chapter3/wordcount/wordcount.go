package main

import (
	"../words"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	filename := os.Args[1]

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("There was an error opening the file:", err)
		return
	}

	text := string(contents)
	count := words.CountWords(text)

	fmt.Printf("There are %d words in your text. \n", count)

	var a = make([]int, 3, 5)
	a = []int{1, 2, 3, 4, 5}
	b := a[1:3]
	b[0] = 100
	fmt.Println(a, b, len(b), cap(b))
}
