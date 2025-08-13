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
	fmt.Println("intArr: ", intArr)
	fmt.Println("slice元素是: ", slice)
	fmt.Println("slice元素个数: ", len(slice))
	fmt.Println("slice容量: ", cap(slice)) // slice目前可以存储的最大元素的个数, 容量可动态变化
}
