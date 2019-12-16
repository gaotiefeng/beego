package api

import (
	"beego/app/Service/Dao"
	"beego/app/constants"
	"beego/app/controllers"
	"beego/app/models"
	"github.com/astaxie/beego"
)

func UserBizUpdate(id int,name string)  (num int64,err error,user models.User){

	num,err,user = Dao.UserDaoUpdate(id, name)

	return num,err,user
}

func UserBizInsert(name string,mobile string,password string) (json interface{}) {

	err,user := Dao.UserDaoInsert(name,mobile,password)

	if err != nil {
		beego.Info("插入失败",err)
		json = controllers.Error(constants.SERVERERROR,"添加失败",user)
	}else {
		json = controllers.Success(constants.SUCCESS,"添加成功",user)
	}
	return json
}

func UserBizDelete(id int) (json interface{}){
	if id == 0 {
		json = controllers.Error(constants.SERVERERROR,"id 不能为空",id)
		return json
	}
	num,err := Dao.UserDaoDelete(id)
	var message string
	if err != nil {
		message = "删除失败"
		json = controllers.Error(constants.SERVERERROR,message,err)
	}else {
		message = "删除成功"
		beego.Info("删除成功",num)
		json = controllers.Success(constants.SUCCESS,message,num)
	}
	return json
}
