package routers


import (
	"beego/app/controllers"
	"beego/app/controllers/api"
	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/",&controllers.BaseController{},"get:Welcome")
	beego.Router("/user/find",&api.UserController{},"get:Find")
	beego.Router("/user/list",&api.UserController{},"get:List")
	beego.Router("/user/register",&api.UserController{},"get:Register")

}