package main

import (
	"math"
	"fmt"
)

type Vertex4 struct {
	X, Y float64
}

func (v *Vertex4) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex4) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex4{3, 4}
	fmt.Printf("Before scaling:%v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling:%v, Abs: %v\n", v, v.Abs())
}
