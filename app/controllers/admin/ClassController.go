package admin

import (
	"beego/app/constants"
	"beego/app/service/Dao"
)

type ClassController struct {
	LoginController
}

//获取下级分类
func (this *ClassController)GetChild()  {
	id,_ := this.GetInt("id")

	class := Dao.GetClass(id)

	this.ResponseSuccess(constants.SUCCESS,"子集数据",class)
}