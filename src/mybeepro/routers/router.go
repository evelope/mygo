package routers

import (
	"mybeepro/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/posttest", &controllers.MainController{})
	beego.Router("/deluser", &controllers.MainController{}, "delete:Deluser")

	// Article 列表
	// --------------增
	beego.Router("/addArticle", &controllers.MainController{}, "post:AddArticle")


	// --------------查
	beego.Router("/getArticle", &controllers.MainController{}, "get:GetArticle")
}
