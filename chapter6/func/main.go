// Copyright (c) 2025 honeok <honeok@disroot.org>

package main

import (
	"fmt"
)

// 将计算的功能放到一个函数中
func cal(num1 float64, num2 float64, operator byte) float64 {
	var result float64
	switch operator {
	case '+':
		result = num1 + num2
	case '-':
		result = num1 - num2
	case '*':
		result = num1 * num2
	case '/':
		result = num1 / num2
	default:
		fmt.Printf("操作符号错误\n")
	}
	return result
}

func main() {

	// 输入两个数, 再输入一个运算符(+ - * /), 得到结果
	var num1 float64 = 5.3
	var num2 float64 = 2.3
	var operator byte = '-'

	// 调用函数
	result := cal(num1, num2, operator)
	fmt.Printf("结果为: %v\n", result)
}
