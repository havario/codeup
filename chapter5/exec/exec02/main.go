// Copyright (c) 2025 honeok <honeok@autistici.org>

package main

import (
	"fmt"
)

func main() {
	// 岳云鹏参加golang考试他和父亲岳不群达成承诺
	// 如果
	// 成绩为100的时候 奖励BMW
	// 成绩为[80,99]的时候, 奖励iPhone7Plus
	// 成绩为[60,80]的时候,奖励iPad
	// 其余没有奖励
	// 从键盘获取岳云鹏期末成绩并加以判断

	var score int
	fmt.Printf("请输入岳云鹏期末成绩: ")
	fmt.Scanln(&score)
	if score > 100 {
		fmt.Println("你有点抽象了!")
	} else if score == 100 {
		fmt.Println("奖励BMW!")
	} else if score > 80 && score <= 99 {
		fmt.Println("奖励iPhone7Plus!")
	} else if score >= 60 && score <= 80 {
		fmt.Println("奖励iPad!")
	} else {
		fmt.Println("废物, 什么都没有!")
	}
}
