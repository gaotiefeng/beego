package main

import (
	"beego/app/models"
	_ "beego/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init()  {
	dataSrouce := "liunian:" + "Gtf2019." + "@tcp(127.0.0.1:3306)/beego?charset=utf8"

	orm.RegisterDataBase("default","mysql",dataSrouce,60)

	orm.RegisterModel(new(models.User))

	beego.BConfig.Listen.HTTPPort = 8030
	beego.BConfig.WebConfig.EnableDocs = true
}

func main(){

	beego.Run()
}
