package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"mybeepro/models"
)

type MainController struct {
	beego.Controller
}

// get 方法
func (c *MainController) Get() {
	c.Data["name"] = "卫贵飞名字测试"

	//-------------------------------查
	//1 有rom 对象
	o := orm.NewOrm()
	//2 有要插入的结构体对象
	user := models.User{}
	//3 指定字段
	user.Id = 2
	//4 查询
	err := o.Read(&user)
	if err != nil {
		beego.Info("查询失败", err)
		return
	}
	c.Data["userList"] = user
	beego.Info("查询成功", user)
	// **********************************




	//--------------------------------改
	//1 有rom 对象
	o = orm.NewOrm()
	//2 有要插入的结构体对象
	user = models.User{}
	//3 指定字段
	user.Id = 1
	//4 查询
	err = o.Read(&user)
	if err == nil {
		user.Name = "www飞"
		//5 更新
		_,errr := o.Update(&user)
		if errr != nil {
			beego.Info("更新失败", err)
			return
		}
	}
	// **********************************







	c.TplName = "index.html"
}

// post 重写 插入User
func (c *MainController) Post() {
	c.Data["name"] = "post新增"

	// --------------------------------增
	// //1 有rom 对象
	// o := orm.NewOrm()
	// //2 有要插入的结构体对象
	// user := models.User{}
	// //3 对象结构体赋值
	// user.Name = "卫贵飞"
	// user.Pwd = "123456"
	// //4 插入  返回值 1 _ 插入几行 2 错 error
	// _, err := o.Insert(&user)
	// if err != nil {
	// 	beego.Info("插入失败", err)
	// 	return
	// }


	// 根据form表单 -*-*-*-*-*-*-******-*-*-*
	username := c.GetString("username")
	pwd := c.GetString("pwd")

	o := orm.NewOrm()
	user := models.User{}
	user.Name = username
	user.Pwd = pwd
	if username=="" || pwd == "" {
		beego.Info("不能为空")
		c.Data["name"] = "不能为空"
		c.TplName = "index.html"
		return
	}
	_, err := o.Insert(&user)
	if err != nil {
		beego.Info("插入失败", err)
		c.TplName = "index.html"
		return
	}
	c.Data["name"] = "post 成功"
	// **********************************


	// --------------------------------删
		//1 有rom 对象
		// o := orm.NewOrm()
		// //2 有要插入的结构体对象
		// user := models.User{}
		// //3 指定字段
		// user.Id = 1
		// //4 查询
		// _,err := o.Delete(&user)
		// if err != nil {
		// 	beego.Info("删除失败", err)
		// 	return
		// }
	// **********************************

	c.TplName = "index.html"
}
