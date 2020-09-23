package admin

import (
	"beego/app/constants"
	"beego/app/service/Dao"
)

type ArticleController struct {
	LoginController
}


func (this *ArticleController) List()  {

	page, _ := this.GetInt("page",0)
	limit, _ := this.GetInt("limit",10)
	offset := (page-1) * limit

	count,list,err := Dao.ArticleDaoList(offset,limit)

	if err != nil {
		this.ResponseError(constants.SERVERERROR,"查询失败",err)
	}

	this.ResponseSuccess(constants.SUCCESS,"查询成功",list,count)
}

func (this *ArticleController) Add(){

	this.Layout = "admin/layout.html"
	this.TplName = "admin/article/add.html"
}
