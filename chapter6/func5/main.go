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

type myFunType func(int, int) int // 这时这个myFunType就是 func(int,int) int类型

// 函数既然是一种数据类型, 因此在go中, 函数可以作为形参, 并且调用
func anyway2(interesting myFunType, num1 int, num2 int) int {
	return interesting(num1, num2)
}

// 支持对返回函数命名
func getSumAndSub(n1 int, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return // 返回两个结果按顺序接收
}

func main() {
	// 函数赋值给一个变量a
	// a为一个函数变量
	a := getSum
	fmt.Printf("a的类型%T, getSum类型是%T\n", a, getSum)

	result1 := a(10, 40)
	fmt.Printf("result1 = %v\n", result1)

	result2 := anyway(getSum, 50, 60)
	fmt.Printf("result2 = %v\n", result2)

	// 举例说明自定义数据类型的使用
	// 给int取了别名，在go中myInt 和 int虽然都是 int 类型，（但是go认为myInt和int两个类型）
	type myInt int
	var num3 myInt
	var num4 int
	num3 = 40
	num4 = int(num3) // 这里依然需要显示转换
	fmt.Println("num3 = ", num3)
	fmt.Println("num4 = ", num4)

	result3 := anyway2(getSum, 500, 600)
	fmt.Printf("result3 = %v\n", result3)

	//
	e, f := getSumAndSub(1, 2)
	fmt.Printf("e = %v\n", e) // sum
	fmt.Printf("f = %v\n", f) // sub
}
