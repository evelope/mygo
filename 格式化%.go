package main

import (
	"fmt"
)

func main() {
	a := 10
	b := "asb"
	c := 'a'
	d := 3.14
	var e bool
	// %T 类型
	a++
	a++
	fmt.Printf("%T,%T%T,%T\n", a, b, c, d)
	// %d,%s,%c,%f  整数 字符串 字符 浮点
	fmt.Printf("%d,%s,%c,%f,%t\n", a, b, c, d, e)
	// %v  自动匹配
	fmt.Printf("%v--%v--%v--%v--%v", a, b, c, d, e)
}
