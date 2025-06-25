package main

import (
	"fmt"
)

func main() {
	// 100以内的数求和, 求出当和第一次大于20的当前数

	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i // 求和
		if sum > 20 {
			fmt.Println("当sum>20时, 当前数是", i)
			break
		}
	}

	// 实现登录验证, 有三次机会 如果用户名为张无忌, 密码888提示登录成功, 否则提示还有几次机会
	var name string
	var passwd string
	loginChance := 3
	for i := 1; i <= 3; i++ {
		fmt.Printf("请输入用户: ")
		fmt.Scanln(&name)
		fmt.Printf("请输入密码: ")
		fmt.Scanln(&passwd)

		if name == "张无忌" && passwd == "888" {
			fmt.Printf("登录成功!\n")
			break
		} else {
			loginChance--
			fmt.Printf("密码错误, 当前还有%v次机会\n", loginChance)
		}
	}
	if loginChance == 0 {
		fmt.Printf("哦豁, 没有登录成功.\n")
	}
}
