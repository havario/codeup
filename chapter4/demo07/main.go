// Copyright (c) 2025 honeok <honeok@duck.com>

package main

import (
	"fmt"
)

// 要求: 可以从控制台接受用户信息 [ 姓名 年龄 薪水 是否通过考试 ]
func main() {

	// 方式1: fmt.fmt.Scanln
	// 1. 先声明需要的变量
	var name string
	var age byte
	var salary float32
	var isPass bool
	fmt.Println("请输入姓名")
	// 当程序执行到fmt.Scanln(&name)时, 程序会停止在这里, 等待用户输入并回车
	fmt.Scanln(&name)
	fmt.Println("请输入年龄")
	fmt.Scanln(&age)
	fmt.Println("请输入薪资")
	fmt.Scanln(&salary)
	fmt.Println("是否通过考试")
	fmt.Scanln(&isPass)

	fmt.Printf("名字是 %v \n 年龄是 %v \n 薪资 %v \n 是否通过考试 %v \n", name, age, salary, isPass)

	// 方式2: fmt.Scanf, 可以按指定的格式输入
	fmt.Println("请输入你的信息: 姓名 年龄 薪水 是否通过考试, 使用空格分割")
	fmt.Scanf("%s %d %f %t", &name, &age, &salary, &isPass)
	fmt.Printf("名字是 %v \n 年龄是 %v \n 薪资 %v \n 是否通过考试 %v \n", name, age, salary, isPass)
}
