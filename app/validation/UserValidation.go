package validation

import (
	"github.com/astaxie/beego/validation"
	"log"
)

type User struct {
	Name   string
	Mobile string
}

func RegisterValidation(name string, mobile string) (err error) {
	u := User{name, mobile}
	valid := validation.Validation{}
	valid.Required(u.Name, "name")
	valid.Mobile(u.Mobile, "mobile")
	if valid.HasErrors() {
		// validation does not pass
		// print invalid message
		for _, err := range valid.Errors {
			if err != nil {
				log.Println(err.Key, err.Message)
				return err
			}
		}
	}
	return
}
