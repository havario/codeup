package main

import (
	"fmt"
)

/*
	切片定义的基本语法
	var 切片名 []类型
	比如 var a []int
*/

func main() {
	// 演示切片的基本使用
	var intArr [5]int = [...]int{1, 22, 33, 66, 99}

	/*
		声明/定义一个切片
		1> slice 就是切片的名称
		2> intArr[1:3]表示slice引用到intArr这个数组
		3> 引用intArr数组的起始下标为1, 最后的下标为3(不包含3)
	*/
	slice := intArr[1:3]
	fmt.Printf("intArr: %v\n", intArr)
	fmt.Printf("slice元素是: %v\n", slice)
	fmt.Printf("slice元素个数: %v\n", len(slice))
	fmt.Printf("slice容量: %v\n", cap(slice)) // slice目前可以存储的最大元素的个数, 容量可动态变化

	fmt.Printf("intArr[1]的地址: %p\n", &intArr[1])
	fmt.Printf("slice[0]的地址: %p slice[0]的值: %v\n", &slice[0], slice[0])
	slice[1] = 34
	fmt.Println()
	fmt.Printf("intArr: %v\n", intArr)
	fmt.Printf("slice元素是: %v\n", slice)
}
