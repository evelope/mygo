package main

import (
	"fmt"
)

// len()  arr长度
var goarr [100]int
var goarr2 = [10]int{10, 50, 80}

func main() {
	fmt.Println(goarr, len(goarr))

	fmt.Println(goarr2, len(goarr2))

	fmt.Println(goarr2[2])

	for i := 0; i < len(goarr2); i++ {
		fmt.Println(goarr2[i])
	}
}
