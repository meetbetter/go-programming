package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Id   int    `json:"user_id" gorm:"column:id"` // 在json处理时会将Id转成user_id
	Name string `json:"user_name" gorm:"column:name"`
	age  int    `json:"user_age" gorm:"column:age"` // 首字母大写才能对其他package可见，否则json.Marshal序列化时不会看见该成员
	sex  string
	addr string
}

func main() {
	var stu = Student{Id: 10086, Name: "张三", age: 28, sex: "男", addr: "西二旗"}

	jsBytes, err := json.Marshal(stu)
	if err != nil {
		fmt.Printf("json marshal err:%v\n", err)
		return
	}

	fmt.Printf("json:%s\n", string(jsBytes))
}
