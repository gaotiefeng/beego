package api

import (
	"beego/app/controllers"
	"github.com/astaxie/beego/orm"
)

type UserController struct {
	controllers.BaseController
}

func (this *UserController) Index()  {
	//TODO 验证
	mobile := this.GetString("mobile")
	o := orm.NewOrm()
	var maps []orm.Params
	o.Raw("SELECT * FROM `user` WHERE mobile = ? LIMIT 1;", mobile).Values(&maps)
	json := map[string]interface{}{"code": 0, "data": maps[0]}
	this.Data["json"] = &json
	this.ServeJSON()
}

func (this *UserController) Register() {
	id := this.GetString("id")
	this.Data["json"] = id
	this.ServeJSON()
}

func (this *UserController) Login()  {
	mobile := this.GetString("mobile")
	password := this.GetString("password")
	json := map[string]interface{}{"mobile":mobile,"password":password}

	this.Data["json"] = map[string]interface{}{"code":200,"data":json}
	this.ServeJSON()
}
