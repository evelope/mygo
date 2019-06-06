package main

import (
	"fmt"
)
type my_struct struct {
	ID int
	Url string
}
func main() {
	// 结构体赋值
	// my_struct{ID: 100,Url: "22222"}
	var my my_struct
	my.ID = 101
	my.Url = "33333"

	fmt.Println(my, my.ID)
}