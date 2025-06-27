// Copyright (c) 2025 honeok <honeok@disroot.org>

package main

import (
	"fmt"
)

func test(n int) {
	if n > 2 {
		n--
		test(n)
	}
	fmt.Printf("n = %v\n", n)
}

func test2(n int) {
	if n > 2 {
		n-- // 递归必须向退出的条件逼近否则无限循环调用
		test2(n)
	} else {
		fmt.Printf("n = %v\n", n)
	}
}

// 递归调用的重要原则:
// 1> 执行一个函数时, 就创建一个新的受保护的独立空间(新函数栈)
// 2> 函数的局部变量是独立的, 不会相互影响
// 3> 递归必须向退出的条件逼近, 否则就是无限递归, 死龟了 :)
// 4> 当一个函数执行完毕 或者遇到return, 就会返回, 遵守谁调用, 就将结果返回给谁, 同时当函数执行完毕或者返回时该函数本身也会被系统销毁

func main() {

	test(4)
	test2(4)
}
