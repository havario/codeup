package main

import (
	"fmt"
)

func main() {
	/* 二维数组
	0 0 0 0 0 0
	0 0 1 0 0 0
	0 2 0 3 0 0
	0 0 0 0 0 0
	*/

	/*声明4行×6列的二维数组
	arr[0] → [6]int
	arr[1] → [6]int
	arr[2] → [6]int
	arr[3] → [6]int
	*/
	var arr [4][6]int
	// 赋初值
	arr[1][2] = 1
	arr[2][1] = 2
	arr[2][3] = 3

	// 遍历二维数组
	for i := 0; i < 4; i++ { // 遍历行
		for j := 0; j < 6; j++ { // 遍历列
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}

	fmt.Println()

	// 直接赋值
	var arr3 [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr3)

	fmt.Println()

	// 直接赋值
	var arr4 [2][3]int = [...][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr4)

	fmt.Println()

	arr5 := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr5)
}
