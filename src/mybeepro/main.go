package main

import (
	_ "mybeepro/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/down1", "download1")
	beego.Run()
};

