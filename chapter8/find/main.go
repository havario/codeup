package main

import "fmt"

func main() {
	/*
		有一个数列: 白眉鹰王 金毛狮王 紫衫龙王 青翼蝠王
		猜数游戏: 从键盘中任意输入一个名称, 判断数列中是否包含此名称 (顺序查找)

		1. 定义一个字符串数组
		2. 控制台接受名称依次比对, 如果有就提示
	*/
	var heroName = ""

	names := [4]string{"白眉鹰王", "金毛狮王", "紫衫龙王", "青翼蝠王"}
	fmt.Printf("请输入查找的人名:")
	fmt.Scanln(&heroName)

	// 顺序查找:  第一种方式
	for i := 0; i < len(names); i++ {
		if heroName == names[i] {
			fmt.Printf("名字已找到: %v 下标: %v\n", heroName, i)
			break
		} else if i == (len(names) - 1) {
			fmt.Printf("名字没有找到: %v\n", heroName)
		}
	}
}
