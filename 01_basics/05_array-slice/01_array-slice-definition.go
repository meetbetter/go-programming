package main

import "fmt"

func main() {
	// 数组是一块连续的内存，存储固定长度的同类型数据
	var arr1 [10]int = [10]int{1, 2, 3, 4, 5}
	fmt.Printf("%T\n", arr1) // [10]int
	fmt.Println(arr1)

	arr2 := [...]int{1, 2, 3, 4, 5} //...数组长度根据成员个数确定
	fmt.Printf("%T\n", arr2)        //[5]int
	fmt.Println(arr2)

	// slice的定义
	sli := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("%T \n", sli)
	fmt.Println(sli)
}
