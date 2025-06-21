// Copyright (c) 2025 honeok <honeok@disroot.org>

package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("你好, 我是你爸 %v\n", i)
		fmt.Printf("i=%v\n", i)
		fmt.Println(&i)
	}

	// for循环的第二种写法
	j := 1       // 循环变量初始化
	for j < 10 { // 循环条件
		fmt.Printf("你好, 我是你爷 %v\n", j)
		j++
	}

	// for循环的第三种写法
	k := 1
	for {
		if k <= 10 {
			fmt.Printf("你好, 我是你祖 %v\n", k)
		} else {
			break
		}
		k++
	}
}
