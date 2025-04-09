// Copyright (c) 2025 honeok <honeok@duck.com>

package main

import (
	"fmt"
)

// byte 范围 0 ~ 255
// 演示golang中字符类型使用
func main() {
	var c1 byte = 'a'
	var c2 byte = '0' // 字符0

	// 当直接输出byte值时, 就是输出了对应的字符Ascii码
	fmt.Println("c1 = ", c1)
	fmt.Println("c2 = ", c2)
	// 如果希望输出对应字符, 需要使用格式化输出即可
	fmt.Printf("c1 = %c c2 = %c\n", c1, c2)

	// var c3 byte = '北' byte value in variable declaration (overflows)
	var c3 int = '北'
	fmt.Printf("c3 = %c c3对应的码值 = %d\n", c3, c3)
}

// 说明
// 1> 如果报错的字符在ASCII表中, 比如[0-1,a-z,A-Z..] 可以直接保存byte
// 2> 如果报错的码值大于255, 这时可以考虑使用int类型
// 3> 如果需要安装字符的方式输出, 需要格式化输出, 即 fmt.Printf("%c, x)