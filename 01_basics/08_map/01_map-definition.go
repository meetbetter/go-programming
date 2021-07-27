package main

import "fmt"

func main() {
	var m1 map[int]string = make(map[int]string)
	m1[0] = "张三"
	m1[1] = "李四"
	m1[2] = "王五"
	fmt.Println("m1:", m1)

	m2 := map[string]string{"k1": "zhangsan", "k2": "lisi"} //map中的key必须支持== != 操作  不能使用函数 切片 map
	fmt.Println("m2:", m2)
}
