package main

import "fmt"

/* 一个slice会指向一个序列的值，并且包含了长度信息 */
func main() {
	s := []int{2,3,5,7,9,11}
	fmt.Println("s = ",s)

	for i := 0; i < len(s);i++ {
		fmt.Printf("s[%d] = %d\n",i,s[i])
	}
}
