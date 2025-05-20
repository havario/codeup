// Copyright (c) 2025 honeok <honeok@duck.com>

package main

import "fmt"

// golang中标识符的使用
func main() {

	// golang中严格区分大小写
	// goalng中认为num和Num是不同的变量
	var num int = 10
	var Num int = 20
	fmt.Printf("num=%v Num=%v\n", num, Num)

	// 标识符不能包含空格
	// var ab c int = 30

	// _ 空标识符, 用于占位
	// var _ int = 40 //error
	// fmt.Println(_)
}
