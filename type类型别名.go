package main

import (
	"fmt"
)

type begin int64

// 函数也是一种数据类型， 通过 type 给一个函数类型起名

type FuncType func(int,int,int) int

func main() {
	var aa begin = 10
	fmt.Println(aa)

	var TestFunc FuncType = fn
	bb := TestFunc(2,4,6);
	fmt.Println(bb, "bbb***")
}


func fn(a,b,c int) int {
	return a+b+c
}