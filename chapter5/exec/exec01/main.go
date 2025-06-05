// Copyright (c) 2025 honeok <honeok@autistici.org>

package main

import (
	"fmt"
)

func main() {
	var n1 int32 = 10
	var n2 int32 = 50
	if n1+n2 >= 50 {
		fmt.Println("hello,world")
	}

	// 声明两个float64变量并赋值. 判断第一个数大于10.0
	// 且第二个数小于20.0 打印两数之和
	var n3 float64 = 11.0
	var n4 float64 = 17.0
	if n3 > 10.0 && n4 < 20.0 {
		fmt.Println("和=", (n3 + n4))
	}

	// 定义两个变量int32, 判断两者之和, 是否能被3又能被5整除,打印提示信息
	var n5 int32 = 5
	var n6 int32 = 10
	if (n5+n6)%3 == 0 && (n5+n6)%5 == 0 {
		fmt.Println("能被3又能被5整除")
	}

	// 判断一个年份是否是闰年, 闰年的条件符合下面二者之一
	// 1> 年份能被4整除, 但不能被100整除
	// 2> 能被400整除
	var year int = 2020
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		fmt.Printf("%v是闰年\n", year)
	}
}
