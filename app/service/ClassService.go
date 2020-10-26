package service

import (
	"beego/app/models"
	"beego/app/service/Dao"
)

func GetClassChild(pid int) []*models.ClassTreeList {

	return Dao.ClassGetAll(pid)
}
