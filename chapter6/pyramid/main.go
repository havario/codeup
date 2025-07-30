package main

import (
	"fmt"
)

// 将打印金字塔的代码封装函数
func printPyramid(totalLevel int) {
	// i 表示层数
	for i := 1; i <= totalLevel; i++ {

		// 在打印星号之前先打印空格
		for k := 1; k <= totalLevel-i; k++ {
			fmt.Printf(" ")
		}

		// j 表示每一层打印多少*
		for j := 1; j <= 2*i-1; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}

func main() {

	// 从终端输入一个整数打印出对应的金字塔
	var num int

	fmt.Printf("请输入打印金字塔的层数: ")
	fmt.Scanln(&num)

	// 调用printPyramid打印金字塔
	printPyramid(num)
}
