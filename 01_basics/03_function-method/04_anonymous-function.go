package main

import "fmt"

// 定义匿名函数方便在同一个函数内部重复使用
func main() {
	f := func(a int, b int) int {
		return a + b
	}

	sum1 := f(1, 2)
	fmt.Println(sum1)
	sum2 := f(2, 3)
	fmt.Println(sum2)
}
