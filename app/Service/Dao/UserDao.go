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

func UserDaoList(offset int,limit int) (int64,[]orm.ParamsList,error) {
	user := new(models.User)
	o := orm.NewOrm()

	var list [] orm.ParamsList
	count,err := o.QueryTable(user).Limit(limit,offset).ValuesList(&list,"id","name","mobile","password","created_at")

	return count,list,err
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

func UserDaoUpdate( id int ,name string) (num int64,err error,user models.User) {
	o := orm.NewOrm()

	user = models.User{}

	user.Id = id
	user.Name = name
	num,err = o.Update(&user)

	return num,err,user
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
