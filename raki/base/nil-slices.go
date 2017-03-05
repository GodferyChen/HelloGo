package main

import "fmt"

/* 一个nil的slice的长度和容量是0 */
func main() {
	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}
}
