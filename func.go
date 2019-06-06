package main

import (
	"fmt"
	// "my_pkg"
)

func main() {
	text_func(1, true, 3)
	text_func_arg(1,2,3,5)
	r,b:=text_func_return1()
	fmt.Println(r,b)
}

func text_func(a int, b bool, c int) (int, bool) {
	fmt.Print(a, b, c)
	return a, b
}

func text_func_arg(args ...int){
	fmt.Print(args)
}

// go 推荐 return
func text_func_return1()(result int,b string){
	result,b = 11, "bbbbbbbbb"
	return
}
