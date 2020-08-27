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

	offset, _ := this.GetInt("offset",0)
	limit, _ := this.GetInt("limit",10)
	offset = (offset - 1)*limit

	count,data,err := Dao.ArticleDaoList(offset,limit)
	if err != nil {
		beego.Informational("article is empty")
	}

	this.ResponseSuccess("文章列表",map[string]interface{}{"count":count,"items":data})
}