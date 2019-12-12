package api

import (
	"beego/app/Service/Dao"
	"beego/app/controllers"
)

type UserController struct {
	controllers.BaseController
}

func (this *UserController) Find()  {
	id, _ := this.GetInt("id")
	json := Dao.UserFind(id)

	this.Data["json"] = json
	this.ServeJSON()
}

func (this *UserController) List()  {
	//TODO
	json := make(map[string]string)
	this.Data["json"] = json
	this.ServeJSON()
}

func (this *UserController) Register() {
	mobile := this.GetString("mobile")
	name := this.GetString("name")

	json := Dao.UserInsert(name,mobile)

	this.Data["json"] = json
	this.ServeJSON()
}

func (this *UserController) Login()  {
	mobile := this.GetString("mobile")
	password := this.GetString("password")
	json := make(map[string]string)
	json["mobile"] = mobile
	json["password"] = password

	this.Data["json"] = map[string]interface{}{"code":200,"data":json}
	this.ServeJSON()
}
