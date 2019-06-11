package main

import (
	"fmt"
)
var zz = 10
var zzz = 100

var xx *int   // 声明指针变量
func main() {
	// & 内存地址
	fmt.Println("zz", zz, &zz)
	fmt.Println("zzz", zzz, &zzz)

	xx = &zz
	fmt.Println("*xx", xx, *xx, *&zzz)
}