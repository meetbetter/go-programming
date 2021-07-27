package main

import "fmt"

func main() {
	var a int64 = 100
	var b string = "hello 世界"
	var c float64 = 1.007

	foo(a)
	foo(b)
	foo(c)
}

func foo(i interface{}) { // 空接口可以接受任何类型
	fmt.Printf("get: %v\n", i)
}
