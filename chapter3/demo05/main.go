// Copyright (C) 2025 honeok <honeok@duck.com>

package main

import "fmt"

// 演示go中+的使用
func main() {
	var i = 300
	var j = 311
	var r = i + j // 加法运算

	fmt.Println("r=", r)

	var str1 = "hello"
	var str2 = "world!"
	var streq = str1 + str2

	fmt.Println("streq=", streq) // 拼接操作
}