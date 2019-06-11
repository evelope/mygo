package main

// slice  切片

// len 长度  append 添加
import (
	"fmt"
)

var Go [10]int
var Golang = []int{10,20,30}
var Golang2 = make([]int, 10)

func main() {
	fmt.Println(Go, Golang,Golang2)

	fmt.Println(len(Go))

	Golang = append(Golang, 4)

	fmt.Println(Golang)
}