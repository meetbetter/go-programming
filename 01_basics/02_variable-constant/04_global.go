package main

import "fmt"

//全局变量  作用域  项目中所有文件 如果在同一包下 直接调用 如果在不用包下 需要使用包名.全局变量名 调用
var Value int = 123

func main() {
	foo()
}

func foo() {
	fmt.Println(Value)
}
