package admin

import "github.com/astaxie/beego"

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Index()  {
	this.TplName = "admin/index.html"
}
