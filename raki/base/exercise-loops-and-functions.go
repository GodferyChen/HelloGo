package main

import (
	"fmt"
	"math"
)

/* 牛顿法计算平方根 */
func Sqrt(x float64) float64 {
	z := 1.0
	for {
		tmp := z - (z*z -x)/(2*z)
		fmt.Println(tmp)
		if tmp == z || math.Abs(tmp-z) < 0.000000000001 {
			break
		}
		z = tmp
	}
	return z
}

func main() {
	fmt.Println(Sqrt(4))
	fmt.Println(math.Sqrt(4))
}
