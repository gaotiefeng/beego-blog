package home

import (
	"beego/app/controllers"
	"beego/app/models"
	"beego/app/service"
	"beego/app/service/Dao"
)

type BlogController struct {
	controllers.BaseController
}

/**
博客首页
 */
func (this *BlogController) Index() {

	classId,_ := this.GetInt("class_id",0)
	this.Data["class_id"] = classId
	childId,_ := this.GetInt("child_id",0)
	this.Data["child_id"] = childId

	homeBanner := service.GetByConfigKey("home_banner")
	this.Data["home_banner"] = homeBanner
	banner := service.GetByConfigKey("banner")
	this.Data["banner"] = banner

	treeList :=Dao.ClassGetAll(0)
	this.Data["class"] = treeList
	count := Dao.ArticleDaoCount(classId,childId)
	this.Data["count"] = count
	//	获取当前页码
	page,_ := this.GetInt("page",1)
	limit,_ := this.GetInt("limit",10)
	offset := limit*(page-1)

	_,data,_ := Dao.ArticleDaoList(offset,limit,classId,childId)
	this.Data["data"] = data
	// 分页数据
	paginatorMap := models.Paginator(page,limit,count)
	this.Data["paginator"]= paginatorMap //分页的数据

	_,hotData,_ := Dao.ArticleDaoHotList(offset,limit,classId,childId)

	this.Data["hot_data"] = hotData

	this.TplName = "home/article/index.html"
}



