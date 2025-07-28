// Copyright 2025 (c) honeok <i@honeok.com>

package main

import (
	"fmt"
)

// 函数外部定义的变量叫全局变量
// 作用域在整个包都有效, 如果首字母为大写, 则作用域在整个程序有效
var (
	age  int    = 50      // 在别的包无法使用
	Name string = "zzwsb" // 在别的包可以使用, 首字母大写全局程序生效
)

// func localVar() {
// 	// age 和 Name的作用域只在localVar函数内部可用
// 	age := 10
// 	Name := "Tom"
// }

func main() {
	fmt.Println("age =", age)
	fmt.Println("Name =", Name)

	// 如果变量在一个代码块中, 比如for/if中, 那么这个变量的作用域就在该代码块
	for i := 0; i <= 10; i++ {
		fmt.Println("i=", i)
	}

	// undefine
	//fmt.Println("i=", i)
}
