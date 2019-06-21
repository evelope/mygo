package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"mybeepro/models"
	// "path"
	// "time"
)

type MainController struct {
	beego.Controller
}

// get 方法
func (c *MainController) Get() {
	c.Data["name"] = "卫贵飞名字测试"

	//-------------------------------查
	//1 有rom 对象
	// o := orm.NewOrm()
	// //2 有要插入的结构体对象
	// user := models.User{}
	// //3 指定字段
	// user.Id = 2
	// //4 查询
	// err := o.Read(&user)
	// if err != nil {
	// 	beego.Info("查询失败", err)
	// 	return
	// }
	// c.Data["userList"] = user
	// beego.Info("查询成功", user)
	// **********************************



	//--------------------------------改
	//1 有rom 对象
	// o = orm.NewOrm()
	// //2 有要插入的结构体对象
	// user = models.User{}
	// //3 指定字段
	// user.Id = 1
	// //4 查询
	// err = o.Read(&user)
	// if err == nil {
	// 	user.Name = "www飞"
	// 	//5 更新
	// 	_,errr := o.Update(&user)
	// 	if errr != nil {
	// 		beego.Info("更新失败", err)
	// 		return
	// 	}
	// }
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
	if username=="" || pwd == "" {
		beego.Info("不能为空")
		c.Data["name"] = "不能为空"
		c.TplName = "index.html"
		return
	}

	o := orm.NewOrm()
	user := models.User{}
	user.Name = username
	user.Pwd = pwd
	_, err := o.Insert(&user)
	if err != nil {
		beego.Info("插入失败", err)
		c.TplName = "index.html"
		return
	}
	c.Data["name"] = "post 成功"
	// **********************************

	c.TplName = "index.html"
}

// delete
func (c *MainController) Deluser() {
	// --------------------------------删
		//1 有rom 对象
		o := orm.NewOrm()
		//2 有要插入的结构体对象
		user := models.User{}
		//3 指定字段
		user.Id = 5
		//4 查询
		_,err := o.Delete(&user)
		if err != nil {
			beego.Info("删除失败", err)
			return
		}
	// **********************************
	c.TplName = "index.html"

}



// ----------------- AddArticle
func (c *MainController) AddArticle () {
	// ------------------- 文件上传
	// // _,h,err := c.GetFile("file")
	// f,h,err := c.GetFile("file")
	// defer f.Close()

	// // 1 限定格式
	// fileext := path.Ext(h.Filename)
	// beego.Info(fileext)
	// if fileext != ".png" && fileext != ".jpg" {
	// 	beego.Info("文件格式错误")
	// 	return
	// }
	// // 2 限定 大小
	// if h.Size > 50000000 {
	// 	beego.Info("文件太大")
	// 	return
	// }

	// // 3 需要重命名
	// filename := time.Now().Format("2006-01-02-15-04-05") + fileext
	// if err != nil {
	// 	beego.Info("上传失败")
	// 	return
	// } else {
	// 	beego.Info("上传成功")
	// 	c.SaveToFile("file", "./static/img/"+filename)
	// }
	// ****************************************************
	

	ob := models.Article{}
	body := c.Ctx.Input.RequestBody
	json.Unmarshal(body, &ob);
	
	o := orm.NewOrm()
	_, err := o.Insert(&ob)
	if err != nil {
		beego.Info("新增失败", err)
		return
	}
	beego.Info("新增成功")
	c.TplName = "index.html"
}