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
		rand.Seed(time.Now().Unix())
		n := rand.Intn(100) + 1
		count++
		if n == 99 {
			break
		}
	}
	fmt.Println("生成99一共使用了", count)
}
