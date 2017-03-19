package main

import (
	"os"
	"fmt"
)

type Reader interface {
	Read(b []byte) (n int,err error)
}

type Write interface {
	Write(b []byte) (n int,err error)
}

type ReadWriter interface {
	Reader
	Write
}

func main() {
	var w Write

	w = os.Stdout

	fmt.Fprintf(w,"hello, writer\n")
}
