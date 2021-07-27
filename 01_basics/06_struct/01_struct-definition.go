package main

import "fmt"

/*
type 结构体名称 struct{
	结构体成员
}
*/
type student struct {
	//结构体成员
	id   int
	name string
	age  int
	sex  string
	addr string
}

func main() {
	var stu student
	stu.id = 10010
	stu.age = 18
	stu.name = "张三"
	stu.addr = "回龙观"
	stu.sex = "男"
	fmt.Println(stu)

	stu2 := student{id: 10086, name: "李四", age: 28, sex: "男", addr: "西三旗"}
	stu2.name = "王五"
	fmt.Println(stu2)
}
