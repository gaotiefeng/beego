package main

import (
	"beego/app/models"
	_ "beego/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {

	host := beego.AppConfig.String("mysqlhost")
	prot := beego.AppConfig.String("mysqlprot")
	user := beego.AppConfig.String("mysqluser")
	pass := beego.AppConfig.String("mysqlpass")
	db := beego.AppConfig.String("mysqldb")

	logs.SetLogger(logs.AdapterConsole, `{"level":1,"color":true}`)
	logs.Debug(host)
	dataSrouce := user + ":" + pass + "@tcp(" + host + ":" + prot + ")/" + db + "?charset=utf8"

	orm.RegisterDataBase("default", "mysql", dataSrouce, 60)

	orm.RegisterModel(new(models.User))

	// beego.BConfig.Listen.HTTPPort = 8030
	beego.BConfig.WebConfig.EnableDocs = true
}

func main() {

	orm.Debug = true
	beego.Run()
}
