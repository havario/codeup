// Copyright (C) 2025 honeok <honeok@duck.com>

package main

import "fmt"

// 变量使用注意事项
func main() {
	// 该区域的数据值可以再同一类型范围不断变化
	var i int = 10
	i = 30
	i = 50

	fmt.Println("i=", i)

	// 变量在同一个作用域(再一个函数或代码块)内不能同名
	// var i int = 59
	// i := 99
}