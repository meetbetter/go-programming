package main

import "fmt"

/*
func (方法接收者)方法名（参数列表）返回值类型｛

	return 返回值
｝
*/
type MyInt int64 // 方法接收者必须是自定义类型

func (m MyInt) sum(a, b int64) int64 {
	sum := a + b
	return sum
}

type Student struct {
	Name string
	Age  int
}

func (s *Student) Change1(name string, age int) { // 方法接收者是指针，则可以在方法内修改接收者指向的内容
	s.Name = name
	s.Age = age
	return
}

func (s *Student) Change2(name string, age int) {
	s = nil //指针只能修改s指向的内容，不能修改s变量本身
	return
}

func main() {
	var m MyInt
	ret := m.sum(10, 20)
	fmt.Println(ret)

	s := &Student{Name: "zhangsan", Age: 18}
	s.Change1("lisi", 20)
	fmt.Printf("after Change1:%v\n", *s)
	s.Change2("lisi", 20)
	fmt.Printf("after Change2:%v\n", *s)
}
