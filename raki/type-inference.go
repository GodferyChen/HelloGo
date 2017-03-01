package main

import "fmt"

/* 类型推导：在定义变量却不显式指定其类型时，变量的类型由（等号）右侧的值推导出来 */
func main() {
	v := false //change me
	fmt.Printf("v is of type %T\n",v)
}
