package main

import (
	"fmt"
)


// var kvs =map[int]string{1: "aa",2: "bb"}

// var kvs map[int]string=map[int]string{1: "aa",2: "bb"}

var kvs map[int]string
// make 初始化 
func main() {
	kvs = make(map[int]string)
	kvs[10] = "asasasasas"
	fmt.Println(kvs)

	delete(kvs,10)
	fmt.Println(kvs)
}