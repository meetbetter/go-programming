package main

import "fmt"

func main() {
	A := make([]int, 5, 10)
	fmt.Printf("1： 值：%v, 地址：%p，len：%d, cap:%d\n", A, &A, len(A), cap(A))
	A = A[2:3] // 左闭右开
	fmt.Printf("2： 值：%v, 地址：%p，len：%d, cap:%d\n", A, &A, len(A), cap(A))
	A = A[2:3:5]
	fmt.Printf("3： 值：%v, 地址：%p，len：%d, cap:%d\n", A, &A, len(A), cap(A))
}
