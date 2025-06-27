// Copyright (c) 2025 honeok <honeok@disroot.org>

package main

import (
	"fmt"

	C "github.com/havario/codeup/chapter6/func/utils"
)

func main() {

	// 输入两个数, 再输入一个运算符(+ - * /), 得到结果
	var num1 float64 = 5.3
	var num2 float64 = 2.3
	var operator byte = '-'
	// 调用函数
	result := C.Cal(num1, num2, operator)
	fmt.Printf("结果为: %v\n", result)

	num1 = 2
	num2 = 3
	operator = '*'
	// 调用函数
	result = C.Cal(num1, num2, operator)
	fmt.Printf("结果为: %v\n", result)
	C.SayOk()
}
