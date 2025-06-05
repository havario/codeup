// Copyright (c) 2025 honeok <honeok@autistici.org>

package main

import (
	"fmt"
)

func main() {
	var age int
	fmt.Printf("请输入你的年龄: ")
	fmt.Scanln(&age)

	if age >= 18 {
		fmt.Println("你已经成年, 要对自己的行为负责!")
	} else {
		fmt.Println("你的年龄不大这次就放过你了.")
	}
}
