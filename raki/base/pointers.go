package main

import "fmt"

/* 指针保存了变量的内存地址。与C不同，Go没有指针运算 */
func main() {
	i, j := 42, 2701

	p := &i         //point to i
	fmt.Println(*p) //read i through the pointer
	*p = 21         //set i through the pointer
	fmt.Println(i)  //see the new value of i

	p = &j         //point to j
	*p = *p / 37   //divide j through the pointer
	fmt.Println(j) //see the new value of j

}
