package main

import (
	"factory/model"
	"fmt"
)


func main() {
	stu := model.NewStudent("tom",88.0)
	fmt.Println(*stu)
	fmt.Println(stu.GetName())
}