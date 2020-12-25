package home

import (
	"beego/app/controllers"
	"beego/app/service"
	"beego/app/service/Dao"
)

type IndexController struct {
	controllers.BaseController
}

/**
博客首页
*/
func (this *IndexController) Index() {

	//分类数据
	treeList :=Dao.ClassGetAll(0)
	this.Data["class"] = treeList
	//banner
	homeBanner := service.GetByConfigKey("home_banner")
	this.Data["home_banner"] = homeBanner
	banner := service.GetByConfigKey("banner")
	this.Data["banner"] = banner



	this.TplName = "home/index/index.html"
}