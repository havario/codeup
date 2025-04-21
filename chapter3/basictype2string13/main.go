// Copyright (c) 2025 honeok <honeok@duck.com>

package main

import (
	"fmt"
	"strconv"
	_ "unsafe"
)

// 基本数据类型转string
func main() {
	var num1 int = 99
	var num2 float64 = 23.456
	var b bool = true
	var ischar byte = 'h'
	var str string

	// 使用第一种方式来转换
	str = fmt.Sprintf("%d", num1)
	fmt.Printf("string type is: %T str=%q\n", str, str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("string type is: %T str=%q\n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("string type is: %T str=%q\n", str, str)

	str = fmt.Sprintf("%c", ischar)
	fmt.Printf("string type is: %T str=%q\n", str, str)

	// 使用strconv包函数
	var num3 int = 99
	var num4 float64 = 23.456
	var b2 bool = true

	str = strconv.FormatInt(int64(num3), 10)
	fmt.Printf("string type is: %T str=%q\n", str, str)

	// 说明:
	// f: 格式
	// 10: 表示小数位保留10位
	// 64: 表示小数是float64
	str = strconv.FormatFloat(num4, 'f', 10, 64)
	fmt.Printf("string type is: %T str=%q\n", str, str)

	str = strconv.FormatBool(b2)
	fmt.Printf("string type is: %T str=%q\n", str, str)

	// strconv包还有一个函数 (Itoa)
	var num5 int64 = 4567
	str = strconv.Itoa(int(num5))
	fmt.Printf("string type is: %T str=%q\n", str, str)
}
