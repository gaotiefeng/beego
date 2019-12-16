package api

import (
	"beego/app/Service/Dao"
	"beego/app/constants"
	"beego/app/controllers"
	"github.com/astaxie/beego"
)

func UserBizDelete(id int) (json interface{}){
	if id == 0 {
		json = controllers.Error(constants.SERVERERROR,"id 不能为空",id)
		return json
	}
	num,err := Dao.UserDelete(id)
	var message string
	if err != nil {
		message = "删除失败"
	}else {
		message = "删除成功"
		beego.Info("删除成功",num)
	}
	json = controllers.Success(constants.SUCCESS,message,num)
	return json
}
