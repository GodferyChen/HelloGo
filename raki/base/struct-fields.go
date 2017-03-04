package main

import "fmt"

/* 结构体字段使用点号来访问 */
type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{2,3}
	v.X = 4
	fmt.Println(v.X)
}
