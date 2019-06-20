package routers

import (
	"mybeepro/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/posttest", &controllers.MainController{})
}
