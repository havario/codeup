// Copyright (c) 2025 honeok <i@honeok.com>

package main

import (
	"fmt"
)

func main() {
	// 使用for循环完成下面的案例编写一个程序, 可以接受一个整数, 表示层数, 打印出金字塔
	// 打印金字塔

	/*
		规律:
		  *   1层1个* 规律: 2 * 层数 - 1  空格 2 规律: 总层数 - 当前层数
		 ***  2层3个* 规律: 2 * 层数 - 1  空格 1
		***** 3层5个* 规律: 2 * 层数 - 1  空格 0


		打印空心金字塔

		  *
		 * *   2层3个 * 规律:  2 * 层数 - 1  空格 1 规律 总层数-当前层数
		*****  3层5个 * 规律:  2 * 层数 - 1  空格 0 规律 总层数-当前层数

		分析: 给每一行打印 * 号时, 需要考虑是打印 * 还是 空格
		每层第一个和最后一个打印 * 其他就应该为空, 即输出空格
		还有一个例外情况, 最后一层是全部 *
	*/

	// 将层数做成一个变量
	var totalLevel int = 20

	// i 表示层数
	for i := 1; i <= totalLevel; i++ {

		// 在打印星号之前先打印空格
		for k := 1; k <= totalLevel-i; k++ {
			fmt.Printf(" ")
		}

		// j 表示每一层打印多少*
		for j := 1; j <= 2*i-1; j++ {
			if j == 1 || j == 2*i-1 || i == totalLevel {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
		}

		fmt.Println()
	}

	// 打印99乘法表
	// i 表示层
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v\t", j, i, j*i)
		}
		fmt.Println("")
	}
}
