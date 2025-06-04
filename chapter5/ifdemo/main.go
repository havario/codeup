package main

import (
	"fmt"
)

func main() {
	// 编写一个程序, 可以输入人的年龄, 如果该同志的年龄大于18岁, 则输入 "你的年龄大于18岁, 要对自己的行为负责"

	var age byte
	fmt.Println("请输入年龄: ")
	fmt.Scanln(&age)

	if age > 18 {
		fmt.Println("你的年龄大于18岁, 要对自己的行为负责!")
	}
}
