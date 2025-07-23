// Copyright 2025 (c) honeok <i@honeok.com>

package main

import (
	"fmt"
)

func swap(n1 *int, n2 *int) {
	// 定义一个临时变量t
	t := *n1
	*n1 = *n2
	*n2 = t
}

func main() {
	a := 10
	b := 20
	swap(&a, &b) // 传入地址
	fmt.Printf("a=%v\nb=%v\n", a, b)
}
