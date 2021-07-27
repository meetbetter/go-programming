package main

import "fmt"

func main() {
	var m1 map[int]string = make(map[int]string)
	m1[0] = "张三"
	m1[1] = "李四"
	m1[2] = "王五"
	fmt.Println("origin m1:", m1)

	m2 := m1 // 浅拷贝，共用底层buckets，修改m1会影响m2
	fmt.Printf("before m1:%v, m2:%v\n", m1, m2)
	delete(m1, 1)
	fmt.Printf("after m1:%v, m2:%v\n", m1, m2)

	// 遍历实现深拷贝，修改m1不影响m3
	m3 := make(map[int]string)
	for k, v := range m1 {
		m3[k] = v
	}
	fmt.Printf("before m1:%v, m3:%v\n", m1, m3)
	delete(m1, 0)
	fmt.Printf("after m1:%v, m3:%v\n", m1, m3)

}
