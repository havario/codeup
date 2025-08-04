// Copyright 2025 (c) honeok <i@honeok.com>

// 日期和时间相关函数的方法使用

package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. 获取当前时间
	now := time.Now()
	fmt.Printf("now=%v now type=%T\n", now, now)

	fmt.Println("")

	// 2. 通过Now可以获取年月日 时分秒
	fmt.Printf("年=%v\n", now.Year())
	fmt.Printf("月=%v\n", int(now.Month()))
	fmt.Printf("日=%v\n", now.Day())
	fmt.Printf("时=%v\n", now.Hour())
	fmt.Printf("分=%v\n", now.Minute())
	fmt.Printf("秒=%v\n", now.Second())

	fmt.Println("")

	// 3. 格式化日期
	// 方式1: 使用SPrintf 或者 Printf
	fmt.Printf("当前年月日 %d-%d-%d %d:%d:%d \n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	dateStr := fmt.Sprintf("当前年月日 %d-%d-%d %d:%d:%d \n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Printf("dateStr=%v\n", dateStr)

	// 方式2: 使用time.Format方法完成
	fmt.Println(now.Format("2006/01/02 15:04:05"))
	fmt.Println()
	fmt.Println(now.Format("2006/01/02"))
	fmt.Println()
	fmt.Println(now.Format("15:04:05"))
	fmt.Println()
}
