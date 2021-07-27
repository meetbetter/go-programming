package main

import "fmt"

func main() {
	//自动推导类型  根据变量后面的值进行类型推导
	f := 3.14 //float64
	
	fmt.Printf("type:%T, value:%f\n", f, f)

	var a bool //默认值为false
	fmt.Printf("%T, %t\n", a, a)

	//多重赋值
	x, y, z := 10, 3.14, "hello"
	fmt.Printf("type:%T, value:%d\n", x, x)
	fmt.Printf("type:%T, value:%f\n", y, y)
	fmt.Printf("type:%T, value:%s\n", z, z)

	// _ 表示匿名变量
	b, _ := 10, 20
	fmt.Println(b)
}
