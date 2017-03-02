package main

import "fmt"

func main() {
	//for循环
	sum := 0
	for i := 0; i < 10;i++  {
		sum += i
	}
	fmt.Println(sum)

	//初始化语句和后置语句都是可选的
	sum1 := 1
	for ; sum1 < 1000;  {
		sum1 += sum1
	}
	fmt.Println(sum1)

	//for是Go的“while”
	sum2 := 1
	for sum2<1000  {
		sum2 += sum2
	}
	fmt.Println(sum2)
	
	//死循环
	//for {
	//
	//}

}
