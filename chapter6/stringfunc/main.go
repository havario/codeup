// Copyright (c) 2025 honeok <i@honeok.com>
//
// 字符串中常用的系统函数

package main

import (
	"fmt"
)

func main() {
	// 统计字符串的长度, 按字节
	str1 := "hello北"                    // golang编码统一为utf-8, ascii码的字符(字母和数字) 占一个字节, 汉字占用3个字节
	fmt.Println("str len= ", len(str1)) // 打印8 5+3

	// 字符串遍历, 同时处理有中文的问题 r := []rune(str)
	str2 := "hello北京"
	r := []rune(str2)
	for i := 0; i < len(r); i++ {
		fmt.Printf("字符=%c\n", r[i])
	}
}
