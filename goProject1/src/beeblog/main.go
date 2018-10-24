package main

import (
	"beeblog/models"
	_ "beeblog/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// modol初始化函数
func init() {
	models.RegisterDB()
}

func main() {
	// 设置debug
	orm.Debug = true
	// 自动建表, name数据库名称， force是运行是总是会删掉重建，verbose是否打印相关信息
	orm.RunSyncdb("default", false, true)

	beego.Run()
}
