// Copyright (c) 2025 honeok <honeok@disroot.org>
//                           <i@honeok.com>

package main

import (
	"fmt"
)

/*
猴子吃桃子的问题:
有一堆桃子, 猴子第一天吃了其中的一半, 并再多吃一个! 以后每天猴子都吃其中的一半, 然后再多吃一个. 当到第十天时, 想再吃时(还没吃), 发现只有1个桃子.
问: 最初共多少个桃子?

分析:
1> 第10天只有一个桃子
2> 第9天 优几个桃子 = (第10天桃子数量+1)*2
3> 规律: 第n天的桃子数据 peach(n)=(peach(n+1)+1)*2
*/

// n 范围 1~10之间
func peach(n int) int {
	if n > 10 || n < 1 {
		fmt.Printf("输入的天数错误\n")
		return 0 // 一个默认值 返回0表示没有得到正确数量
	}
	if n == 10 {
		return 1 //第10天只有一个桃子
	} else {
		return (peach(n+1) + 1) * 2
	}
}

func main() {
	fmt.Printf("第1天的桃子数量是: %v\n", peach(1))
	fmt.Printf("第10天的桃子数量是: %v\n", peach(10))
}
