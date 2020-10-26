package home

import (
	"beego/app/controllers"
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
	this.Data["count"] = Dao.ArticleDaoCount()

	this.TplName = "home/article/index.html"
}



