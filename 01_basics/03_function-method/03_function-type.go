package main

import "fmt"

func foo1() {
	fmt.Println("foo1 entry")
}

func foo2(a int, b int) int {
	fmt.Println("foo2 entry")
	return a + b
}

func main() {
	fmt.Printf("%T\n", foo1)
	fmt.Printf("%T\n", foo2)

	var f func(int, int) int //定义函数类型的变量

	f = foo2 //给变量f赋值
	sum := f(100, 200)
	fmt.Println(sum)

}
