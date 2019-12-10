package routers


import (
	"beego/app/controllers"
	"beego/app/controllers/api"
	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/",&controllers.BaseController{},"get:Welcome")
	beego.Router("/user/list",&api.UserController{},"get:List")

}