package main

import (
	"errors"
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

// 读取一个配置文件config.ini的信息
// 如果文件名传入不正确抛出自定义错误
func readConfig(fileName string) (err error) {
	if fileName == "config.ini" {
		return nil
	} else {
		// 返回自定义错误
		return errors.New("读取文件错误!")
	}
}

func test02() {
	err := readConfig("config.ini")
	if err != nil {
		// 如果读取文件错误发送错误, 就输出这个错误并终止程序
		panic(err)
	}

	fmt.Printf("test02后续代码继续执行\n")
}

func main() {

	// 测试go的错误
	// test()

	// 测试自定义错误的使用
	test02()
	fmt.Printf("main函数代码\n")
}
