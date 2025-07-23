// Copyright 2025 (c) honeok <i@honeok.com>

package main

import (
	"fmt"

	U "github.com/havario/codeup/chapter6/funcinit/utils"
)

var age = pre()

// 为了看到全局变量是先被初始化的, 这里先写函数完成初始化
func pre() int {
	fmt.Println("pre()")
	return 90
}

// init函数, 通常可以在init函数中完成初始化工作
func init() {
	fmt.Println("init()")
}

// 如果一个文件同时包含全局变量定义, init函数和main函数, 则执行的流程是 变量定义 -> init函数 -> main函数
func main() {
	fmt.Println("main(), age = ", age)
	fmt.Println("Age=", U.Age, "Name=", U.Name)
}
