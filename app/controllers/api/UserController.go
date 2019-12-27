package api

import (
	"beego/app/Service/Biz/api"
	"beego/app/Service/Dao"
	"beego/app/Service/Formatter"
	"beego/app/constants"
	"beego/app/controllers"
	"github.com/astaxie/beego"
)

type UserController struct {
	controllers.BaseController
}


func (this *UserController) Login()  {
	//TODO request verify
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
	password := this.GetString("password","123456")
	json := api.UserBizInsert(name,mobile,password)

	this.Data["json"] = json
	this.ServeJSON()
}

func (this *UserController) Find() {
	id, _ := this.GetInt("id")
	if id == 0 {
		this.ResponseError(constants.SERVERERROR,"id 不能为空",id)
	}
	err,user := Dao.UserFind(id)

	find := Formatter.UserFormatter(user)
	if err != nil {
		beego.Info("查询失败",err)
		this.ResponseError(constants.SERVERERROR,"查询失败",err)
	}else {
		this.ResponseSuccess("查询成功",find)
	}
}

func (this *UserController) List()  {
	offset, _ := this.GetInt("offset",0)
	limit, _ := this.GetInt("limit",10)

	count,list,err := Dao.UserDaoList(offset,limit)

	if err != nil {
		this.ResponseError(constants.SERVERERROR,"查询失败",err)
	}
	data := map[string]interface{}{"count":count,"items":list,}

	this.ResponseSuccess("查询成功",data)
}


func (this *UserController) Update()  {

	id,_ := this.GetInt("id")
	name := this.GetString("name")
	if id == 0 {
		this.ResponseError(constants.SERVERERROR,"id 不能为空",id)
	}
	err,user := Dao.UserFind(id)

	if err != nil {
		beego.Info("查询失败",err)
		this.ResponseError(constants.SERVERERROR,"id 不存在",err)
	}
	num,err,user := api.UserBizUpdate(id,name)

	if err != nil {
		beego.Info("更新失败",err)
		this.ResponseError(constants.SERVERERROR,"更新失败",err)
	}
	beego.Info(num)
	this.ResponseSuccess("更新成功",user)
}

func (this *UserController) Delete() {

	id,_ := this.GetInt("id")

	json := api.UserBizDelete(id)

	this.Data["json"] = json
	this.ServeJSON()
}
