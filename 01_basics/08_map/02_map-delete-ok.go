package main

import "fmt"

func main() {
	m := map[int]string{100: "aaa", 200: "bbb", 250: "ccc"}
	fmt.Println(m)

	// 二值判断法，ok判断key是否存在
	value, ok := m[100]
	if ok {
		fmt.Printf("1, key exist, value:%s\n", value)
	} else {
		fmt.Printf("1, key not exist\n")
	}

	delete(m, 100) //delete(map，key)，删除map中的key-value

	value, ok = m[100]
	if ok {
		fmt.Printf("2, key exist, value:%s\n", value)
	} else {
		fmt.Printf("2, key not exist\n")
	}
}
