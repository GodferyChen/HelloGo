package main

import "fmt"

/* 结构体字段可以通过结构体指针来访问。通过指针间接的访问是透明的 */
type Bean struct {
	X int
	Y int
}

func main() {
	v := Bean{3,4}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}
