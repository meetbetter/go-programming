package main

import "fmt"

func bubbleSort(slice []int) { //切片作为函数参数，内部成员修改可以影响外部，因为slice结构体中包含一个指向底层数组的指针
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-1-i; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}
func main() {
	slice := []int{9, 1, 5, 6, 10, 8, 3, 7, 2, 4}
	bubbleSort(slice)
	fmt.Println(slice)
}
