package main

import "fmt"

func main() {
	var s1 string = "hello world"
	fmt.Printf("s1 len:%d\n", len(s1))
	for i, v := range s1 {
		fmt.Printf("i:%d, v:%c\n", i, v)
	}

	var s2 string = "hello 世界"
	fmt.Printf("s2 len:%d\n", len(s2))
	for i, v := range s2 {
		fmt.Printf("i:%d, v:%c\n", i, v)
	}

}
