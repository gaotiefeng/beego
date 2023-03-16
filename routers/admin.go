package routers

import (
	"beego/app/controllers"
	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/admin", &controllers.BaseController{}, "get:Welcome")

}
