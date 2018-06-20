package main

import "fmt"
import "time"

func main() {
    p := fmt.Println

    t := time.Now()
    p(t.Format(time.RFC3339))

    t1, _ := time.Parse(
        time.RFC3339,
        "2012-11-01T22:08:41+00:00")
    p(t1)

    p(t.Format("3:04PM"))
    p(t.Format("Mon Jan _2 15:04:05 2007"))
    p(t.Format("2010-01-01T22:23:43.999999-07:00"))

    form := "3 04 PM"
    t2, _ := time.Parse(form, "8 41 PM")
    p(t2)
}
