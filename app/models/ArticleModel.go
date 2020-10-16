package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Article struct {
	Id      int		 `orm:"column(id);auto" description:"id" json:"id"`
	Name    string	`orm:"column(name)" description:"name" json:"name"`
	Content	string	`orm:"column(content)" description:"content" json:"content"`
	ClassId	int	`orm:"column(class_id)" description:"class_id" json:"class_id"`
	ChildId	int	`orm:"column(child_id)" description:"child_id" json:"child_id"`
	Image string `orm:"column(image)" description:"image" json:"image"`
	Click      int		 `orm:"column(click)" description:"click" json:"click"`
	CreatedAt	time.Time	`orm:"column(created_at)"  description:"created_at" json:"created_at"`
}

func (u *Article) TableName() string {
	// db table name
	return "article"
}

func init() {
	orm.RegisterModel(new(Article))
}
