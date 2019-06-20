package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	Id int
	// unique 唯一
	Name string `orm:"unique"`
	Pwd string
}

// digits(4);decimal(2)  12.24   总位数4  小数位 2
type Article struct {
	// pk 主键 ， auto 自增(int int32 int64 unit31 unit64)
	Id int `orm:"pk;auto"`
	// size 长度
	Aname string `orm:"size(20)"`
	// auto_now  每次提交都更新最新时间
	// auto_now_add  第一次保存 才 设置时间
	Atime time.Time `orm:"auto_now"`
	// default 默认值   null 允许为空
	Acount int `orm:"default(0);null"`
	Acontent string
	Aimg string
}
func init(){
	// 设置数据库基本信息
	orm.RegisterDataBase("default","mysql","root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")

	// 映射model数据    new(User) 表
	orm.RegisterModel(new(User),new(Article))

	// 生成表 1  default 默认数据库  2 是否强制更新 重新生成表（表建错了）  3 是否可见
	orm.RunSyncdb("default", false, true)
}