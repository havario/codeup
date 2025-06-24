// Copyright (c) 2025 honeok <honeok@disroot.org>

package main

import (
	"fmt"
)

func main() {
	// 打印金字塔

	// 将层数做成一个变量
	var totalLevel int = 20

	// i表示层数
	for i := 1; i <= totalLevel; i++ {
		// 打印*前打印空格
		for k := 1; k <= totalLevel-i; k++ {
			fmt.Print(" ")
		}

		// j 表示每层打印多少个*
		for j := 1; j <= 2*i-1; j++ {
			if j == 1 || j == 2*i-1 || i == totalLevel {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Printf("\n")
	}

	// 打印99乘法表
	// i 表层
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v\t", j, i, j*i)
		}
		fmt.Println("")
	}
}
