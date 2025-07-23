// Copyright 2025 (c) honeok <i@honeok.com>

// 匿名函数

package main

import (
	"fmt"
)

var (
	// fun1就是全局匿名函数, 在整个程序有效
	Fun1 = func(n1 int, n2 int) int {
		return n1 * n2
	}
)

func main() {
	// 在定义函数时候就直接调用, 这种方式匿名函数只能调用一次

	// 求两个数的和, 使用匿名函数的方式⭐
	result1 := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Printf("result1 = %v\n", result1)

	// 将匿名函数func(n1 int, n2 int) int赋给变量a
	// 则a的数据类型为函数类型, 此时可以通过a完成调用
	a := func(n1 int, n2 int) int {
		return n1 - n2
	}
	result2 := a(10, 30)
	fmt.Printf("result2 = %v\n", result2)

	result3 := a(90, 30)
	fmt.Printf("result3 = %v\n", result3)

	// 全局匿名函数的使用
	result4 := Fun1(4, 9)
	fmt.Printf("result4 = %v\n", result4)
}
