package main

import "fmt"

func main() {
	var a = 100
	fmt.Printf("before a:%v\n", a)

	if a > 0 {
		a, err := demo() // 常见错误，此处的a是使用:定义的新的局部变量

		if err != nil {
			fmt.Printf("执行demo结果出错,:%v,%d\n", err, a)
			return
		}

		fmt.Printf("inner a:%p, %d\n", &a, a)
	}

	// 此处输出值为100
	fmt.Printf("after a:%p, %d\n", &a, a)
}

func demo() (int, error) {
	return 200, nil
}
