package main

import "fmt"

/*对slice切片*/
func main() {
	s := []int{2,3,5,7,9,11}
	fmt.Println("s = ",s)
	fmt.Println("s[1:4]",s[1:4])

	//省略下标代表从0开始
	fmt.Println("s[:3]",s[:3])

	//省略上标代表到len(s)结束
	fmt.Println("s[4:]",s[4:])
}
