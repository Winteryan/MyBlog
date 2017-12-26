package main

import (
	"hello/models"
	_ "hello/routers"

	"github.com/astaxie/beego"
)

func main() {
	models.Init()
	beego.Run()
}

//func init() {
//	orm.RegisterDriver("sqlite", orm.DRSqlite)
//	orm.RegisterDataBase("default", "sqlite3", "database/orm_test.db")
//	orm.RegisterModel(new(models.Article))
//}
