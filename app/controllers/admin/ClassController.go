package admin

import (
	"beego/app/constants"
	"beego/app/service/Admin"
	"beego/app/service/Dao"
	"github.com/astaxie/beego"
)

type ClassController struct {
	LoginController
}

//获取下级分类
func (this *ClassController)GetChild()  {
	id,_ := this.GetInt("id")

	class := Dao.ClassGetAll(id)

	this.ResponseSuccess(constants.SUCCESS,"子集数据",class)
}

//分类页面
func (this *ClassController)Index()  {

	this.Layout = "admin/layout.html"
	this.TplName = "admin/class/index.html"
}
//分类添加页面
func (this *ClassController)Add()  {
	id,_ := this.GetInt("id",0)

	classInfo,_ := Dao.ClassFirst(id)
	this.Data["classInfo"] = classInfo

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
//更新数据
func (this *ClassController) Save() {
	id,_ := this.GetInt("id",0)
	name := this.GetString("name")
	sort,_ := this.GetInt("sort",0)
	pid,_ := this.GetInt("pid",0)
	num,err := Admin.ServiceAdminClassSave(id,name,sort,pid)
	if err != nil {
		this.ResponseError(constants.SERVERERROR,"添加失败",num)
	}
	this.ResponseSuccess(constants.SUCCESS,"添加成功",num)
}

func (this *ClassController) Del()  {
	id,_ := this.GetInt("id",0)
	num, err := Dao.ClassDaoDel(id)
	if err != nil {
		this.ResponseError(constants.SERVERERROR,"删除失败",num)
	}
	this.ResponseSuccess(constants.SUCCESS,"删除成功",num)
}