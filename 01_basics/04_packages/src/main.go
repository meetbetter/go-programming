package main

import (
	"fmt"

	"go-programming/01_basics/04_packages/pkg2"

	"go-programming/01_basics/04_packages/pkg1"
)

/*
只有首字母大写的变量和函数才能在package外使用
init()在main函数前执行
多个package的init按import顺序和引用关系执行init
同一个package下的多个init按照文件名字典序执行
单个文件的init按顺序从上到下执行
**但是开发中不要依赖此顺序**
*/

var GlobalData int = 123

/*
go run src/main.go

pkg2 init is running
pkg1 file01 init is running, Data1:100
pkg1 file02 init is running, Data1:100
main GlobalData:123, pkg1 data:100
pkg1 DoSomething
pkg2 DoSomething
*/

func main() {
	fmt.Printf("main GlobalData:%d, pkg1 data:%d\n", GlobalData, pkg1.Data1)
	pkg1.DoSomething()
	pkg2.DoSomething()
}
