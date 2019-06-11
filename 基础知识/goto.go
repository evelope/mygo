package main

import (
	"fmt"
)

func main() {
	aa:=0
	End_my:
	aa++
	fmt.Println(aa)
	
	fmt.Printf("请回车 :")
	fmt.Scanf("%s", &aa)
	goto End_my // 直接去定义的地方
	
	fmt.Println(22)
	
	
	fmt.Println(33)
}