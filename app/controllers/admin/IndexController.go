package admin

import "github.com/astaxie/beego"

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Index()  {
	this.Layout = "admin/layout.html"
	this.TplName = "admin/index.html"
}
