// Copyright (c) 2025 honeok <honeok@disroot.org>

package main

import (
	"fmt"
)

func main() {
	// 参加百米运动会, 如果用时8秒以内进入决赛
	// 否则提示淘汰, 并且根据性别提示进入男子组或女子组
	// 输入成绩和性别

	var second float64
	fmt.Println("请输入秒数")
	fmt.Scanln(&second)

	if second <= 8 {
		// 进入决赛
		var gender string
		fmt.Println("请输入性别")
		fmt.Scanln(&gender)
		if gender == "male" {
			fmt.Println("进入决赛男子组")
		} else if gender == "female" {
			fmt.Println("进入决赛女子组")
		} else {
			fmt.Println("进入决赛中性人组")
		}
	} else {
		fmt.Println("out you!")
	}

	// 出票系统: 根据淡旺季月份和年龄 打印票价

	var month int32
	var age byte
	price := 60.0 // 旺季原始票价
	fmt.Printf("请输入旅游月份: ")
	fmt.Scanln(&month)
	fmt.Printf("请输入游客年龄: ")
	fmt.Scanln(&age)

	// 旺季
	if month >= 4 && month <= 10 {
		if age >= 60 {
			fmt.Printf("%v月 年龄: %v 票价: %v\n", month, age, price/3)
		} else if age >= 18 {
			fmt.Printf("%v月 年龄: %v 票价: %v\n", month, age, price)
		} else {
			// 儿童
			fmt.Printf("%v月 年龄: %v 票价: %v\n", month, age, price/2)
		}
	} else {
		// 淡季
		if age >= 18 && age < 60 {
			fmt.Printf("淡季成人票价: 40\n")
		} else {
			fmt.Printf("淡季儿童和老人票价: 20\n")
		}
	}
}
