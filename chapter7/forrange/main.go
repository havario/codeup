package main

import (
	"fmt"
)

func main() {

	// 演示for-range遍历数组
	heroes := [...]string{"Lisa", "Jisoo", "Jennie", "Rosé"}
	for hero, value := range heroes {
		fmt.Printf("hero=%v value=%v\n", hero, value)
	}

	for _, value := range heroes {
		fmt.Printf("value=%v\n", value)
	}
}
