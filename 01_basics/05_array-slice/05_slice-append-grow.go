package main

import "fmt"

func main01() {

}

func main() {

	var sli1 []int //空切片
	fmt.Println("1:", len(sli1), cap(sli1))
	fmt.Printf("%p\n", sli1) //指向内存地址编号为0空间

	sli1 = append(sli1, 1)
	fmt.Println("2:", len(sli1), cap(sli1))

	sli1 = append(sli1, 2)
	fmt.Println("3:", len(sli1), cap(sli1))

	sli1 = append(sli1, 2, 3, 4, 5, 7)
	fmt.Println("4:", len(sli1), cap(sli1))

	var sli2 []int
	for i := 0; i <= 10000; i++ {
		//如果没有超过1024 按照上一次倍数扩容 超过1024  按照上一次的1/4扩容
		sli2 = append(sli2, i)
		fmt.Println(len(sli2), cap(sli2))
	}
}
