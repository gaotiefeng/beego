package Formatter

import (
	"beego/app/models"
)

func UserFormatter(model models.User) map[string]interface{} {
	return map[string]interface{}{
		"id" : model.Id,
		"name" : model.Name,
		"mobile" : model.Mobile,
		"password" : model.Password,
		"created_at" : model.CreatedAt,
	}
}
