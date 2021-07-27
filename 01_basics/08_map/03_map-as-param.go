package main

import "fmt"

func foo(m map[int]string) { //map是引用类型，函数内可以改变外部传进来的实参的值，因为m底层是通过指针指向存储数据的buckets
	delete(m, 101)
}

func main() {

	m := make(map[int]string)
	m[100] = "aaa"
	m[101] = "bbb"
	m[102] = "ccc"
	m[103] = "ddd"

	fmt.Printf("before m:%v, len:%d\n", m, len(m))
	foo(m)
	fmt.Printf("after m:%v, len:%d\n", m, len(m))
}
