// Copyright (c) 2025 honeok. All rights reserved.

package main

import (
	"fmt"
)

// 演示golang中字符类型使用
func main() {
	var c1 byte = 'a'
	var c2 byte = '0' // 字符0

	// 当直接输出byte值时, 就是输出了对应的字符Ascii码
	fmt.Println("c1 = ", c1)
	fmt.Println("c2 = ", c2)
	// 如果希望输出对应字符, 需要使用格式化输出即可
	fmt.Printf("c1=%c c2=%c", c1, c2)
}