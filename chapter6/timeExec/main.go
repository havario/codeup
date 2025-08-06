package main

import (
	"fmt"
	"strconv"
	"time"
)

func test03() {

	str := ""
	for i := 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
}

func main() {
	// 在执行前获取当前时间戳
	startTime := time.Now().Unix()

	test03()

	endTime := time.Now().Unix()
	fmt.Printf("执行test03函数耗费时间为: %v秒\n", endTime-startTime)
}
