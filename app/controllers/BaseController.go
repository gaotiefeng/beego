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

func Success(code int,message string,data interface{}) (json interface{}){

	json = map[string]interface{}{"code":code,"message":message,"data":data}

	return json
}

func Error(code int,message string, data interface{}) (json interface{}) {

	json = map[string]interface{}{"code":code,"message":message,"data":data}
	return json
}