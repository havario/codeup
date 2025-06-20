package main

import (
	"fmt"
)

func main() {
	// 使用switch 把小写类型的char型转换为大写(接受键盘输入)
	// 只转换a, b, c, d, e 其他输出 'other'

	var char byte
	fmt.Printf("请输入一个字符: ")
	fmt.Scanf("%c", &char)

	switch char {
	case 'a':
		fmt.Printf("A\n")
	case 'b':
		fmt.Printf("B\n")
	case 'c':
		fmt.Printf("C\n")
	case 'd':
		fmt.Printf("D\n")
	case 'e':
		fmt.Printf("E\n")
	default:
		fmt.Printf("Other\n")
	}

	// 学生成绩大于60分的, 输出合格, 低于60分的 输出不合格 (输入的成绩不能大于100)
	var score float64
	fmt.Printf("请输入学生成绩: ")
	fmt.Scanln(&score)

	switch int(score / 60) {
	case 1:
		fmt.Printf("成绩合格\n")
	case 0:
		fmt.Printf("成绩不合格\n")
	default:
		fmt.Printf("无效成绩\n")
	}

	// 根据用户指定的月份
	// 打印该月份的季节 3 4 5春季 6 7 8夏季 9 10 11 秋季 12 1 2 春季
	var month byte
	fmt.Printf("请输入查询月份: ")
	fmt.Scanln(&month)

	switch month {
	case 3, 4, 5:
		fmt.Printf("spring\n")
	case 6, 7, 8:
		fmt.Printf("summer\n")
	case 9, 10, 11:
		fmt.Printf("autumn\n")
	case 12, 1, 2:
		fmt.Printf("spring\n")
	default:
		fmt.Printf("Wrong value\n")
	}
}
