package Dao

import (
	"beego/app/models"
	"github.com/astaxie/beego/orm"
	"log"
	"time"
)

func UserFind(id int) (err error, user models.User) {
	user = models.User{Id: id}
	o := orm.NewOrm()
	err = o.Read(&user)
	return err,user
}

func UserMobile(mobile string)  (err error, user models.User){
	user = models.User{Mobile: mobile}
	o := orm.NewOrm()
	err = o.Read(&user,"Mobile")
	log.Println(err)

	return err, user
}

func UserDaoList(offset int,limit int) (int64,*[]models.User,error) {
	user := new(models.User)
	o := orm.NewOrm()
	userModel := new([]models.User)
	count,err := o.QueryTable(user).Limit(limit,offset).All(userModel)

	//var list [] orm.Params
	//count,err := o.QueryTable(user).Limit(limit,offset).Values(&list,"id","name","mobile","password","created_at")

	return count,userModel,err
}

func UserDaoInsert(name string,mobile string,password string) (err error, user models.User) {
	//orm object
	o := orm.NewOrm()
	//struct object
	user = models.User{}
	//对结构体对象赋值
	user.Name = name
	user.Mobile = mobile
	user.Password = password
	user.CreatedAt = time.Now()
	//insert
	_,err = o.Insert(&user)
	return err,user
}

func UserDaoUpdate( id int ,name string) (num int64,err error,user models.User) {
	o := orm.NewOrm()

	user = models.User{}

	user.Id = id
	user.Name = name
	num,err = o.Update(&user, "Name")

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
