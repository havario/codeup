// Copyright (c) 2025 honeok <honeok@disroot.org>

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

	// var score int
	// fmt.Printf("请输入岳云鹏期末成绩: ")
	// fmt.Scanln(&score)
	// if score > 100 {
	// 	fmt.Println("你有点抽象了!")
	// } else if score == 100 {
	// 	fmt.Println("奖励BMW!")
	// } else if score > 80 && score <= 99 {
	// 	fmt.Println("奖励iPhone7Plus!")
	// } else if score >= 60 && score <= 80 {
	// 	fmt.Println("奖励iPad!")
	// } else {
	// 	fmt.Println("废物, 什么都没有!")
	// }

	// 男大当婚女大当嫁, 需求: 高 180cm以上
	// 如果三个条件同时满足, 则: 一定要嫁人
	// 如果条件为真, 则：嫁吧, 比上不足比下有余
	// 如果三个条件都不满足, 则 你爱你妈逼嫁不嫁
	var height int32
	var money float32
	var handsome bool
	fmt.Println("请输入身高")
	fmt.Scanln(&height)
	fmt.Println("请输入财富")
	fmt.Scanln(&money)
	fmt.Println("请输入颜值")
	fmt.Scanln(&handsome)
	if height > 180 && money >= 1.0 && handsome {
		fmt.Println("一定嫁")
	} else if height > 180 || money >= 1.0 || handsome {
		fmt.Println("比上不足比下有余")
	} else {
		fmt.Println("你爱你妈逼嫁不嫁")
	}
}
