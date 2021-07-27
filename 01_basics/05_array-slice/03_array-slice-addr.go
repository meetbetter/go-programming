package main

import "fmt"

func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Printf("arr %p, %T\n", &arr, arr) //数组名的地址 就是数组首元素地址
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%p\n", &arr[i]) //数组中元素的内存地址
	}

	fmt.Println("------------------")

	sli := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("sli %p, %T \n", &sli, sli) //切片名的地址 不是数组首元素地址，而是slice结构体的地址
	for i := 0; i < len(sli); i++ {
		fmt.Printf("%p\n", &sli[i]) //切片中元素的内存地址
	}
}
