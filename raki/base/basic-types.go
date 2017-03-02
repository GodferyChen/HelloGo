package main

import (
	"fmt"
)

var(
	ToBe bool = false
	MaxInt uint64 = 1 << 64 - 1
	z complex128 = 7+10i
)

/* 基本类型 const 关键字用于声明常量 */
func main() {
	const f = "%T(%v)\n"
	fmt.Printf(f,ToBe,ToBe)
	fmt.Printf(f,MaxInt,MaxInt)
	fmt.Printf(f,z,z)

	var num  = 3.7E+1 + 5.98E-2i
	fmt.Printf("浮点数 %E 表示的是 %f。 \n",num,num)
}
