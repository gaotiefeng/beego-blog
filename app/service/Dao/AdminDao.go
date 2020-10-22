package Dao

import (
	"beego/app/models"
	"github.com/astaxie/beego/orm"
)

func GetByAdmin(name string)  (models.Admin,error){
	admin := models.Admin{UserName: name}
	o := orm.NewOrm()
	err := o.Read(&admin,"user_name")

	return admin,err
}
