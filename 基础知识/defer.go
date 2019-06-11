package main

import (
	"fmt"
)


// defer 延迟 执行，  / 多个 倒着执行
 
func main() {
	fmt.Println(1);
	defer fmt.Println(2);
	defer fmt.Println(3);
	defer func () {  // 切换defer注释查看效果
		x:=0
		a:=10/x
		fmt.Println(4,a);
	}()
	defer fmt.Println(5);
}
