package utils

import (
	"fmt"
)

// 将计算的功能放到一个函数中
func Cal(num1 float64, num2 float64, operator byte) float64 {
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
