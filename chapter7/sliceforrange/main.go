// Copyright 2025 honeok <i@honeok.com>

package main

import (
	"fmt"
)

func main() {

	// 使用常规for循环遍历切片
	var arr [5]int = [...]int{10, 20, 30, 40, 50}
	// slice := arr[1:4] // 20 30 40
	slice := arr[:] // 20 30 40 (0-结尾)
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%v]=%v \n", i, slice[i])
	}

	fmt.Println()

	// 使用for--range遍历切片
	for i, v := range slice {
		fmt.Printf("i=%v v=%v\n", i, v)
	}

	fmt.Println()

	/*
		切片可以继续切片
		slice已经是一个切片
	*/
	slice2 := slice[2:4] // 30 40
	slice2[0] = 100      // 同理slice[0]也会变化, 因为slice和slice2指向同一个数据空间, 因此slice2变化slice也会变化
	fmt.Println("slice2=", slice2)

	/*
		使用append内置函数, 可以对切片进行动态追加

		切片append操作的本质就是对数组扩容
		go底层会创建一下新的数组 newArr(安装后的扩容大小)
		将slice原来包含的元素拷贝到新的数组newArr
		slice重新引用到newArr
		注意newArr是在底层来维护的, 程序员不可见
	*/

	slice3 := []int{100, 200, 300}
	//通过append直接给slice3追加具体的元素
	slice3 = append(slice3, 400, 500, 600)
	// 通过append将切片slice3追加给slice3
	slice3 = append(slice3, slice3...)

	fmt.Println("slice3", slice3)

	// 通过append将切片slice3追加给slice3
	slice3 = append(slice3, slice3...)
	fmt.Println("slice3", slice3)

	fmt.Println()

	/*
		切片的拷贝操作

		1. copy(para1, para2) 参数的数据类型都是切片
		2. 按照下面的代码来看, slice4和slice5的数据空间是独立的, 相互不影响也就是说slice4[0]=999 slice5[0]仍然是1
	*/
	var slice4 []int = []int{1, 2, 3, 4, 5}
	var slice5 = make([]int, 10)
	fmt.Println(slice4)
	copy(slice5, slice4)
	fmt.Println(slice5)
}
