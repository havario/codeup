// Copyright (c) 2025 honeok <honeok@disroot.org>

package main

import (
	"fmt"
)

func main() {
	// 使用while的方式输出10句hello world

	// 循环变量初始化
	i := 0
	for {
		if i > 10 {
			break
		}
		fmt.Printf("hello,world! current: %v\n", i)
		i++ // 循环迭代
	}

	fmt.Printf("exit var i = %v\n", i)

	// do...while 的实现
	j := 0
	for {
		fmt.Printf("hello ok! %v\n", j)
		j++ // 循环迭代
		if j > 10 {
			break
		}
	}
}
