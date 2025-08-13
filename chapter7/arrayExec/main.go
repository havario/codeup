package main

import (
	"fmt"
)

func main() {
	/*
		1. 创建一个byte类型的26个元素的数组, 分别放置'A'-'Z'
		使用for循环访问所有元素并打印出来 提示: 字符数据运算 'A' + 1 -> 'B'

		思路:
		1> 声明一个数组 var myChars [26]byte
		2> 使用for循环, 利用字符可以运算的特点来赋值 'A'+1 -> 'B'
		3> 使用for循环打印即可
	*/
	var myChars [26]byte
	for i := 0; i < len(myChars); i++ {
		myChars[i] = 'A' + byte(i) // 注意需要将 i => byte
	}

	for i := 0; i < len(myChars); i++ {
		fmt.Printf("%c ", myChars[i])
	}

	fmt.Println()
	fmt.Println()

	/*
		请求出一个数组的最大值, 并得到对应的下标

		思路:
		1> 声明一个数组 intArr := [...]int{1, -1, 9, 90, 11, 9000}
		2> "假定"第一个元素就是最大值, 下标为0
		3> 从第二个数开始循环对比, 发现有更大的数则交换
	*/
	intArr := [...]int{1, -1, 9, 90, 11, 9000}
	maxVal := intArr[0]
	maxValIndex := 0 // "假定"第一个元素就是最大值, 下标为0

	for i := 1; i < len(intArr); i++ { // 从第二个元素开始比较
		// 从第二个数开始循环对比, 发现有更大的数则交换
		if maxVal < intArr[i] {
			maxVal = intArr[i]
			maxValIndex = i
		}
	}
	fmt.Printf("maxVal=%v maxValIndex=%v\n", maxVal, maxValIndex)

	fmt.Println()

	/*
		请求出一个数组的"和"和平均值
		要求: 使用for range

		思路:
		1> 声明一个数组 intArr := [...]int{1, -1, 9, 90, 11}
		2> 求和
		> 求平均值
	*/
	intArr2 := [...]int{1, -1, 9, 90, 12}
	sum := 0                      // 求和的变量, 默认为0
	for _, val := range intArr2 { // 遍历intArr2数组
		// 累计求和
		sum += val
	}
	// 如何让平均值保留到小数点到小数?
	fmt.Printf("sum=%v 平均值=%v\n", sum, float64(sum)/float64(len(intArr2)))
}
