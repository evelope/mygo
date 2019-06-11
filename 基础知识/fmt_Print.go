package main
import (
	"fmt"
)
// aa := "asas"  **

var (
	aa int= 10
)
var dd, ee, ff = 10,"qqqq", true;
var gg = 11;
func main() {
	// 自动推导
	aa := "asas"
	bb,cc := 10, 20

	bb,cc = cc,bb

	fmt.Println(aa,bb,cc,dd,ee,ff)

	fmt.Printf("aa = %s,bb = %d,cc = %d,dd = %d,ee = %s,ff = %t\n",aa,bb,cc,dd,ee,ff)
	//%T  变量所属类型
	//%d 值
	fmt.Printf("%T\n",aa)
	fmt.Println(&gg)
}