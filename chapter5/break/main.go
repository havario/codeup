// Copyright (c) 2025 honeok <honeok@disroot.org>

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// 为了生成一个随机数 还需要为rand设置一个种子
	// rand.Seed(time.Now().Unix())
	// 随机生成1~100的整数
	// n := rand.Intn(100) + 1
	// fmt.Println(n)

	var count int = 0
	for {
		rand.Seed(time.Now().UnixNano()) // 纳秒
		n := rand.Intn(100) + 1          // 默认是0~99
		fmt.Printf("n=%v\n", n)
		count++
		if n == 99 {
			break
		}
	}
	fmt.Println("生成99一共使用了", count)

lable2:
	// 演示一下break指定标签的形式来使用
	for i := 0; i < 4; i++ {
		// lable1: // 设置一个标签
		for j := 0; j < 10; j++ {
			if j == 2 {
				// break // break默认情况下会终止最近的for循环
				break lable2
			}
			fmt.Println("j=", j)
		}
	}
}
