package main

import "golang.org/x/tour/reader"
import "strings"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (mr MyReader) Read(b []byte) (int, error) {
    r := strings.NewReader("A")
    n, err := r.Read(b)
    return n, err
}

func main() {
    reader.Validate(MyReader{})
}
