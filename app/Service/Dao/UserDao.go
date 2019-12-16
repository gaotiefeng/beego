package Dao

import (
	"beego/app/models"
	"github.com/astaxie/beego/orm"
	"time"
)

func UserFind(id int) (err error, user models.User) {
	user = models.User{Id: id}
	o := orm.NewOrm()
	err = o.Read(&user)
	/*json := make(map[string]interface{})
	if err != nil {
		beego.Info("查询失败",err)
		json = map[string]interface{}{"code":404,"message":err}
	}else {
		json = map[string]interface{}{"code":200,"data":&user}
	}*/
	return err,user
}

func UserDaoInsert(name string,mobile string) (err error, user models.User) {
	//orm object
	o := orm.NewOrm()
	//struct object
	user = models.User{}
	//对结构体对象赋值
	user.Name = name
	user.Mobile = mobile
	user.CreatedAt = time.Now()
	//insert
	_,err = o.Insert(&user)
	/*json := make(map[string]interface{})
	if err!= nil {
		beego.Info("插入失败",err)
		json = map[string]interface{}{"code":500,"message":err}
	}else {
		json = map[string]interface{}{"code":200,"data":&user}
	}*/
	return err,user
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

func UserDaoDelete(id int) (num int64, err error) {
	//orm object
	o := orm.NewOrm()
	//struct object
	user := models.User{}
	user.Id = id
	num, err = o.Delete(&user)

	return num, err
}
