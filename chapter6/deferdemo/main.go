// Copyright 2025 (c) honeok <i@honeok.com>

// 延时

package main

import (
	"fmt"
)

func sum(n1, n2 int) int {
	// 当执行到defer时, 暂时不执行, 会将defer后面的语句压入一个独立栈中, 姑且叫defer栈
	// 当函数执行完毕后再从defer栈按照 先入后出 的方式出栈, 执行
	defer fmt.Println("ok1 n1 =", n1)
	defer fmt.Println("ok2 n2 =", n2)

	result := n1 + n2 // result 30
	fmt.Println("ok3 result =", result)
	return result
}

func main() {
	result := sum(10, 20)
	fmt.Println("result", result)
}
