package admin

import (
	"beego/app/constants"
	"beego/app/models"
	"beego/app/service/Admin"
	"beego/app/service/Dao"
	"github.com/astaxie/beego"
)

type ArticleController struct {
	LoginController
}
//页面
func (this *ArticleController) Index()  {
	this.Layout = "admin/layout.html"
	this.TplName = "admin/article/list.html"
}
//json文章列表
func (this *ArticleController) List()  {

	classId, _ := this.GetInt("class_id",0)
	childId, _ := this.GetInt("child_id",0)
	page, _ := this.GetInt("page",0)
	limit, _ := this.GetInt("limit",10)
	offset := (page-1) * limit

	_,list,err := Dao.ArticleDaoList(offset,limit,classId,childId)
	count := Dao.ArticleDaoCount()
	if err != nil {
		this.ResponseError(constants.SERVERERROR,"查询失败",err)
	}

	this.ResponseListSuccess(constants.SUCCESS,"查询成功",list,count)
}

//添加文章页面
func (this *ArticleController) Add(){
	id, _ := this.GetInt("id",0)
	article := models.Article{}
	classInfo := models.Class{}
	childInfo := models.Class{}
	if id != 0 {
		article,_ = Dao.ArticleDaoFind(id)
		classInfo,_ = Dao.ClassFirst(article.ClassId)
		childInfo,_ = Dao.ClassFirst(article.ChildId)
	}
	this.Data["article"] = article
	this.Data["classInfo"] = classInfo
	this.Data["childInfo"] = childInfo

	class,_ := Dao.ClassFindParent(0)
	this.Data["class"] = class
	classAll,_ := Dao.ClassFindAll()
	this.Data["class_all"] = classAll
	beego.Info(class)
	this.Layout = "admin/layout.html"
	this.TplName = "admin/article/add.html"
}

//添加|更新文章
func (this *ArticleController) Save() {
	id, _ := this.GetInt("id",0)
	name := this.GetString("name")
	classId,_ := this.GetInt("class_id",0)
	childId,_ := this.GetInt("child_id",0)
	content := this.GetString("content")
	image := this.GetString("image")
	num,err := Admin.ArticleSave(id,name,content,image,classId,childId)
	if err != nil {
		this.ResponseError(constants.SERVERERROR,"添加失败",num)
	}
	this.ResponseSuccess(constants.SUCCESS,"添加成功",num)
}
//删除文章
func (this *ArticleController) Del(){
	id, _ := this.GetInt("id",0)

	num, err := Dao.ArticleDaoDelete(id)
	if err != nil {
		this.ResponseError(constants.SERVERERROR,"删除失败",num)
	}
	this.ResponseSuccess(constants.SUCCESS,"删除成功",num)
}

//全部文章数据
func (this *ArticleController) DataAll() {
	_,data,err := Dao.ArticleDaoDataAll()
	if err != nil {
		this.ResponseError(constants.SERVERERROR,"查询失败",data)
	}
	this.ResponseSuccess(constants.SUCCESS,"数据列表",data)
}