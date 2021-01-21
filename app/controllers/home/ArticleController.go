package home

import (
	"beego/app/controllers"
	"beego/app/models"
	"beego/app/service"
	"beego/app/service/Dao"
	"github.com/astaxie/beego"
)
type ArticleController struct {
	controllers.BaseController
}

/**
文章列表
*/
func (this *ArticleController) Index() {

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
	treeList :=Dao.ClassGetAll(0)
	this.Data["class"] = treeList
	//上一篇
	pre,_ :=Dao.ArticleDaoPre(id)
	this.Data["pre"] = pre
	next,_ :=Dao.ArticleDaoNext(id)
	this.Data["next"] = next
	//点击加1
	var articleClick = models.Article{Id: article.Id,Click: article.Click+1}
	Dao.ArticleDaoUpdateClick(articleClick)
	//相关文章
	_,articleList,_ := Dao.ArticleDaoList(0,10,article.ClassId,0)
	this.Data["articleList"] = articleList

	this.TplName = "home/article/detail.html"
}