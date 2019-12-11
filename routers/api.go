package routers


import (
	"beego/app/controllers"
	"beego/app/controllers/api"
	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/",&controllers.BaseController{},"get:Welcome")
	beego.Router("/user/index",&api.UserController{},"get:Index")
	beego.Router("/user/register",&api.UserController{},"get:Register")

}