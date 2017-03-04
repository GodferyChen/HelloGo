package main

import "fmt"

/* 一个结构体（struct）就是一个字段的集合 */
type Entry struct {
	X float64
	Y string
	Z bool
}

func main() {
	fmt.Println(Entry{2,"hello",true})
}
