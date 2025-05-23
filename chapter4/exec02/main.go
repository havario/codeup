package main

import (
	"fmt"
)

func main() {
	// 面试题: 有两个变量, a 和 b 要求将其进行交换, 但是不允许使用中间变量, 打印最终结果
	var a int = 10
	var b int = 20

	a = a + b //
	b = a - b // b = a + b - b == > b = a
	a = a - b // a = a + b - a == > a = b
	fmt.Printf("a=%v b=%v\n", a, b)
}
