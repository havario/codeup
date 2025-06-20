// Copyright (c) 2025 honeok <honeok@disroot.org>

package main

import (
	"fmt"
)

// 写一个简单的函数
func test(char byte) byte {
	return char + 1
}

func main() {
	// 编写一个程序, 该程序可以接受一个字符,  比如 a b c d c f g
	// a 表示星期一, b 表示星期二 ... 根据用户的输入显示相依的信息

	// 要求: 使用switch语句完成

	var input byte
	fmt.Printf("请输入一个字符: ")
	fmt.Scanf("%c", &input)

	switch test(input) {
	case 'a':
		fmt.Printf("周一买卖稀\n")
	case 'b':
		fmt.Printf("周二买卖淡\n")
	case 'c':
		fmt.Printf("周三买卖惨\n")
	case 'd':
		fmt.Printf("周四买卖次\n")
	case 'e':
		fmt.Printf("周五买卖苦\n")
	case 'f':
		fmt.Printf("周六吃块肉\n")
	case 'g':
		fmt.Printf("周天过大年\n")
	// ...
	default:
		fmt.Printf("输正确啊傻逼\n")
	}

	// switch 的数据类型必须和case一致
	n1 := 5
	n2 := 20
	switch n1 {
	case n2, 20, 5: // case后面可以有多个表达式
		fmt.Println("ok1")
	default:
		fmt.Println("无匹配项")
	}

	// switch后也可以不带表达式, 类似if --else 分支使用
	var age int = 10
	switch {
	case age == 10:
		fmt.Printf("age是10\n")
	case age == 20:
		fmt.Printf("age是20\n")
	default:
		fmt.Printf("没有匹配项\n")
	}

	// case中也可以对范围进行判断
	var score int = 30
	switch {
	case score > 90:
		fmt.Printf("成绩优秀\n")
	case score >= 70 && score <= 90:
		fmt.Printf("成绩良\n")
	default:
		fmt.Printf("真你妈垃圾\n")
	}

	// switch后也可以直接声明/定义一个变量, 分号结束 不推荐
	switch grade := 90; {
	case grade > 90:
		fmt.Printf("成绩优秀\n")
	case grade >= 70 && grade <= 90:
		fmt.Printf("成绩良\n")
	default:
		fmt.Printf("真你妈垃圾\n")
	}

	// switch的穿透 fallthrought
	var num int = 10
	switch num {
	case 10:
		fmt.Printf("ok1\n")
		fallthrough // 默认只能穿透一层
	case 20:
		fmt.Printf("ok2\n")
		fallthrough
	case 30:
		fmt.Printf("ok3\n")
	default:
		fmt.Printf("错误匹配\n")
	}
}
