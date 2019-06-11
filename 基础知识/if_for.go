package main

import (
	"fmt"
)

func main() {
	con := 10
	if con == 11 {
		fmt.Println(1)
	} else {
		fmt.Println(2)
	}

	var bb = 20
	// 值初始化
	if bb=11;bb== 11 {
		fmt.Println("bb")
	}

	var aa = "1"
	switch aa="2"; aa {
	case "2":
		fmt.Println("2")
		// break
		// 关键字  某一个执行 ，下面一个 无条件执行
		fallthrough
	case "3":
		fmt.Println("3")
		// break
		fallthrough
	default: 
		fmt.Println("hh")
	}


	switch {
	case aa == "2":
		fmt.Println("22222")
	default: 
		fmt.Println("hhh")
	}

	// if _,err:=func(){};err!=nil {
	// 	fmt.Print(3)
	// } else {
	// 	fmt.Print(4)
	// }
	fmt.Print("\n-----------\n")


	// for  map  slice ...
	for i:=0;i<10;i++ {
		fmt.Print(i)
	}

	//	死循环
	// for {
	// 	fmt.Print("o")
	// }

}
