/*
Copyright (c) 2025 honeok <i@honeok.com>
SPDX-License-Identifier: MIT
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println()
	/*
		定义二维数组, 用于保存三个班, 每个班五名同学成绩
		并求出每个班级平均分, 以及所有班级平均分
	*/

	// 定义二维数组 3行5列
	var scores [3][5]float64

	// 循环输入成绩
	for i := 0; i < len(scores); i++ {
		for k := 0; k < len(scores[i]); k++ {
			fmt.Printf("请输入第%d班的第%d个学生的成绩: ", i+1, k+1)
			fmt.Scanln(&scores[i][k])
		}
	}

	// 遍历输入成绩后的二维数组, 统计平均分

	totalSum := 0.0 // 变量用于累计所有班级的总分
	for i := 0; i < len(scores); i++ {
		sum := 0.0 // 存储班级总分
		for k := 0; k < len(scores[i]); k++ {
			sum += scores[i][k]
		}
		totalSum += sum
		fmt.Printf("第%d班的总分为%v, 平均分%v\n", i+1, sum, sum/float64(len(scores[i])))
	}
	fmt.Printf("所有班级总分为%v, 所有班级平均分%v\n", totalSum, totalSum/float64(15))
}
