package main

/*
递归函数
*/
import "fmt"

func main() {
	fmt.Println(factorial(10))
}

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	ret := factorial(num - 1)
	ret *= num //1*2*3*4*5
	return ret
}
