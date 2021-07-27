package main

import "fmt"

/*
func 函数名（函数参数列表）函数返回值类型｛
	函数体
	return 返回值
｝
*/

func main() {
	a, b := 20, 10
	ret := sum(a, b)
	fmt.Println(ret)

	ret = sum2(a, b)
	fmt.Println(ret)
}

func sum(a int, b int) (sum int) {
	sum = a + b
	return
}

func sum2(a int, b int) int {
	return sum(a, b)
}
