// Copyright (c) 2025 honeok <honeok@duck.com>

package main

import (
	"fmt"
)

// 演示一把 & 和 * 的使用
func main() {
	a := 100
	fmt.Println("a的地址是:", &a)

	var ptr *int = &a
	fmt.Println("ptr指向的值", *ptr)

	var n int
	var i int = 10
	var j int = 12
	// 传统的三元运算
	// n = i > j ? i : j
	if i > j {
		n = i
	} else {
		n = j
	}
	fmt.Println("n=", n)
}
