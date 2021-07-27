package main

import "fmt"

var a = 100

func main() {
	fmt.Printf("main %p\n", &a)
	demo1()
	fmt.Printf("demo1 修改后 %v\n", a)
	demo2(a)
	fmt.Printf("demo2 修改后 %v\n", a)
}

func demo1() {
	// 此处的a是全部变量
	fmt.Printf("demo1 %p\n", &a)
	a = 200
}

func demo2(a int) {
	// 此处的a是局部变量
	fmt.Printf("demo2 %p\n", &a)
	a = 300
}
