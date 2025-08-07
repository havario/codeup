package main

import (
	"fmt"
)

func test() {
	// 使用defer + recover 捕获处理异常
	defer func() {
		err := recover() // 内置函数, 可以捕获异常
		if err != nil {  // nil是零值, 如果非0说明捕获到错误
			fmt.Printf("err=%v\n", err)
			// 这里可以将错误信息发送给管理员
			fmt.Printf("发送邮件给sb@gmail.com\n")
		}
	}()

	num1 := 10
	num2 := 0
	result := num1 / num2
	fmt.Printf("result=%v\n", result) // panic: runtime error: integer divide by zero
}

func main() {

	// 测试go的错误
	test()
	fmt.Printf("main函数代码\n")
}
