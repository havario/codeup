// Copyright 2025 (c) honeok <i@honeok.com>

// 闭包

package main

import (
	"fmt"
	"strings"
)

// // 累加器
// func AddUpper() func(int) int {
// 	// 仅第一次初始化
// 	var n int = 10
// 	var str = "hello"
// 	return func(x int) int {
// 		n = n + x
// 		str += string(x)
// 		fmt.Println("str=", str)
// 		return n
// 	}
// }

// func main() {
// 	// 使用累加器
// 	f := AddUpper()
// 	fmt.Println(f(1)) // 11
// 	fmt.Println(f(2)) // 13
// 	fmt.Println(f(3)) // 16
// }

// 闭包的最佳实践

// 1. 编写一个函数makeSuffix(suffix string) 可以接收一个文件后缀名（比如.jpg) 并返回一个闭包
// 2. 调用闭包可以传入一个文件名 如果该文件名没有指定的后缀 (比如.jpg) 则返回文件名.jpg 如果
// 3. 要求使用闭包的方式完成
// 4. strings.HasSuffix 该函数可以判断某个字符串是否有指定的后缀

func makeSuffix(suffix string) func(string) string {

	return func(name string) string {
		// 如果 name 没有指定的后缀则加上, 否则就返回原来的名字
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}

		return name
	}
}

func main() {
	// 测试makeSuffix能否正常使用
	// 返回一个闭包
	f := makeSuffix(".jpg")
	fmt.Println("文件名处理后: ", f("winter"))
	fmt.Println("文件名处理后: ", f("bird.jpg"))
}
