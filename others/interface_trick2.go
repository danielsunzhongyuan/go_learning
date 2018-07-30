package main

import (
	"fmt"
)

type People interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {
}

func live() Student{
	var stu Student
	return stu
}

func main() {
	x := live()
	fmt.Println(x)
    

    a := new(Student)
	if a == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}
   
}
