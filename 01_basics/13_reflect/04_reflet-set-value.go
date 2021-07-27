package main

import (
	"fmt"
	"reflect"
)

func setValue(x interface{}) {
	v := reflect.ValueOf(x)

	switch v.Elem().Kind() { // 必须使用Elem()获取指针的kind
	case reflect.Int64:
		v.Elem().SetInt(28)
	default:
		fmt.Printf("unsupport value type:%T\n", x)
	}
}
func main() {
	var a int64 = 18
	setValue(&a) //必须用指针，否则reflect.Value.Elem处理非指针会panic
	fmt.Println("a: ", a)
	var b float32 = 18.0
	setValue(&b)
	fmt.Println("b: ", b)
}
