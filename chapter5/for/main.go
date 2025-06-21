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

	// 字符串遍历 方式1
	// var str string = "hello, world!"
	// for i := 0; i < len(str); i++ {
	// 	fmt.Printf("%c\n", str[i])
	// }

	// 字符串遍历 方式1
	var str string = "hello, world!乌克兰"
	change := []rune(str) // 就是把str转为 []rune
	for i := 0; i < len(change); i++ {
		fmt.Printf("%c\n", change[i])
	}

	// 字符串遍历 方式2 for range
	// for-range变量字符串是按照字符进行遍历, 而不是字节
	str = "wocaonima!俄罗斯"
	for index, value := range str {
		fmt.Printf("index=%v value=%c \n", index, value)
	}
}
