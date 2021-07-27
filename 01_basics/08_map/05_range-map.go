package main

import "fmt"

func main() {
	var m1 map[int]string = make(map[int]string)
	m1[0] = "张三"
	m1[1] = "李四"
	m1[2] = "王五"

	for k, v := range m1 { // map的遍历是无序的
		fmt.Println(k, v)
	}
}
