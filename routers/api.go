package routers

import (
	"beego/app/controllers"
	"beego/app/controllers/api"
	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/",&controllers.BaseController{},"get:Welcome")

	beego.Router("/user/login",&api.UserController{},"get:Login")
	beego.Router("/user/register",&api.UserController{},"get:Register")
	beego.Router("/user/find",&api.UserController{},"get:Find")
	beego.Router("/user/list",&api.UserController{},"get:List")
	beego.Router("/user/update",&api.UserController{},"get:Update")
	beego.Router("/user/delete",&api.UserController{},"get:Delete")


}