package models

import "time"

type User struct {
	Id        int       `orm:"column(id);auto" description:"id" json:"id"`
	Name      string    `orm:"column(name)" description:"name" json:"name"`
	Mobile    string    `orm:"column(mobile)" description:"mobile" json:"mobile"`
	Password  string    `orm:"column(password)" description:"password" json:"password"`
	CreatedAt time.Time `orm:"column(created_at)"  description:"created_at" json:"created_at"`
}
