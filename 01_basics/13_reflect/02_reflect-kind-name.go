package main

/*

从relfect.Value中获取接口interface的信息
*/
import (
	"fmt"
	"reflect"
)

type student struct {
	name string
	age  int
}

type people struct {
	name string
	age  int
}

func main() {
	var num = 1.2345

	t := reflect.TypeOf(num)
	v := reflect.ValueOf(num)

	fmt.Printf("type.Kind:%v, type.Name:%v, value:%v\n", t.Kind(), t.Name(), v)

	s := student{
		name: "zhangsan",
		age:  18,
	}
	t2 := reflect.TypeOf(s)
	v2 := reflect.ValueOf(s)
	fmt.Printf("type.Kind:%v, type.Name:%v, value:%v\n", t2.Kind(), t2.Name(), v2)

	p := people{
		name: "tony",
		age:  30,
	}
	t3 := reflect.TypeOf(p)
	v3 := reflect.ValueOf(p)
	fmt.Printf("type.Kind:%v, type.Name:%v, value:%v\n", t3.Kind(), t3.Name(), v3)
}
