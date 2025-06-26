package main

import (
	"fmt"
)

func main() {
	// continue实现 打印1~100之内的奇数 要求使用for循环+continue
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("奇数是: %v\n", i)
	}

	// 从键盘读入个数不确定的整数, 并判断读入的正数和负数的个数, 输入为0时结束程序 (使用for循环, break, continue完成)
	// positive  negative
	var positiveCount int = 0 // 统计正数个数
	var negativeCount int = 0 // 统计负数个数
	var num int
	for {
		fmt.Printf("请输入一个整数: ")
		fmt.Scanln(&num)
		if num == 0 {
			break
		}
		if num > 0 {
			positiveCount++
			continue
		}
		negativeCount++
	}
	fmt.Printf("正数个数是: %v 负数的个数是: %v\n", positiveCount, negativeCount)

	// 某人有100000元, 每经过一次路口需要交费, 规则如下
	// 当现金 > 50000元时, 每次交5%
	// 当现金 < 50000元时, 每次交1000
	// 计算该人可以经过多少次路口, 使用 for break 完成
	principal := 100000.0 // 总金额
	passCount := 0        // 通行计数器
	for {
		if principal > 50000 {
			also := principal * 0.05
			principal -= also
		} else if principal >= 1000 {
			principal -= 1000
		} else {
			break
		}
		passCount++
	}
	fmt.Printf("经过%d次路口\n", passCount)
}
