package main

import "fmt"

type PersonInfo struct {
    ID string
    Name string
    Address string
}

func test(x int) int {
    if x == 0 {
        return 5
    } else {
        return x
    }
}

func test_multiple_args(args ...int) {
    test_multiple_args2(args...)
    test_multiple_args2(args[1:]...)
}

func test_multiple_args2(args ...int) {
    for _, arg := range args {
        fmt.Println(arg)
    }
}

func main() {
    var personDB map[string] PersonInfo
    personDB = make(map[string] PersonInfo)

    personDB["12345"] = PersonInfo{"12345", "Tom", "Room 203,..."}
    personDB["1"] = PersonInfo{"1", "Jack", "Room 101,..."}
    person, ok := personDB["12345"]

    if ok {
        fmt.Println("Found person", person.Name, "with ID 1234.")
    } else {
        fmt.Println("Did not find person with ID 1234.")
    }

    fmt.Println(test(0))
    fmt.Println("multiple_args")
    test_multiple_args(1,2,3,4,5)
}

