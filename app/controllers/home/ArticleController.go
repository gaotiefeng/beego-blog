package home

import (
	"beego/app/controllers"
	"beego/app/service/Dao"
	"github.com/astaxie/beego"
)
type ArticleController struct {
	controllers.BaseController
}

func (this *ArticleController) List()  {

	classId, _ := this.GetInt("class_id",0)
	childId, _ := this.GetInt("child_id",0)
	offset, _ := this.GetInt("offset",0)
	limit, _ := this.GetInt("limit",10)
	offset = (offset - 1)*limit

	count,data,err := Dao.ArticleDaoList(offset,limit,classId,childId)
	if err != nil {
		beego.Informational("article is empty")
	}

	this.ResponseSuccess("文章列表",map[string]interface{}{"count":count,"items":data})
}

/**

 */
func (this *ArticleController) Detail()  {
	id,_ := this.GetInt("id",0)

	if id == 0 {
		this.TplName = "404.tpl"
		return
	}

	article,error := Dao.ArticleDaoFind(id)
	if error != nil {
		this.TplName = "500.tpl"
		return
	}
	this.Data["data"] = article
	class,_ := Dao.ClassFirst(article.ClassId)
	this.Data["classInfo"] = class
	//头部
	treeList :=Dao.GetClass(0)
	this.Data["class"] = treeList

	this.TplName = "home/article/detail.html"
}