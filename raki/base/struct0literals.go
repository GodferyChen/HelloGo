package main

import "fmt"

type Vertex2 struct {
	X, Y int
}

var (
	v1 = Vertex2{1, 2}  //类型为Vertex
	v2 = Vertex2{X: 1}  //Y:0被省略
	v3 = Vertex2{}      //X:0和Y:0
	p  = &Vertex2{1, 2} //类型为 *Vertex
)

func main() {
	fmt.Println(v1, v2, v3, p)
}
