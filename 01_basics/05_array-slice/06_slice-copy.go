package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4, 5}

	s2 := s1 //浅拷贝，修改s1仍会影响s2，因为底层是共用同一个数组
	fmt.Printf("1, s1:%v, s2:%v\n", s1, s2)
	s2[0] = 100
	fmt.Printf("2, s1:%v, s2:%v\n", s1, s2)

	s3 := make([]int, 10) // 长度要足以接收源数组，否则只能拷贝部分
	copy(s3, s1)          //copy(目标，源)，深拷贝，拷贝的是内容，源切片和目标切片不是同一个底层数组
	fmt.Printf("3, s1:%v, s3:%v\n", s1, s3)
	s1[0] = 200
	fmt.Printf("4, s1:%v, s3:%v\n", s1, s3)
}
