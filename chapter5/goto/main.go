package main

import (
	"fmt"
)

func main() {

	// goto的使用

	num := 20
	fmt.Println("ok1")
	if num >= 10 {
		goto lable1
	}
	fmt.Println("ok2")
	fmt.Println("ok3")
	fmt.Println("ok4")
	fmt.Println("ok5")
lable1:
	fmt.Println("ok6")
	fmt.Println("ok7")
	fmt.Println("ok8")
	fmt.Println("ok9")
	fmt.Println("ok10")
}
