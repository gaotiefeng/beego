package api

import (
	"beego/app/Service/Biz/api"
	"beego/app/Service/Dao"
	"beego/app/constants"
	"beego/app/controllers"
	"github.com/astaxie/beego"
)

type UserController struct {
	controllers.BaseController
}


func (this *UserController) Login()  {
	//TODO
	mobile := this.GetString("mobile")
	password := this.GetString("password")
	json := make(map[string]string)
	json["mobile"] = mobile
	json["password"] = password

	this.Data["json"] = map[string]interface{}{"code":200,"data":json}
	this.ServeJSON()
}

func (this *UserController) Register() {
	mobile := this.GetString("mobile")
	name := this.GetString("name")

	json := api.UserBizInsert(name,mobile)

	this.Data["json"] = json
	this.ServeJSON()
}

func (this *UserController) Find() {
	id, _ := this.GetInt("id")
	var json interface{}
	if id == 0 {
		this.ResponseError(constants.SERVERERROR,"id 不能为空",id)
	}
	err,user := Dao.UserFind(id)

	if err != nil {
		beego.Info("查询失败",err)
		json = controllers.Error(constants.SERVERERROR,"查询失败",err)
	}else {
		json = controllers.Success(constants.SERVERERROR,"查询成功",&user)
	}
	this.Data["json"] = json
	this.ServeJSON()
}

func (this *UserController) List()  {
	//TODO
	json := make(map[string]string)
	this.Data["json"] = json
	this.ServeJSON()
}


func (this *UserController) Update()  {
	//TODO
	id,_ := this.GetInt("id")
	name := this.GetString("name")
	json := Dao.UserUpdate(id,name)

	this.Data["json"] = map[string]interface{}{"code":200,"data":json}
	this.ServeJSON()
}

func (this *UserController) Delete() {

	id,_ := this.GetInt("id")

	json := api.UserBizDelete(id)

	this.Data["json"] = json
	this.ServeJSON()
}
