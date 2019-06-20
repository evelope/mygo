package main

import (
	_ "mybeepro/routers"
	_ "mybeepro/models"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/down1", "download1")
	beego.Run()
};

