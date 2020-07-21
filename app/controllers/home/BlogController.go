package home

import "beego/app/controllers"

type BlogController struct {
	controllers.BaseController
}

func (this *BlogController) Index() {
	this.Data["content"] = "welcome to blog"
	this.Data["language"] = "go"
	this.TplName = "index.html"
}
