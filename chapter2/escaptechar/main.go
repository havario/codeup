package main

import "fmt" // fmt包主要提供格式化，输出、输入的函数

func main()  {
	// 转义字符的使用
	fmt.Println("tom\tjack")

	fmt.Println("hello\nworld")

	fmt.Println("D:\\go_code\\codeup")

	fmt.Println("ranbot say \"i love you\"")

	// \r 回车 从当前行开始最前面开始输出，覆盖之前的内容
	fmt.Println("天龙八部雪山飞狐\r张飞")
}