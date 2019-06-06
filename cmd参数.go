package main

import (
	"fmt"
	"os"
)


// defer 延迟 执行，  / 多个 倒着执行
 
func main() {
	// 接受用户 参数
	list := os.Args
	n:= len(list)
	fmt.Println(n,"**")
	for i,x:=range list {
		fmt.Println(i,x,"---")
	}
}
