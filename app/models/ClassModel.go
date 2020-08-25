package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Class struct {
	ClassId int	`orm:"column(class_id);auto" description:"class_id" json:"class_id"`
	Name string	`orm:"column(class_name)" description:"class_name" json:"class_name"`
	Sort	int	`orm:"column(sort)" description:"sort" json:"sort"`
	Pid	int	`orm:"column(pid)" description:"pid" json:"pid"`
	Link	string	`orm:"column(link)" description:"link" json:"link"`
	CreatedAt time.Time	`orm:"column(created_at)"  description:"created_at" json:"created_at"`
}

type ClassTreeList struct {
	ClassId int	`orm:"column(class_id);auto" description:"class_id" json:"class_id"`
	Name string	`orm:"column(class_name)" description:"class_name" json:"class_name"`
	Sort	int	`orm:"column(sort)" description:"sort" json:"sort"`
	Pid	int	`orm:"column(pid)" description:"pid" json:"pid"`
	Link	string	`orm:"column(link)" description:"link" json:"link"`
	CreatedAt time.Time	`orm:"column(created_at)"  description:"created_at" json:"created_at"`
	Children []*ClassTreeList	`json:"children"`
}

func (u *Class) TableName() string {
	// db table name
	return "class"
}

func init() {
	orm.RegisterModel(new(Class))
}