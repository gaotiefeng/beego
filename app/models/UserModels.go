package models

import "time"

type User struct {
	Id      int		`orm:"column(id);auto" description:"id"`
	Name    string	`orm:"column(name)" description:"name"`
	Mobile	string	`orm:"column(mobile)" description:"mobile"`
	Password string `orm:"column(password)" description:"password"`
	CreatedAt	time.Time	`orm:"column(created_at)"  description:"created_at"`
}
