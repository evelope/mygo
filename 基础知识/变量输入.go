package main

import (
	"fmt"
)

func main() {
	var a string
	fmt.Printf("请问墨问是 ？ ： ")
	// 阻塞等待用户输入
	fmt.Scanf("%s", &a)

	fmt.Println("墨问是 ：沙雕！")
	fmt.Println("呸")
	fmt.Println("墨问是 ：", a)
}