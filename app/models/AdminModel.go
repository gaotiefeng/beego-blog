package models

import "time"

type admin struct {
	Id      int		 `orm:"column(id);auto" description:"id" json:"id"`
	UserName    string	`orm:"column(user_name)" description:"user_name" json:"user_name"`
	Password	string	`orm:"column(password)" description:"password" json:"password"`
	Image string `orm:"column(image)" description:"image" json:"image"`
	CreatedAt	time.Time	`orm:"column(created_at)"  description:"created_at" json:"created_at"`
}
