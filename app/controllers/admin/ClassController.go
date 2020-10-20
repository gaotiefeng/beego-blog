package admin

import (
	"beego/app/constants"
	"beego/app/service/Dao"
	"github.com/astaxie/beego"
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

//分类页面
func (this *ClassController)Index()  {

	this.Layout = "admin/layout.html"
	this.TplName = "admin/class/index.html"
}
//分类添加页面
func (this *ClassController)Add()  {

	class,_ := Dao.ClassFindParent(0)
	this.Data["class"] = class
	beego.Info(class)
	this.Layout = "admin/layout.html"
	this.TplName = "admin/class/add.html"
}
//所有分类数据
func (this *ClassController)List()  {
	page,_ := this.GetInt("page")
	limit,_ := this.GetInt("limit")
	offset := this.Page(page,limit)

	list,err := Dao.ClassList(offset,limit)
	count := Dao.ClassDaoCount()
	if err != nil {
		this.ResponseError(constants.SERVERERROR,"err",err)
	}
	this.ResponseListSuccess(constants.SUCCESS,"查询成功",list,count)
}