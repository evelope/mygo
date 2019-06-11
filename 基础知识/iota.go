package main
import (
	"fmt"
)

// iota 枚举
func main() {

	// 必须是  const
	const (
		aa = iota
		bb,cc,dd = iota,iota,iota
		ee = iota
	)

	fmt.Println(aa,bb,cc,dd,ee)



	var t complex128 = 2.1+3.14i
	fmt.Println("real(t) = ", real(t), ", imag(t) = ", imag(t))
}