package validation

import (
	"beego/app/constants"
	"beego/app/controllers"
	"github.com/astaxie/beego/validation"
	"log"
)

type User struct {
	Name string
	Mobile string
}

func RegisterValidation(name string,mobile string) (json interface{}){
	u := User{name, mobile}
	valid := validation.Validation{}
	valid.Required(u.Name, "name")
	valid.MaxSize(u.Name, 15, "nameMax")
	valid.Mobile(u.Mobile, "mobile")
	if valid.HasErrors() {
		// validation does not pass
		// print invalid message
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			return controllers.Error(constants.SERVERERROR,"参数错误",err)
		}
	}
	return map[string]interface{}{}
}
