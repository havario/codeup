//  变量学习
package main

import "fmt"

// 定义全局变量
var number1 = 100
var number2 = 200
var checkstr1 = "Rockylinux"

// 全局变量的一次性声明
var (
	number3 = 300
	number4 = 400
	checkstr2 = "Alpine"
)

func main()  {
	// 演示golang如何一次性声明多个变量
	var num1, num2, num3 int
	fmt.Println("num1=", num1, "num2=", num2, "num3=", num3)

	// 一次性声明多个变量
	var num4, name, num5 = 611, "ysl", 1231
	fmt.Println("num4=",num4, "name=",name, "num5=",num5)

	// 一次性声明多个变量, 同样可以使用类型推导
	num6, nam, num7 := 3112, "Jack", 911
	fmt.Println("num6=",num6, "nam=",nam, "num7=",num7)

	// 输出全局变量
	fmt.Println("number1=",number1, "number2=",number2, "checkstr1=",checkstr1)
	fmt.Println("number3=",number3, "number4=",number4, "checkstr2=",checkstr2)
}