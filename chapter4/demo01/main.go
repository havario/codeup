// Copyright (c) 2025 honeok <honeok@duck.com>

package main

import "fmt"

// golang的除法和取模 / and %
func main() {

	// golang的除法/
	// 说明: 如果运算的数都是整数, 那么除后, 去掉小数部分, 保留整数部分
	fmt.Println(10 / 4) // 2

	var n1 float32 = 10 / 4
	fmt.Println(n1) // 结果还是2

	// 如果希望保留小数部分, 则需要有浮点数参与运算
	var n2 float32 = 10.0 / 4
	fmt.Println(n2)

	// golang的取模%
	// 看一个公式  a % b = a - a / b * b
	fmt.Println("10%3=", 10%3)     // 1
	fmt.Println("-10%3=", -10%3)   // -1
	fmt.Println("10%-3=", 10%-3)   // 10 - 10 / -3 * -3
	fmt.Println("-10%-3=", -10%-3) // 1
}
