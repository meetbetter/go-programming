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

	fmt.Printf("type %v, value:%v\n", t, v)

}
