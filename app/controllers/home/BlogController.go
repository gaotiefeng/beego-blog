package home

import (
	"beego/app/controllers"
	"beego/app/service/Dao"
)

type BlogController struct {
	controllers.BaseController
}

/**
博客首页
 */
func (this *BlogController) Index() {

	treeList :=Dao.GetClass(0)
	this.Data["class"] = treeList
	this.Data["count"] = Dao.ArticleDaoCount()

	this.TplName = "home/article/index.html"
}



