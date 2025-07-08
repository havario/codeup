// Copyright (c) 2025 honeok <honeok@disroot.org>
//                           <i@honeok.com>

package main

import (
	"fmt"
)

/*
在go中, 函数也是一种数据类型
可以赋值给一个变量, 则该变量就是一个函数类型的变量, 通过该变量可以对函数调用
*/

func getSum(n1 int, n2 int) int {
	return n1 + n2
}

// 函数既然是一种数据类型, 因此在go中, 函数可以作为形参, 并且调用
func anyway(interesting func(int, int) int, num1 int, num2 int) int {
	return interesting(num1, num2)
}

func main() {
	// 函数赋值给一个变量a
	// a为一个函数变量
	a := getSum
	fmt.Printf("a的类型%T, getSum类型是%T\n", a, getSum)

	result1 := a(10, 40)
	fmt.Printf("result = %v\n", result1)

	result2 := anyway(getSum, 50, 60)
	fmt.Printf("result = %v\n", result2)
}
