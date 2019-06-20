package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id int
	Name string
	Pwd string
}
func init(){
	// 设置数据库基本信息
	orm.RegisterDataBase("default","mysql","root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")

	// 映射model数据    new(User) 表
	orm.RegisterModel(new(User))

	// 生成表 1  default 默认数据库  2 是否强制更新 重新生成表（表建错了）  3 是否可见
	orm.RunSyncdb("default", false, true)
}