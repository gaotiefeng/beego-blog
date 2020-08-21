package home

import (
	"beego/app/constants"
	"beego/app/controllers"
	"beego/app/service/Dao"
)

type BlogController struct {
	controllers.BaseController
}

/**

 */
func (this *BlogController) Index() {

	offset, _ := this.GetInt("offset",0)
	limit, _ := this.GetInt("limit",10)

	count,data,err := Dao.ArticleDaoList(offset,limit)
	if err != nil {
		this.ResponseError(constants.SERVERERROR,"查询失败",err)
	}

	this.Data["count"] = count
	this.Data["data"] = data
	this.Data["content"] = "welcome to blog"
	this.Data["language"] = "go controller"
	this.TplName = "home/index.html"
}

/**

 */
func (this *BlogController) Detail()  {

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
	this.TplName = "home/detail.html"
}

func (this *BlogController) Test() {

	this.Data["json"] = map[string]interface{}{"code": 200, "message": "hello word test"}
	this.ServeJSON()
}