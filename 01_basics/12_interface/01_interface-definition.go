package main

import "fmt"

type Animal interface {
	Call()
	Run()
}

type Cat struct{} // Cat类型实现了Animal接口下的全部方法，则可以将Cat的实体赋值给Animal接口，这种动态思想称为鸭子类型

func (c Cat) Call() {
	fmt.Println("cat miao miao miao")
}
func (c Cat) Run() {
	fmt.Println("cat is running")
}

type Dog struct{} // Dog类型实现了Animal接口下的全部方法

func (c Dog) Call() {
	fmt.Println("dog wow wow wow")
}
func (c Dog) Run() {
	fmt.Println("dog is running")
}

func main() {
	var i Animal
	var c Cat
	var d Dog

	i = c
	i.Call()
	i.Run()

	i = d
	i.Call()
	i.Run()
}
