package api

import (
	"beego/app/controllers"
)

type UserController struct {
	controllers.BaseController
}

func (c *UserController) List()  {
	c.Data["json"] = map[string]interface{}{"code": 200, "message": "api user list"}
	c.ServeJSON()
}
