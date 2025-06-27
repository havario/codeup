// Copyright (c) 2025 honeok <honeok@disroot.org>

package main

import (
	"fmt"
)

// 编写一个函数
func test(n1 int) {
	n1 += 1
	fmt.Printf("func n1 = %v\n", n1)
}

// 一个函数 getSum
// ------------------------>int 表示返回参数列表
func getSum(n1 int, n2 int) int {
	sum := n1 + n2
	fmt.Printf("getSum sum= %v\n", sum)
	// 当函数有return的时候, 就是将结果返回给调用者
	// 即谁调用就返回给谁
	return sum
}

// 编写函数, 可以两个数的和和差, 并返回结果
func getSumAndSub(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}

// 调用一个函数的时候, 会给函数分配一个新的空间, 编译器会通过自身的处理让这个新的空间和其他的栈的空间区分开
// 在每个函数对应的栈中, 数据空间是独立的, 不会混淆
// 当一个函数执行完成后, 程序会销毁这个函数对应的栈空间

func main() {

	n1 := 10
	// 调用test函数
	test(n1)
	fmt.Printf("main n1 = %v\n", n1)

	sum := getSum(10, 20)
	fmt.Printf("main sum = %v\n", sum)

	// 调用函数 getSumAndSub
	result1, result2 := getSumAndSub(1, 2)
	fmt.Printf("result1 = %v  result2 = %v\n", result1, result2)

	// 希望忽略某个返回值, 则使用 _ 符号占位
	_, result3 := getSumAndSub(50, 100)
	fmt.Printf("result3 = %v\n", result3)
}
