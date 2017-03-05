package main

import "fmt"

func main() {
	var a []int
	print("a",a)

	//append works on nil slices
	a = append(a,0)
	print("a",a)

	//the slice grows as needed
	a = append(a,1)
	print("a",a)

	//we can add more than one element at a time
	a = append(a,2,3,4)
	print("a",a)
}

func print(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",s,len(x),cap(x),x)
}
