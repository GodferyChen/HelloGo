package main

import (
	"math"
	"fmt"
)

type Abser interface {
	Abs() float64
}

type MFloat float64

type Vertex5 struct {
	X,Y float64
}

func (f MFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex5) Abs() float64 {
	return math.Sqrt(v.X*v.X+v.Y*v.Y)
}

func main() {
	var a Abser
	f := MFloat(-math.Sqrt2)
	v := Vertex5{3,4}

	a = f
	a = &v

	a = v

	fmt.Println(a.Abs())
}
