// Copyright (C) 2025 honeok <honeok@duck.com>

package main

import "fmt"

// 演示go中整数类型的使用
func main() {
	var i int = 1
	fmt.Println("i=", i)

	// 测试int8的范围 -128 ~ 127
	// int16 int32 int64 类推
	var j int8 = 127
	fmt.Println("j=", j)

	// 测试uint8的范围 0 ~ 255
	var k uint8 = 255
	fmt.Println("k=", k)

	// int uint rune byte 使用
	var a int = 8900
	var b uint = 1
	var c byte = 255
	fmt.Println("a=", a, "b=", b, "c=", c)
}