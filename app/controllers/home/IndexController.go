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
	for key,value := range treeList{
		_, art, _ := Dao.ArticleDaoList(0,6,value.ClassId,0)
		treeList[key].Art = art
	}
	this.Data["class"] = treeList
	//banner
	homeBanner := service.GetByConfigKey("home_banner")
	this.Data["home_banner"] = homeBanner
	banner := service.GetByConfigKey("banner")
	this.Data["banner"] = banner
	//最近热文
	_,articleModel,_ := Dao.ArticleDaoHotList(0,20,0,0)
	this.Data["article"] = articleModel


	this.TplName = "home/index/index.html"
}