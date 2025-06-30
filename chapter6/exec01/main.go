// Copyright (c) 2025 honeok <honeok@disroot.org>
//                           <i@honeok.com>

package main

import (
	"fmt"
)

/*
使用递归的方式, 求出斐波那契数 1,1,2,3,4,5,6,7,8...
给你一个整数n, 求出它的斐波那契数是多少?

思路:
1> 当 n = 1 || n == 2 返回1
2> 当 n >= 2 返回前面两个数的和 f(n-1) + f(n-2)
*/

func fibonacci(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

func main() {
	fmt.Println("result = ", fibonacci(3))
	fmt.Println("result = ", fibonacci(4))
	fmt.Println("result = ", fibonacci(5))
	fmt.Println("result = ", fibonacci(6))
}
