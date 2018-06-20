package main

import s "strings"
import "fmt"

var p = fmt.Println

func main() {
    p("Contains:", s.Contains("test", "es"))
    p()
    p("len:", len("hello"))
}
