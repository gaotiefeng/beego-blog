package home

import "beego/app/controllers"

type BlogController struct {
	controllers.BaseController
}

/**

 */
func (this *BlogController) Index() {
	this.Data["content"] = "welcome to blog"
	this.Data["language"] = "go controller"
	this.TplName = "home/index.html"
}

/**

 */
func (this *BlogController) Detail()  {

	this.TplName = "home/blog.html"
}

func (this *BlogController) Test() {

	println("12211221")
	this.Data["json"] = map[string]interface{}{"code": 200, "message": "hello word test"}
	this.ServeJSON()
}