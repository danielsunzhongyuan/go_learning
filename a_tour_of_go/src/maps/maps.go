package main

import (
    "fmt"
    "reflect"
    "strings"
    "golang.org/x/tour/wc"
)

type Vertex struct {
    Lat, Long float64
}

var m map[string]Vertex

var n = map[string]Vertex{
    "Bell Labs": Vertex{
        40.68433, -74.39967,
    },
    "Google": Vertex{
        37.42202, -122.08408,
    },
}

// above definition, Vertex is redundant
var o = map[string]Vertex{
    "Bell Labs": {40.68433, -74.39967},
    "Google":    {37.42202, -122.08408},
}

func WordCount(s string) map[string]int {
    s_list := strings.Fields(s)
    //ret := make(map[string]int)
    ret := map[string]int {}
    for _, word := range s_list {
        ret[word] = ret[word] + 1
        //tmp, ok := ret[word]
        //if ok {
        //    ret[word] = tmp + 1
        //} else {
        //    ret[word] = 1
        //}
    }
    return ret
}

func main() {
    m = make(map[string]Vertex)
    m["bell"] = Vertex{
        40.3, -23.03,
    }
    fmt.Println(m["bell"])
    fmt.Println(reflect.TypeOf(n))
    fmt.Println(o)

    p := make(map[string]int)
    p["answer"] = 42
    fmt.Println("The value:", p["answer"])

    p["answer"] = 48
    fmt.Println("The value:", p["answer"])

    delete(p, "answer")
    fmt.Println("The value:", p["Answer"])

    v, ok := p["answer"]
    fmt.Println("The value:", v, "Present?", ok)


    wc.Test(WordCount)
}
