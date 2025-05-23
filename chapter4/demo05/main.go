// Copyright (c) 2025 honeok <honeok@duck.com>

package main

import (
	"fmt"
)

func test() int {
	return 90
}

func main() {
	// 赋值运算符的使用
	var i int
	i = 10 // 基本赋值
	fmt.Printf("基本赋值 i=%+v\n", i)

	// 有两个变量, a 和 b 要求将其进行交换, 最终打印结果
	// a =9 b = 2   -->   a = 2 b = 9
	a := 9
	b := 2
	fmt.Printf("交换前的情况是 a=%+v, b=%+v\n", a, b)
	// 定义一个临时变量
	tempVar := a
	a = b
	b = tempVar
	fmt.Printf("交换后的情况是 a=%+v, b=%+v\n", a, b)

	// 符合赋值的操作
	a += 17 // 等价于a = a + 7
	fmt.Printf("%+v\n", a)

	var c int
	c = a + 3
	fmt.Println(c)

	// 赋值运算符的左边 只能是变量 右边可以是变量 表达式 常量值
	// 表达式: 任何有值都可以看做表达式
	var d int
	d = a //
	fmt.Println(d)
	d = 8 + 2*8 // = 的右边是表达式
	fmt.Println(d)
	d = test() + 90 // = 的右边是表达式
	fmt.Println(d)
	d = 890 // 常量
	fmt.Println(d)
}
