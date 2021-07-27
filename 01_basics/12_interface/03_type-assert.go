package main

import "fmt"

func main() {
	var i interface{}
	var a int64 = 100
	var b string = "hello 世界"
	var c float64 = 1.007

	i = a
	v := i.(int64) // 类型断言
	fmt.Printf("value:%v, type:%T\n", v, v)

	i = b
	v2, ok := i.(string) // 类型断言时，使用二值接收防止断言失败panic
	if ok {
		fmt.Printf("value:%v, type:%T\n", v2, v2)
	} else {
		fmt.Println("断言失败")
	}

	i = c
	switch i.(type) { // 使用switch判断接口中数据的类型
	case int64:
		fmt.Printf("int64, value:%v\n", i)
	case string:
		fmt.Printf("string, value:%v\n", i)
	case float64:
		fmt.Printf("float64, value:%v\n", i)
	default:
		fmt.Printf("type unknown")
	}
}
