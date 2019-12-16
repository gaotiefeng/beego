package Dao

import (
	"beego/app/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

func UserFind(id int) (model interface{}) {
	user := models.User{Id: id}
	o := orm.NewOrm()
	err := o.Read(&user)
	json := make(map[string]interface{})
	if err != nil {
		beego.Info("查询失败",err)
		json = map[string]interface{}{"code":404,"message":err}
	}else {
		json = map[string]interface{}{"code":200,"data":&user}
	}
	return json
}

func UserInsert(name string,mobile string) (model interface{}) {
	//orm object
	o := orm.NewOrm()
	//struct object
	user := models.User{}
	//对结构体对象赋值
	user.Name = name
	user.Mobile = mobile
	user.CreatedAt = time.Now()
	//insert
	_,err := o.Insert(&user)
	json := make(map[string]interface{})
	if err!= nil {
		//beego.Info("插入失败",err)
		json = map[string]interface{}{"code":500,"message":err}
	}else {
		json = map[string]interface{}{"code":200,"data":&user}
	}
	return json
}

func UserUpdate( id int ,name string) (model interface{}) {
	o := orm.NewOrm()

	user := models.User{}

	user.Id = id
	err := o.Read(&user)
	json := make(map[string]interface{})
	if err == nil {
		user.Name = name
		num,err := o.Update(&user)
		if err != nil {
			json = map[string]interface{}{"code":200,"message":"更新失败"}
		}else {
			json = map[string]interface{}{"code":200,"message":"更新成功","data":num}
		}
	}else {
		json = map[string]interface{}{"code":200,"message":"数据不存在"}
	}
	return json
}

func UserDelete(id int) (num int64, err error) {
	//orm object
	o := orm.NewOrm()
	//struct object
	user := models.User{}
	user.Id = id
	num, err = o.Delete(&user)
	/*json := make(map[string]interface{})
	if err != nil {
		json = map[string]interface{}{"code":500,"message":err}
	}else {
		beego.Info("删除成功",num)
		json = map[string]interface{}{"code":200,"data":num}
	}*/
	return num, err
}
