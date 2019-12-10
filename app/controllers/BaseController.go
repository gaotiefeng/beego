package controllers

import (
	//"fmt"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) Welcome() {
	c.Data["json"] = map[string]interface{}{"code": 200, "message": "hello word"}
	c.ServeJSON()
}