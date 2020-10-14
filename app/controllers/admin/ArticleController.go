package admin

import (
	"beego/app/constants"
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

	page, _ := this.GetInt("page",0)
	limit, _ := this.GetInt("limit",10)
	offset := (page-1) * limit

	count,list,err := Dao.ArticleDaoList(offset,limit)

	if err != nil {
		this.ResponseError(constants.SERVERERROR,"查询失败",err)
	}

	this.ResponseListSuccess(constants.SUCCESS,"查询成功",list,count)
}
//添加文章页面
func (this *ArticleController) Add(){

	this.Layout = "admin/layout.html"
	this.TplName = "admin/article/add.html"
}

func (this *ArticleController) Save() {
	id, _ := this.GetInt("id",0)
	name := this.GetString("name")
	classId,_ := this.GetInt("class_id",0)
	content := this.GetString("content")
	image := this.GetString("image")
	beego.Error(id)
	beego.Error(name)
	beego.Error(content)
	beego.Error(classId)
	beego.Error(image)
	this.ResponseSuccess(constants.SUCCESS,"添加成功",image)
}
