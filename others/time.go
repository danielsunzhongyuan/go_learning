package main

import "fmt"
import "time"

func main() {
    p := fmt.Println

    now := time.Now()
    p(now)

    then := time.Date(2015, 07, 01, 20, 30, 23, 672387123, time.UTC)
    p(then)

    p(then.Year())
    p(then.Month())
    p(then.Second())
    p(then.Nanosecond())
    p(then.Location())
    p(then.Weekday())

    p(then.Before(now))

    diff := now.Sub(then)
    p(diff)

    p(diff.Hours())

    p(then.Add(diff))
}
