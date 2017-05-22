package main

import (
	"strconv"
	"fmt"
)

func main() {
	var grid = [][]int{
		{0, 1, 0, 0},
		{1, 1, 1, 0},
		{0, 1, 0, 0},
		{1, 1, 0, 0},
	}
	fmt.Println(islandPerimeter(grid))
}

func islandPerimeter(grid [][]int) int {
	res := 0
	n := len(grid)
	m := len(grid[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				res += 4
				if i > 0 && grid[i-1][j] == 1 {
					res -= 2
				}
				if j > 0 && grid[i][j-1] == 1 {
					res -= 2
				}
			}
		}
	}
	return res
}

/**
反转字符串
 */
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

/**
Fizz Buzz
 */
func fizzBuzz(n int) []string {
	s := make([]string, n)
	for i := 1; i <= len(s); i++ {
		if i%3 == 0 && i%5 == 0 {
			s[i-1] = "FizzBuzz"
		} else if i%3 == 0 {
			s[i-1] = "Fizz"
		} else if i%5 == 0 {
			s[i-1] = "Buzz"
		} else {
			s[i-1] = strconv.Itoa(i) //int to string method
		}
	}
	return s
}

/**
Number Complement
 */
func findComplement(num int) int {
	temp := num
	temp |= temp >> 1
	temp |= temp >> 2
	temp |= temp >> 4
	temp |= temp >> 8
	temp |= temp >> 16
	temp = temp - (temp >> 1)
	return int(uint32(^num) & ((uint32(temp) << 1) - 1))
}

/**
The Hamming distance between two integers is the number of positions at which the corresponding bits are different.
Given two integers x and y, calculate the Hamming distance.
解题思路：将两个数异或起来，遍历异或结果的每一位，统计为1的个数
 */
func hammingDistance(x int, y int) int {
	res := 0
	exc := x ^ y
	for i := 0; i < 32; i++ {
		res += int(uint32(exc)>>uint32(i)) & 1
	}
	return res
}
