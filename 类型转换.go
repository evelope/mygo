package main

import (
	"fmt"
)

func main() {
	// var flag bool
	// bool 不能转换
	// fmt.Println(int(falg))

	var ch byte
	ch = 'a'
	var t int

	t = int(ch) // 类型转换
	fmt.Println(t, int(ch))

}
