package main

import (
	"fmt"
)

type People interface {
	Speak(string) string
}

type Stduent struct{}

func (stu *Stduent) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func main() {
	// it must be a ptr, Student does not implement the interface "People"
	var peo People = &Stduent{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}
