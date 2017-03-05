package main

import "fmt"

/* 可以通过赋值给_来忽略序号和值 */
func main() {
	pow := make([]int,10)
	for i := range pow{
		pow[i] = 1 << uint(i)
	}
	for _,value := range pow {
		fmt.Printf("%d\n",value)
	}
}
